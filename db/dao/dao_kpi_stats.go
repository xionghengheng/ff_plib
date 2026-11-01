package dao

import (
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

type KPIBusinessInterface interface {
	GetKPIBusinessStats(begin, end int64, currentMonth bool) (*model.KPIBusinessStats, error)
}
type KPIBusinessInterfaceImp struct{}

var ImpKPIBusiness KPIBusinessInterface = &KPIBusinessInterfaceImp{}

// 付费与退款后的订单均保留原成交事实；按课包所属学员判定首购。
const kpiPaidOrders = ` FROM payment_orders o JOIN course_packages p ON p.package_id=o.package_id
 WHERE p.package_type=2 AND COALESCE(p.is_gift_package,0)=0 AND o.purchase_type=3
 AND o.order_status IN (1,4) AND o.payment_time>0 AND o.payment_amount>0 `
const kpiFirstPaid = `SELECT p.uid, MIN(o.payment_time) first_paid_ts` + kpiPaidOrders + ` GROUP BY p.uid`

func (imp *KPIBusinessInterfaceImp) GetKPIBusinessStats(begin, end int64, currentMonth bool) (*model.KPIBusinessStats, error) {
	result := &model.KPIBusinessStats{Daily: []model.KPIDailyStat{}}
	if end <= begin {
		return result, nil
	}
	var sales, first, lessons []model.KPIDailyStat
	if err := db.Get().Raw(`SELECT FLOOR((o.payment_time-?)/86400) day_index,
 COALESCE(SUM(o.payment_amount),0) gmv_fen, COALESCE(SUM(o.course_cnt),0) sold_lessons`+kpiPaidOrders+`
 AND o.payment_time>=? AND o.payment_time<? GROUP BY day_index`, begin, begin, end).Scan(&sales).Error; err != nil {
		return nil, err
	}
	if err := db.Get().Raw(`SELECT FLOOR((first_paid_ts-?)/86400) day_index, COUNT(*) new_paid_users
 FROM (`+kpiFirstPaid+`) f WHERE first_paid_ts>=? AND first_paid_ts<? GROUP BY day_index`, begin, begin, end).Scan(&first).Error; err != nil {
		return nil, err
	}
	if err := db.Get().Raw(`SELECT FLOOR((l.write_off_ts-?)/86400) day_index, COUNT(DISTINCT l.lesson_id) lessons
 FROM course_package_single_lessons l JOIN course_packages p ON p.package_id=l.package_id
 WHERE p.package_type=2 AND COALESCE(p.is_gift_package,0)=0 AND COALESCE(l.is_gift_lesson,0)=0
 AND l.status=2 AND l.write_off_ts>=? AND l.write_off_ts<? GROUP BY day_index`, begin, begin, end).Scan(&lessons).Error; err != nil {
		return nil, err
	}
	result.Daily = append(result.Daily, sales...)
	result.Daily = append(result.Daily, first...)
	result.Daily = append(result.Daily, lessons...)
	if err := db.Get().Raw(`SELECT COUNT(*) refund_count, COALESCE(SUM(refund_amount),0)*100 refund_fen
 FROM course_package_refund_record WHERE refund_ts>=? AND refund_ts<? AND refund_amount>0`, end-30*86400, end).Scan(result).Error; err != nil {
		return nil, err
	}
	// 当前课包剩余资产不能用于还原历史月份风险，历史月份明确返回null。
	if currentMonth {
		type counts struct{ Count int64 }
		var zero, overdue counts
		active := ` FROM course_packages p WHERE p.package_type=2 AND COALESCE(p.is_gift_package,0)=0 AND p.remain_cnt>0 `
		completed := ` AND NOT EXISTS (SELECT 1 FROM course_package_single_lessons l JOIN course_packages q ON q.package_id=l.package_id
   WHERE l.uid=p.uid AND q.package_type=2 AND COALESCE(q.is_gift_package,0)=0 AND COALESCE(l.is_gift_lesson,0)=0 AND l.status=2) `
		if err := db.Get().Raw(`SELECT COUNT(DISTINCT p.uid) count` + active + completed).Scan(&zero).Error; err != nil {
			return nil, err
		}
		if err := db.Get().Raw(`SELECT COUNT(DISTINCT p.uid) count`+active+completed+`
   AND EXISTS (SELECT 1 FROM (`+kpiFirstPaid+`) f WHERE f.uid=p.uid AND f.first_paid_ts<?)
   AND NOT EXISTS (SELECT 1 FROM course_package_single_lessons l JOIN course_packages q ON q.package_id=l.package_id
    WHERE l.uid=p.uid AND q.package_type=2 AND COALESCE(q.is_gift_package,0)=0 AND COALESCE(l.is_gift_lesson,0)=0 AND l.status=1)`, end-7*86400).Scan(&overdue).Error; err != nil {
			return nil, err
		}
		result.ZeroLessonUsers = &zero.Count
		result.FirstLessonOverdueUsers = &overdue.Count
	}
	return result, nil
}

// 只通过客资DAO访问客资表；一条客资只计一次，避免节点join放大人数。
func (imp *UserTrackManagementInterfaceImp) GetKPITrackStats(begin, end int64) ([]model.KPITrackStat, error) {
	result := []model.KPITrackStat{}
	if end <= begin {
		return result, nil
	}
	err := db.Get().Raw(`SELECT COALESCE(t.advisor_name,'') advisor_name, COUNT(*) leads,
 SUM(CASE WHEN EXISTS(SELECT 1 FROM user_track_node n WHERE n.track_id=t.track_id AND n.stage=4 AND n.node_create_ts<?) THEN 1 ELSE 0 END) appointments,
 SUM(CASE WHEN EXISTS(SELECT 1 FROM user_track_node n WHERE n.track_id=t.track_id AND n.stage=5 AND n.node_create_ts<?) THEN 1 ELSE 0 END) visits,
 SUM(CASE WHEN EXISTS(SELECT 1 FROM user_track_node n JOIN payment_orders o ON o.package_id=n.pay_package_id
  JOIN course_packages p ON p.package_id=o.package_id
  WHERE n.track_id=t.track_id AND n.stage=7 AND n.node_create_ts<? AND o.order_status IN (1,4)
  AND o.payment_time>0 AND o.payment_time<? AND o.payment_amount>0 AND o.purchase_type=3
  AND p.package_type=2 AND COALESCE(p.is_gift_package,0)=0) THEN 1 ELSE 0 END) deals
 FROM user_track t WHERE t.created_ts>=? AND t.created_ts<? GROUP BY t.advisor_name`, end, end, end, end, begin, end).Scan(&result).Error
	return result, err
}

func (imp *UserTrackManagementInterfaceImp) GetKPITrackDailyStats(begin, end int64) ([]model.KPIDailyStat, error) {
	result := []model.KPIDailyStat{}
	if end <= begin {
		return result, nil
	}
	var leads, appointments []model.KPIDailyStat
	if err := db.Get().Raw(`SELECT FLOOR((created_ts-?)/86400) day_index, COUNT(*) leads FROM user_track
 WHERE created_ts>=? AND created_ts<? GROUP BY day_index`, begin, begin, end).Scan(&leads).Error; err != nil {
		return nil, err
	}
	if err := db.Get().Raw(`SELECT FLOOR((first_ts-?)/86400) day_index, COUNT(*) appointments
 FROM (SELECT track_id,MIN(node_create_ts) first_ts FROM user_track_node WHERE stage=4 GROUP BY track_id) n
 WHERE first_ts>=? AND first_ts<? GROUP BY day_index`, begin, begin, end).Scan(&appointments).Error; err != nil {
		return nil, err
	}
	result = append(result, leads...)
	result = append(result, appointments...)
	return result, nil
}
