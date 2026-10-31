package dao

import (
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

// AddCoursePackageRefundRecord 新增退费或课时归零历史记录。
func (imp *CoursePackageRefundRecordInterfaceImp) AddCoursePackageRefundRecord(record *model.CoursePackageRefundRecordModel) error {
	return db.Get().Table("course_package_refund_record").Create(record).Error
}

// GetCoursePackageRefundRecordList 按 ID 倒序游标分页查询全量历史。
func (imp *CoursePackageRefundRecordInterfaceImp) GetCoursePackageRefundRecordList(lastID int64, limit int) ([]model.CoursePackageRefundRecordModel, error) {
	var records []model.CoursePackageRefundRecordModel
	var err error
	if lastID > 0 {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_refund_record
			WHERE id < ?
			ORDER BY id DESC
			LIMIT ?`, lastID, limit).Scan(&records).Error
	} else {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_refund_record
			ORDER BY id DESC
			LIMIT ?`, limit).Scan(&records).Error
	}
	return records, err
}

// GetCoursePackageRefundRecordListByUid 按用户查询历史，按 ID 倒序游标分页。
func (imp *CoursePackageRefundRecordInterfaceImp) GetCoursePackageRefundRecordListByUid(uid, lastID int64, limit int) ([]model.CoursePackageRefundRecordModel, error) {
	var records []model.CoursePackageRefundRecordModel
	var err error
	if lastID > 0 {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_refund_record
			WHERE uid = ? AND id < ?
			ORDER BY id DESC
			LIMIT ?`, uid, lastID, limit).Scan(&records).Error
	} else {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_refund_record
			WHERE uid = ?
			ORDER BY id DESC
			LIMIT ?`, uid, limit).Scan(&records).Error
	}
	return records, err
}

// GetCoursePackageRefundRecordListByPackageId 按课包查询历史，按 ID 倒序游标分页。
func (imp *CoursePackageRefundRecordInterfaceImp) GetCoursePackageRefundRecordListByPackageId(packageID string, lastID int64, limit int) ([]model.CoursePackageRefundRecordModel, error) {
	var records []model.CoursePackageRefundRecordModel
	var err error
	if lastID > 0 {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_refund_record
			WHERE package_id = ? AND id < ?
			ORDER BY id DESC
			LIMIT ?`, packageID, lastID, limit).Scan(&records).Error
	} else {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_refund_record
			WHERE package_id = ?
			ORDER BY id DESC
			LIMIT ?`, packageID, limit).Scan(&records).Error
	}
	return records, err
}
