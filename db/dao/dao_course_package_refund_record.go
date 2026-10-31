package dao

import (
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

// AddCoursePackageRefundRecord 新增退费或课时归零历史记录。
func (imp *CoursePackageRefundRecordInterfaceImp) AddCoursePackageRefundRecord(record *model.CoursePackageRefundRecordModel) error {
	return db.Get().Table("course_package_refund_record").Create(record).Error
}

// GetCoursePackageRefundRecordList 按退费时间、ID 倒序游标分页查询全量历史。
func (imp *CoursePackageRefundRecordInterfaceImp) GetCoursePackageRefundRecordList(lastRefundTs, lastID int64, limit int) ([]model.CoursePackageRefundRecordModel, error) {
	var records []model.CoursePackageRefundRecordModel
	var err error
	if lastID > 0 {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_refund_record
			WHERE refund_ts < ? OR (refund_ts = ? AND id < ?)
			ORDER BY refund_ts DESC, id DESC
			LIMIT ?`, lastRefundTs, lastRefundTs, lastID, limit).Scan(&records).Error
	} else {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_refund_record
			ORDER BY refund_ts DESC, id DESC
			LIMIT ?`, limit).Scan(&records).Error
	}
	return records, err
}
