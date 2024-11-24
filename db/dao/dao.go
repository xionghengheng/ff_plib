package dao

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
	"strings"
	"time"
)

const user_tableName = "user_info"

func (imp *UserInterfaceImp) UpsertUser(user *model.UserInfoModel) error {
	cli := db.Get()
	return cli.Table(user_tableName).Save(user).Error
}

func (imp *UserInterfaceImp) UpdateUserPhone(uid int64, phone string) error {
	cli := db.Get()
	return cli.Table(user_tableName).Model(&model.UserInfoModel{}).Where("user_id = ?", uid).Update("phone_number", phone).Error
}

func (imp *UserInterfaceImp) UpdateUserInfo(uid int64, mapUpdates map[string]interface{}) error {
	cli := db.Get()
	return cli.Table(user_tableName).Model(&model.UserInfoModel{}).Where("user_id = ?", uid).Updates(mapUpdates).Error
}

func (imp *UserInterfaceImp) GetUser(uid int64) (*model.UserInfoModel, error) {
	var err error
	var user = new(model.UserInfoModel)
	cli := db.Get()
	err = cli.Table(user_tableName).Where("user_id = ?", uid).First(user).Error
	return user, err
}

func (imp *UserInterfaceImp) GetUserByOpenId(openid string) (*model.UserInfoModel, error) {
	var err error
	var user = new(model.UserInfoModel)

	cli := db.Get()
	err = cli.Table(user_tableName).Where("wechat_id = ?", openid).First(user).Error

	return user, err
}

func (imp *UserInterfaceImp) GetUserByPhone(phone string) (*model.UserInfoModel, error) {
	var err error
	var user = new(model.UserInfoModel)
	cli := db.Get()
	err = cli.Table(user_tableName).Where("phone_number = ?", phone).First(user).Error
	return user, err
}

func (imp *UserInterfaceImp) GetUserByTraceId(traceid string) (*model.UserInfoModel, error) {
	var err error
	var user = new(model.UserInfoModel)
	cli := db.Get()
	err = cli.Table(user_tableName).Where("head_pic_safe_trace_id = ?", traceid).First(user).Error
	return user, err
}

func (imp *UserInterfaceImp) RemoveUser(openid string) error {
	return db.Get().Table(user_tableName).Where("wechat_id = ?", openid).Delete(&model.UserInfoModel{}).Error
}

func (imp *UserInterfaceImp) GetAllUser() ([]model.UserInfoModel, error) {
	var err error
	var allUser []model.UserInfoModel
	cli := db.Get()
	err = cli.Table(user_tableName).Find(&allUser).Error
	return allUser, err
}

func (imp *UserInterfaceImp) GetUserByCoachId(coachId int) (*model.UserInfoModel, error) {
	var err error
	var user = new(model.UserInfoModel)
	cli := db.Get()
	err = cli.Table(user_tableName).Where("coach_id = ?", coachId).First(user).Error
	return user, err
}

const sms_tableName = "verification_codes"

func (imp *SmsInterfaceImp) AddCode(code *model.SmsVerificationCodeModel) error {
	cli := db.Get()
	return cli.Table(sms_tableName).Save(code).Error
}

func (imp *SmsInterfaceImp) UpdateCodeUserStatus(uniqid string) error {
	cli := db.Get()
	return cli.Table(sms_tableName).Model(&model.SmsVerificationCodeModel{}).Where("unid = ?", uniqid).Update("used", 1).Error
}

func (imp *SmsInterfaceImp) GetCode(uniqid string) (*model.SmsVerificationCodeModel, error) {
	var err error
	var code = new(model.SmsVerificationCodeModel)

	cli := db.Get()
	err = cli.Table(sms_tableName).Where("unid = ? AND used = ?", uniqid, 0).Order("createts DESC").First(code).Error

	return code, err
}

const gym_tableName = "gym_info"

func (imp *GymInterfaceImp) GetGymList() ([]model.GymInfoModel, error) {
	var err error
	var vecGyms []model.GymInfoModel

	cli := db.Get()
	err = cli.Table(gym_tableName).Find(&vecGyms).Order("gym_id ASC").Error
	return vecGyms, err
}

func (imp *GymInterfaceImp) GetGymInfoByGymId(gymId int) (model.GymInfoModel, error) {
	var err error
	var stGym model.GymInfoModel

	cli := db.Get()
	err = cli.Table(gym_tableName).Where("gym_id = ?", gymId).Find(&stGym).Error
	return stGym, err
}

const course_tableName = "courses"

func (imp *CourseInterfaceImp) GetCourseList() ([]model.CourseModel, error) {
	var err error
	var vecCourseModel []model.CourseModel

	cli := db.Get()
	err = cli.Table(course_tableName).Find(&vecCourseModel).Order("course_id ASC").Error
	return vecCourseModel, err
}

func (imp *CourseInterfaceImp) GetCourseById(id int) (*model.CourseModel, error) {
	var err error
	var couse = new(model.CourseModel)

	cli := db.Get()
	err = cli.Table(course_tableName).Where("course_id = ?", id).First(couse).Error

	return couse, err
}

const coach_tableName = "coaches"

func (imp *CoachInterfaceImp) GetCoachListByGymId(gymId int) ([]model.CoachModel, error) {
	//var err error
	//var vecCoachModel []model.CoachModel

	//cli := db.Get()
	//err = cli.Table(coach_tableName).Where("gym_id = ?", gymId).Order("priority DESC").Find(&vecCoachModel).Error
	//return vecCoachModel, err
	var vecCoachModel []model.CoachModel
	cli := db.Get()
	return vecCoachModel, cli.Raw("SELECT * FROM coaches WHERE gym_id = ? ORDER BY priority DESC", gymId).Scan(&vecCoachModel).Error
}

func (imp *CoachInterfaceImp) GetCoachById(id int) (*model.CoachModel, error) {
	var err error
	var coach = new(model.CoachModel)
	cli := db.Get()
	err = cli.Table(coach_tableName).Where("coach_id = ?", id).First(coach).Error
	return coach, err
}

func (imp *CoachInterfaceImp) GetCoachAll() ([]model.CoachModel, error) {
	var err error
	var vecCoachModel []model.CoachModel
	cli := db.Get()
	err = cli.Table(coach_tableName).Find(&vecCoachModel).Limit(100).Error
	return vecCoachModel, err
}

const course_package_tableName = "course_packages"

func (imp *CoursePackageInterfaceImp) GetCoursePackageById(packageId string) (*model.CoursePackageModel, error) {
	var err error
	var coursePackage = new(model.CoursePackageModel)
	cli := db.Get()
	err = cli.Table(course_package_tableName).Where("package_id = ?", packageId).First(coursePackage).Error
	return coursePackage, err
}

func (imp *CoursePackageInterfaceImp) GetPayCoursePackageList(uid int64) ([]model.CoursePackageModel, error) {
	var err error
	var vecCoursePackageModel []model.CoursePackageModel
	cli := db.Get()
	err = cli.Table(course_package_tableName).Where("uid = ? AND package_type = ?", uid, model.Enum_PackageType_PaidPackage).Order("ts DESC").Find(&vecCoursePackageModel).Error
	return vecCoursePackageModel, err
}

func (imp *CoursePackageInterfaceImp) GetTrailCoursePackage(uid int64) (*model.CoursePackageModel, error) {
	var err error
	var coursePackage = new(model.CoursePackageModel)
	cli := db.Get()
	err = cli.Table(course_package_tableName).Where("uid = ? AND package_type = ?", uid, model.Enum_PackageType_TrialFree).First(coursePackage).Error
	return coursePackage, err
}

func (imp *CoursePackageInterfaceImp) GetAllCoursePackageListByCoachId(coachId int, limit int) ([]model.CoursePackageModel, error) {
	var err error
	var vecCoursePackageModel []model.CoursePackageModel
	cli := db.Get()
	err = cli.Table(course_package_tableName).Where("coach_id = ?", coachId).Order("ts DESC").Find(&vecCoursePackageModel).Limit(limit).Error
	return vecCoursePackageModel, err
}

func (imp *CoursePackageInterfaceImp) GetPayCoursePackageListByCoachId(coachId int, limit int) ([]model.CoursePackageModel, error) {
	var err error
	var vecCoursePackageModel []model.CoursePackageModel
	cli := db.Get()
	err = cli.Table(course_package_tableName).Where("coach_id = ? AND package_type = ?", coachId, model.Enum_PackageType_PaidPackage).Order("ts DESC").Find(&vecCoursePackageModel).Limit(limit).Error
	return vecCoursePackageModel, err
}

func (imp *CoursePackageInterfaceImp) GetTrailCoursePackageListByCoachId(coachId int, limit int) ([]model.CoursePackageModel, error) {
	var err error
	var vecCoursePackageModel []model.CoursePackageModel
	cli := db.Get()
	err = cli.Table(course_package_tableName).Where("coach_id = ? AND package_type = ?", coachId, model.Enum_PackageType_TrialFree).Order("ts DESC").Find(&vecCoursePackageModel).Limit(limit).Error
	return vecCoursePackageModel, err
}

func (imp *CoursePackageInterfaceImp) GetListByCoachIdAndLastFinishLessonTs(coachId int, limit int) ([]model.CoursePackageModel, error) {
	var err error
	var vecCoursePackageModel []model.CoursePackageModel
	cli := db.Get()
	err = cli.Table(course_package_tableName).Select("package_id, uid, last_lesson_ts").Where("coach_id = ?", coachId).Order("last_lesson_ts ASC").Find(&vecCoursePackageModel).Limit(limit).Error
	return vecCoursePackageModel, err
}

func (imp *CoursePackageInterfaceImp) GetPayCoursePackageListByCoachIdAndUid(coachId int, uid int64) ([]model.CoursePackageModel, error) {
	var err error
	var vecCoursePackageModel []model.CoursePackageModel
	cli := db.Get()
	err = cli.Table(course_package_tableName).Where("coach_id = ? AND uid = ? AND package_type = ?", coachId, uid, model.Enum_PackageType_PaidPackage).Order("ts DESC").Find(&vecCoursePackageModel).Error
	return vecCoursePackageModel, err
}

func (imp *CoursePackageInterfaceImp) GetAllPackageListByCoachIdAndUid(coachId int, uid int64) ([]model.CoursePackageModel, error) {
	var err error
	var vecCoursePackageModel []model.CoursePackageModel
	cli := db.Get()
	err = cli.Table(course_package_tableName).Where("coach_id = ? AND uid = ?", coachId, uid).Order("ts DESC").Find(&vecCoursePackageModel).Error
	return vecCoursePackageModel, err
}

func (imp *CoursePackageInterfaceImp) GetCoursePackageListByUid(uid int64) ([]model.CoursePackageModel, error) {
	var err error
	var vecCoursePackageModel []model.CoursePackageModel

	cli := db.Get()
	err = cli.Table(course_package_tableName).Where("uid = ?", uid).Order("ts DESC").Find(&vecCoursePackageModel).Error
	return vecCoursePackageModel, err
}

func (imp *CoursePackageInterfaceImp) AddCoursePackage2Uid(stCoursePackageModel *model.CoursePackageModel) error {
	cli := db.Get()
	return cli.Table(course_package_tableName).Save(stCoursePackageModel).Error
}

func (imp *CoursePackageInterfaceImp) FindSamePackage(uid int64, gymId int, coachId int, courseId int) (*model.CoursePackageModel, error) {
	var err error
	var coursePackage = new(model.CoursePackageModel)
	cli := db.Get()
	err = cli.Table(course_package_tableName).Where("uid = ? AND gym_id = ? AND coach_id = ? AND course_id = ?",
		uid, gymId, coachId, courseId).First(coursePackage).Error
	return coursePackage, err
}

func (imp *CoursePackageInterfaceImp) AddLessonAndSubCourseCnt(packageId string, uniId string, stCoursePackageSingleLessonModel *model.CoursePackageSingleLessonModel) error {
	cli := db.Get()
	err := cli.Transaction(func(tx *gorm.DB) error {

		err := cli.Table(course_package_single_lesson_tableName).Save(stCoursePackageSingleLessonModel).Error
		if err != nil {
			fmt.Printf("Save lesson err, err:%+v packageId:%s stCoursePackageSingleLessonModel:%+v\n", err, packageId, stCoursePackageSingleLessonModel)
			tx.Rollback()
			return err
		}

		var coursePackage = new(model.CoursePackageModel)
		err = cli.Table(course_package_tableName).Where("package_id = ?", packageId).First(coursePackage).Error
		if err != nil {
			fmt.Printf("get coursePackage err, err:%+v packageId:%s\n", err, packageId)
			tx.Rollback()
			return err
		}

		var vecUniidList []string
		if len(coursePackage.UniidList) > 0 {
			vecUniidList = strings.Split(coursePackage.UniidList, ",")
			for _, v := range vecUniidList {
				if v == uniId {
					fmt.Printf("dup uniid coursePackage return, packageId:%s uniId:%s\n", packageId, uniId)
					return nil
				}
			}
		}

		newSlice := append([]string{uniId}, vecUniidList...)
		if len(newSlice) > 30 {
			newSlice = newSlice[0:30]
		}
		newUniidList := ""
		for _, v := range newSlice {
			newUniidList += ("," + v)
		}
		newUniidList = strings.TrimLeft(newUniidList, ",")
		newUniidList = strings.TrimRight(newUniidList, ",")
		mapUpdates := map[string]interface{}{}
		mapUpdates["remain_cnt"] = gorm.Expr("remain_cnt - ?", 1)
		mapUpdates["uniid_list"] = newUniidList
		err = cli.Table(course_package_tableName).Model(&model.CoursePackageModel{}).Where("package_id = ?", packageId).Updates(mapUpdates).Error
		if err != nil {
			fmt.Printf("update remain_cnt err, err:%+v packageId:%s mapUpdates:%+v\n", err, packageId, mapUpdates)
			tx.Rollback()
			return err
		}

		return nil
	})
	return err
}

func (imp *CoursePackageInterfaceImp) SubCourseCnt(packageId string, uniId string) error {
	cli := db.Get()
	err := cli.Transaction(func(tx *gorm.DB) error {

		var coursePackage = new(model.CoursePackageModel)
		err := cli.Table(course_package_tableName).Where("package_id = ?", packageId).First(coursePackage).Error
		if err != nil {
			fmt.Printf("get coursePackage err, err:%+v packageId:%s\n", err, packageId)
			tx.Rollback()
			return err
		}

		var vecUniidList []string
		if len(coursePackage.UniidList) > 0 {
			vecUniidList = strings.Split(coursePackage.UniidList, ",")
			for _, v := range vecUniidList {
				if v == uniId {
					fmt.Printf("dup uniid coursePackage return, packageId:%s uniId:%s\n", packageId, uniId)
					return nil
				}
			}
		}

		newSlice := append([]string{uniId}, vecUniidList...)
		if len(newSlice) > 30 {
			newSlice = newSlice[0:30]
		}
		newUniidList := ""
		for _, v := range newSlice {
			newUniidList += ("," + v)
		}
		newUniidList = strings.TrimLeft(newUniidList, ",")
		newUniidList = strings.TrimRight(newUniidList, ",")
		mapUpdates := map[string]interface{}{}
		mapUpdates["remain_cnt"] = gorm.Expr("remain_cnt - ?", 1)
		mapUpdates["uniid_list"] = newUniidList
		err = cli.Table(course_package_tableName).Model(&model.CoursePackageModel{}).Where("package_id = ?", packageId).Updates(mapUpdates).Error
		if err != nil {
			fmt.Printf("update remain_cnt err, err:%+v packageId:%s mapUpdates:%+v\n", err, packageId, mapUpdates)
			tx.Rollback()
			return err
		}

		return nil
	})
	return err
}

func (imp *CoursePackageInterfaceImp) AddCourseCnt(packageId string, cnt int) error {
	cli := db.Get()

	err := cli.Transaction(func(tx *gorm.DB) error {
		err := cli.Table(course_package_tableName).Model(&model.CoursePackageModel{}).Where("package_id = ?", packageId).UpdateColumn("remain_cnt", gorm.Expr("remain_cnt + ?", cnt)).Error
		if err != nil {
			fmt.Printf("update remain_cnt err, packageId:%s cnt:%d\n", packageId, cnt)
			tx.Rollback()
			return err
		}

		err = cli.Table(course_package_tableName).Model(&model.CoursePackageModel{}).Where("package_id = ?", packageId).UpdateColumn("total_cnt", gorm.Expr("total_cnt + ?", cnt)).Error
		if err != nil {
			fmt.Printf("update total_cnt err, packageId:%s cnt:%d\n", packageId, cnt)
			tx.Rollback()
			return err
		}
		return nil
	})
	return err
}

func (imp *CoursePackageInterfaceImp) AddRemainCourseCnt(packageId string, cnt int) error {
	cli := db.Get()
	return cli.Table(course_package_tableName).Model(&model.CoursePackageModel{}).Where("package_id = ?", packageId).UpdateColumn("remain_cnt", gorm.Expr("remain_cnt + ?", cnt)).Error
}

func (imp *CoursePackageInterfaceImp) UpdateCoursePackage(uid int64, packageId string, mapUpdates map[string]interface{}) error {
	cli := db.Get()
	return cli.Table(course_package_tableName).Model(&model.CoursePackageModel{}).Where("uid = ? AND package_id = ?",
		uid, packageId).Updates(mapUpdates).Error
}

func (imp *CoursePackageInterfaceImp) GetAllCoursePackageList(ts int64) ([]model.CoursePackageModel, error) {
	cli := db.Get()
	var vecCoursePackageModel []model.CoursePackageModel
	var err error
	if ts != 0 {
		err = cli.Raw("SELECT * FROM course_packages WHERE ts < ? ORDER BY ts DESC Limit 500", ts).Scan(&vecCoursePackageModel).Error
	} else {
		err = cli.Raw("SELECT * FROM course_packages ORDER BY ts DESC Limit 500").Scan(&vecCoursePackageModel).Error
	}
	return vecCoursePackageModel, err
}

func (imp *CoursePackageInterfaceImp) GetAllTrailCoursePackageList(ts int64) ([]model.CoursePackageModel, error) {
	cli := db.Get()
	var vecCoursePackageModel []model.CoursePackageModel
	var err error
	if ts != 0 {
		err = cli.Raw("SELECT * FROM course_packages WHERE package_type = 1 AND ts < ? ORDER BY ts DESC Limit 500", ts).Scan(&vecCoursePackageModel).Error
	} else {
		err = cli.Raw("SELECT * FROM course_packages WHERE package_type = 1 ORDER BY ts DESC Limit 500").Scan(&vecCoursePackageModel).Error
	}
	return vecCoursePackageModel, err
}

const course_package_single_lesson_tableName = "course_package_single_lessons"

func (imp *CoursePackageSingleLessonInterfaceImp) GetSingleLessonListByPackageId(uid int64, packageId string) ([]model.CoursePackageSingleLessonModel, error) {
	var err error
	var vecCoursePackageSingleLessonModel []model.CoursePackageSingleLessonModel
	cli := db.Get()
	err = cli.Table(course_package_single_lesson_tableName).Where("uid = ? AND package_id = ?", uid, packageId).Order("schedule_beg_ts DESC, create_ts DESC").Find(&vecCoursePackageSingleLessonModel).Error
	return vecCoursePackageSingleLessonModel, err
}

func (imp *CoursePackageSingleLessonInterfaceImp) GetSingleLessonByAppointmentId(uid int64, appointmentID int) ([]model.CoursePackageSingleLessonModel, error) {
	var err error
	var lessons []model.CoursePackageSingleLessonModel
	cli := db.Get()
	err = cli.Table(course_package_single_lesson_tableName).Select("lesson_id, package_id, appointment_id, scheduled_by_coach, status, create_ts, is_confirm").
		Where("uid = ? AND appointment_id = ?", uid, appointmentID).Find(&lessons).Error
	return lessons, err
}

func (imp *CoursePackageSingleLessonInterfaceImp) AddSingleLesson2Package(stCoursePackageSingleLessonModel *model.CoursePackageSingleLessonModel) error {
	cli := db.Get()
	return cli.Table(course_package_single_lesson_tableName).Save(stCoursePackageSingleLessonModel).Error
}

func (imp *CoursePackageSingleLessonInterfaceImp) GetSingleLessonById(uid int64, lessonId string) (*model.CoursePackageSingleLessonModel, error) {
	var err error
	var lesson = new(model.CoursePackageSingleLessonModel)
	cli := db.Get()
	err = cli.Table(course_package_single_lesson_tableName).Where("uid = ? AND lesson_id = ?", uid, lessonId).First(lesson).Error
	return lesson, err
}

func (imp *CoursePackageSingleLessonInterfaceImp) UpdateSingleLesson(uid int64, lessonId string, mapUpdates map[string]interface{}) error {
	cli := db.Get()
	return cli.Table(course_package_single_lesson_tableName).Model(&model.CoursePackageSingleLessonModel{}).
		Where("uid = ? AND lesson_id = ?", uid, lessonId).Updates(mapUpdates).Error
}

func (imp *CoursePackageSingleLessonInterfaceImp) GetSingleLessonListNotFinish(nowTs int64, limit int) ([]model.CoursePackageSingleLessonModel, error) {
	var err error
	var vecCoursePackageSingleLessonModel []model.CoursePackageSingleLessonModel
	cli := db.Get()

	//如果当前时间已经超过了课程终止时间，还没有核销，那么则认为用户旷课，或者是教练忘记核销了
	err = cli.Table(course_package_single_lesson_tableName).Where("status = ? AND schedule_end_ts < ? ", model.En_LessonStatus_Scheduled, nowTs).Find(&vecCoursePackageSingleLessonModel).Limit(limit).Error
	return vecCoursePackageSingleLessonModel, err
}

func (imp *CoursePackageSingleLessonInterfaceImp) GetSingleLessonListFinishNotSendMsgWriteComment(nowTs int64, limit int) ([]model.CoursePackageSingleLessonModel, error) {
	var err error
	var vecCoursePackageSingleLessonModel []model.CoursePackageSingleLessonModel
	cli := db.Get()
	err = cli.Table(course_package_single_lesson_tableName).Where("status = ? AND schedule_end_ts < ?  AND send_msg_write_comment = false", model.En_LessonStatusCompleted, nowTs).Find(&vecCoursePackageSingleLessonModel).Limit(limit).Error
	return vecCoursePackageSingleLessonModel, err
}

func (imp *CoursePackageSingleLessonInterfaceImp) GetTodaySingleLessonListNotSendMsgGoLesson(ts int64, limit int) ([]model.CoursePackageSingleLessonModel, error) {
	var err error
	var vecCoursePackageSingleLessonModel []model.CoursePackageSingleLessonModel
	cli := db.Get()

	// schedule_beg_ts 过滤出距离开课前2小时的课程
	// schedule_beg_ts <= 当天凌晨24点，过滤出今天的课程
	err = cli.Table(course_package_single_lesson_tableName).Where("status = ? AND send_msg_go_lesson = false AND schedule_beg_ts <= ?",
		model.En_LessonStatus_Scheduled, ts).Find(&vecCoursePackageSingleLessonModel).Limit(limit).Error
	return vecCoursePackageSingleLessonModel, err
}

func (imp *CoursePackageSingleLessonInterfaceImp) GetSingleLessonListMissed(limit int) ([]model.CoursePackageSingleLessonModel, error) {
	var err error
	var vecCoursePackageSingleLessonModel []model.CoursePackageSingleLessonModel
	cli := db.Get()
	err = cli.Table(course_package_single_lesson_tableName).Where("status = ? AND write_off_missed_return_cnt = false", model.En_LessonStatusMissed).
		Find(&vecCoursePackageSingleLessonModel).Limit(limit).Error
	return vecCoursePackageSingleLessonModel, err
}

func (imp *CoursePackageSingleLessonInterfaceImp) GetCompletedSingleLessonListByCoachId(coachId int, uBegTs int64) ([]model.CoursePackageSingleLessonModel, error) {
	var err error
	var vecCoursePackageSingleLessonModel []model.CoursePackageSingleLessonModel
	cli := db.Get()
	err = cli.Table(course_package_single_lesson_tableName).Where("coach_id = ? AND schedule_beg_ts > ? AND status = ?",
		coachId, uBegTs, model.En_LessonStatusCompleted).Find(&vecCoursePackageSingleLessonModel).Error
	return vecCoursePackageSingleLessonModel, err
}

func (imp *CoursePackageSingleLessonInterfaceImp) GetAllSingleLessonList(createTs int64) ([]model.CoursePackageSingleLessonModel, error) {
	cli := db.Get()
	var vecRes []model.CoursePackageSingleLessonModel
	var err error
	if createTs != 0 {
		err = cli.Raw("SELECT * FROM course_package_single_lessons WHERE create_ts < ? ORDER BY create_ts DESC Limit 500", createTs).Scan(&vecRes).Error
	} else {
		err = cli.Raw("SELECT * FROM course_package_single_lessons ORDER BY create_ts DESC Limit 500").Scan(&vecRes).Error
	}
	return vecRes, err
}

const payment_order_tableName = "payment_orders"

func (imp *PaymentOrderInterfaceImp) UpdateOrderSucc(orderId string, uid int64, mapUpdates map[string]interface{}) error {
	cli := db.Get()
	return cli.Table(payment_order_tableName).Model(&model.PaymentOrderModel{}).Where("order_id = ? AND payer_uid = ?",
		orderId, uid).Updates(mapUpdates).Error
}

func (imp *PaymentOrderInterfaceImp) GetOrderById(orderId string, uid int64) (*model.PaymentOrderModel, error) {
	var err error
	var code = new(model.PaymentOrderModel)
	cli := db.Get()
	err = cli.Table(payment_order_tableName).Where("order_id = ? AND payer_uid = ?", orderId, uid).First(code).Error
	return code, err
}

func (imp *PaymentOrderInterfaceImp) GetOrder(orderId string) (*model.PaymentOrderModel, error) {
	var err error
	var order = new(model.PaymentOrderModel)
	cli := db.Get()
	err = cli.Table(payment_order_tableName).Where("order_id = ?", orderId).First(order).Error
	return order, err
}

func (imp *PaymentOrderInterfaceImp) AddOrder(stPaymentOrderModel model.PaymentOrderModel) error {
	cli := db.Get()
	return cli.Table(payment_order_tableName).Save(stPaymentOrderModel).Error
}

func (imp *PaymentOrderInterfaceImp) GetOrderList(uid int64) ([]model.PaymentOrderModel, error) {
	var err error
	var vecPaymentOrderModel []model.PaymentOrderModel
	cli := db.Get()
	err = cli.Table(payment_order_tableName).Where("payer_uid = ?", uid).Order("order_time DESC").Find(&vecPaymentOrderModel).Error
	return vecPaymentOrderModel, err
}

func (imp *PaymentOrderInterfaceImp) GetOrderListByCoachId(coachId int, begTs int64) ([]model.PaymentOrderModel, error) {
	var err error
	var vecPaymentOrderModel []model.PaymentOrderModel
	cli := db.Get()
	err = cli.Table(payment_order_tableName).Where("coach_id = ? AND payment_time > ?", coachId, begTs).Find(&vecPaymentOrderModel).Error
	return vecPaymentOrderModel, err
}

func (imp *PaymentOrderInterfaceImp) GetRefundOrderByPackageId(uid int64, packageId string) (*model.PaymentOrderModel, error) {
	var err error
	var order = new(model.PaymentOrderModel)
	cli := db.Get()
	err = cli.Table(payment_order_tableName).Where("payer_uid = ? AND package_id = ? AND order_status = ?",
		uid, packageId, model.Enum_Pay_Status_Refunded).First(order).Error
	return order, err
}

func (imp *PaymentOrderInterfaceImp) GetOrderByPackageId(uid int64, packageId string) ([]model.PaymentOrderModel, error) {
	var err error
	var vecPaymentOrderModel []model.PaymentOrderModel
	cli := db.Get()
	err = cli.Table(payment_order_tableName).Where("payer_uid = ? AND package_id = ?", uid, packageId).Find(&vecPaymentOrderModel).Error
	return vecPaymentOrderModel, err
}

const coach_appointments_tableName = "coach_appointments"

func (imp *AppointmentInterfaceImp) SetAppointmentSchedule(stCoachAppointmentModel model.CoachAppointmentModel) error {
	cli := db.Get()
	return cli.Table(coach_appointments_tableName).Save(stCoachAppointmentModel).Error
}

func (imp *AppointmentInterfaceImp) GetAppointmentScheduleHasUidFromBegTs(gymId int, coachId int, dayBegTs int64) ([]model.CoachAppointmentModel, error) {
	var err error
	var vecCoachAppointmentModel []model.CoachAppointmentModel
	cli := db.Get()
	//err = cli.Table(coach_appointments_tableName).Where("coach_id = ? AND gym_id = ? AND appointment_date >= ?", coachId, gymId, dayBegTs).Order("start_time ASC").Find(&vecCoachAppointmentModel).Error
	err = cli.Raw("SELECT * FROM coach_appointments WHERE coach_id = ? AND gym_id = ? AND appointment_date >= ? AND user_id > 0 ORDER BY start_time ASC Limit 100",
		coachId, gymId, dayBegTs).Scan(&vecCoachAppointmentModel).Error
	return vecCoachAppointmentModel, err
}

func (imp *AppointmentInterfaceImp) GetAppointmentScheduleFromBegTs(gymId int, coachId int, dayBegTs int64) ([]model.CoachAppointmentModel, error) {
	var err error
	var vecCoachAppointmentModel []model.CoachAppointmentModel
	cli := db.Get()
	err = cli.Table(coach_appointments_tableName).Where("coach_id = ? AND gym_id = ? AND appointment_date >= ?", coachId, gymId, dayBegTs).Order("start_time ASC").Find(&vecCoachAppointmentModel).Error
	return vecCoachAppointmentModel, err
}

func (imp *AppointmentInterfaceImp) GetAppointmentScheduleOneDay(gymId int, coachId int, dayBegTs int64) ([]model.CoachAppointmentModel, error) {
	var err error
	var vecCoachAppointmentModel []model.CoachAppointmentModel
	cli := db.Get()
	err = cli.Table(coach_appointments_tableName).Where("coach_id = ? AND gym_id = ? AND appointment_date = ?", coachId, gymId, dayBegTs).Order("start_time ASC").Find(&vecCoachAppointmentModel).Error
	return vecCoachAppointmentModel, err
}

func (imp *AppointmentInterfaceImp) DelAppointmentByCoach(appointmentID int, coachId int) error {
	var err error
	cli := db.Get()
	err = cli.Table(coach_appointments_tableName).Where("appointment_id = ? AND coach_id = ?", appointmentID, coachId).Delete(model.CoachAppointmentModel{}).Error
	return err
}

func (imp *AppointmentInterfaceImp) GetUserAppointmentRecordOneDay(uid int64, dayBegTs int64) ([]model.CoachAppointmentModel, error) {
	var err error
	var vecCoachAppointmentModel []model.CoachAppointmentModel
	cli := db.Get()
	err = cli.Table(coach_appointments_tableName).Where("user_id = ? AND appointment_date = ?", uid, dayBegTs).Order("start_time DESC").Find(&vecCoachAppointmentModel).Error
	return vecCoachAppointmentModel, err
}

func (imp *AppointmentInterfaceImp) GetUserAppointmentRecordFromBegTs(uid int64, dayBegTs int64, limit int) ([]model.CoachAppointmentModel, error) {
	var err error
	var vecCoachAppointmentModel []model.CoachAppointmentModel
	cli := db.Get()
	err = cli.Raw("SELECT * FROM coach_appointments WHERE user_id = ? AND appointment_date >= ? ORDER BY appointment_date DESC Limit ?",
		uid, dayBegTs, limit).Scan(&vecCoachAppointmentModel).Error
	//err = cli.Table(coach_appointments_tableName).Where("user_id = ? AND appointment_date >= ?", uid, dayBegTs).Order("appointment_date DESC").Limit(limit).Find(&vecCoachAppointmentModel).Error
	return vecCoachAppointmentModel, err
}

func (imp *AppointmentInterfaceImp) GetUserAppointmentRecord(uid int64, limit int) ([]model.CoachAppointmentModel, error) {
	var err error
	var vecCoachAppointmentModel []model.CoachAppointmentModel
	cli := db.Get()
	err = cli.Table(coach_appointments_tableName).Where("user_id = ?", uid).Order("appointment_date DESC").Limit(limit).Find(&vecCoachAppointmentModel).Error
	return vecCoachAppointmentModel, err
}

func (imp *AppointmentInterfaceImp) GetAppointmentById(appointmentID int) (*model.CoachAppointmentModel, error) {
	var err error
	var appointment = new(model.CoachAppointmentModel)
	cli := db.Get()
	err = cli.Table(coach_appointments_tableName).Where("appointment_id = ?", appointmentID).First(appointment).Error
	return appointment, err
}

func (imp *AppointmentInterfaceImp) GetAppointmentByBegTsAndEndTs(gymId int, coachid int, begTs int64, endTs int64) (*model.CoachAppointmentModel, error) {
	var err error
	var appointment = new(model.CoachAppointmentModel)
	cli := db.Get()
	err = cli.Table(coach_appointments_tableName).Where("coach_id = ? AND gym_id = ? AND start_time = ? AND end_time = ?",
		coachid, gymId, begTs, endTs).First(appointment).Error
	return appointment, err
}

func (imp *AppointmentInterfaceImp) SetAppointmentBooked(uid int64, appointmentID int, courseId int) (error, model.CoachAppointmentModel) {
	var err error
	cli := db.Get().Table(coach_appointments_tableName)

	// 先获取再更新的原子操作
	var stCoachAppointmentModel model.CoachAppointmentModel
	err = cli.Transaction(func(tx *gorm.DB) error {
		// 获取用户记录
		if err := tx.First(&stCoachAppointmentModel, "appointment_id = ?", appointmentID).Error; err != nil {
			fmt.Printf("get err, uid:%d appointmentID:%d\n", uid, appointmentID)
			return err
		}

		if stCoachAppointmentModel.Status != model.Enum_Appointment_Status_Available {
			return errors.New("book status unavailable")
		}

		// 更新用户记录
		mapUpdates := map[string]interface{}{}
		mapUpdates["status"] = model.Enum_Appointment_Status_UnAvailable
		mapUpdates["user_id"] = uid
		mapUpdates["user_course_id"] = courseId
		mapUpdates["update_ts"] = time.Now().Unix()
		stCoachAppointmentModel.Status = model.Enum_Appointment_Status_UnAvailable
		stCoachAppointmentModel.UserID = uid
		stCoachAppointmentModel.UserCourseID = courseId
		stCoachAppointmentModel.UpdateTs = time.Now().Unix()

		// 更新用户数据，使用 Update 方法
		if err := tx.Model(&model.CoachAppointmentModel{}).Where("appointment_id = ?", appointmentID).Updates(mapUpdates).Error; err != nil {
			fmt.Printf("update err, uid:%d appointmentID:%d mapUpdates:%+v\n", uid, appointmentID, mapUpdates)
			tx.Rollback()
			return err
		}

		return nil
	})
	return err, stCoachAppointmentModel
}

// 用户取消约课，将课程变回可用状态，即所有用户都可预约
func (imp *AppointmentInterfaceImp) CancelAppointmentBooked(uid int64, lessonID string, appointmentID int) error {
	var err error
	cli := db.Get().Table(coach_appointments_tableName)

	// 先获取再更新的原子操作
	err = cli.Transaction(func(tx *gorm.DB) error {
		var stCoachAppointmentModel model.CoachAppointmentModel
		// 获取用户记录
		if err := tx.First(&stCoachAppointmentModel, "appointment_id = ?", appointmentID).Error; err != nil {
			fmt.Printf("get err, uid:%d appointmentID:%d\n", uid, appointmentID)
			return err
		}

		if stCoachAppointmentModel.Status == model.Enum_Appointment_Status_Available {
			return errors.New("book already been cancel")
		}

		// 更新用户记录
		mapUpdates := map[string]interface{}{}
		mapUpdates["status"] = model.Enum_Appointment_Status_Available
		mapUpdates["user_id"] = 0
		mapUpdates["user_course_id"] = 0
		mapUpdates["canceled_course"] = lessonID
		mapUpdates["update_ts"] = time.Now().Unix()
		stCoachAppointmentModel.Status = model.Enum_Appointment_Status_Available
		stCoachAppointmentModel.UserID = 0
		stCoachAppointmentModel.UserCourseID = 0
		stCoachAppointmentModel.UpdateTs = time.Now().Unix()

		// 更新用户数据，使用 Update 方法
		if err := tx.Model(&model.CoachAppointmentModel{}).Where("appointment_id = ?", appointmentID).Updates(mapUpdates).Error; err != nil {
			fmt.Printf("update err, uid:%d appointmentID:%d mapUpdates:%+v\n", uid, appointmentID, mapUpdates)
			tx.Rollback()
			return err
		}

		return nil
	})
	return err
}

const invitation_code_tableName = "invitation_code"

func (imp *InvitationCodeInterfaceImp) GetCode(code string) (*model.InvitationCodeModel, error) {
	var err error
	var invitationCode = new(model.InvitationCodeModel)
	cli := db.Get()
	err = cli.Table(invitation_code_tableName).Where("invitation_code = ?", code).First(invitationCode).Error
	return invitationCode, err
}

func (imp *InvitationCodeInterfaceImp) AddCode(stInvitationCodeModel *model.InvitationCodeModel) error {
	cli := db.Get()
	return cli.Table(invitation_code_tableName).Save(stInvitationCodeModel).Error
}

func (imp *InvitationCodeInterfaceImp) UpdateCode(code string, uid int64) error {
	var err error
	cli := db.Get().Table(invitation_code_tableName)

	// 先获取再更新的原子操作
	err = cli.Transaction(func(tx *gorm.DB) error {
		var stInvitationCodeModel model.InvitationCodeModel

		// 获取用户记录
		if err := tx.First(&stInvitationCodeModel, "invitation_code = ?", code).Error; err != nil {
			fmt.Printf("get err, uid:%d code:%s\n", uid, code)
			return err
		}

		if stInvitationCodeModel.IsUsed == true {
			return errors.New("code has used")
		}

		// 更新用户记录
		mapUpdates := map[string]interface{}{}
		mapUpdates["is_used"] = true
		mapUpdates["used_uid"] = uid
		mapUpdates["used_ts"] = time.Now().Unix()

		// 更新用户数据，使用 Update 方法
		if err := tx.Model(&model.InvitationCodeModel{}).Where("invitation_code = ?", code).Updates(mapUpdates).Error; err != nil {
			fmt.Printf("update err, uid:%d code:%s mapUpdates:%+v\n", uid, code, mapUpdates)
			tx.Rollback()
			return err
		}

		return nil
	})
	return err
}

func (imp *InvitationCodeInterfaceImp) GetCount() (int64, error) {
	cli := db.Get()
	var count int64
	err := cli.Table(invitation_code_tableName).Model(&model.InvitationCodeModel{}).Count(&count).Error
	return count, err
}

const coach_page_banner_tableName = "coach_page_banner"

func (imp *CoachPageBannerInterfaceImp) GetBannerList() ([]model.CoachPageBannerModel, error) {
	var err error
	var vecCoachPageBannerModel []model.CoachPageBannerModel
	cli := db.Get()
	err = cli.Table(coach_page_banner_tableName).Order("priority DESC").Find(&vecCoachPageBannerModel).Error
	return vecCoachPageBannerModel, err
}

func getTodayEndTs() int64 {
	// 获取当前时间
	now := time.Now()

	// 创建当天零点的时间
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// 加一天，得到次日零点的时间
	nextMidnight := midnight.Add(24 * time.Hour)

	// 返回次日零点的时间戳
	return nextMidnight.Unix()
}
