package dao

import (
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

// AddCoursePackageCoachChangeRecord 新增课包更换教练操作记录。
func (imp *CoursePackageCoachChangeRecordInterfaceImp) AddCoursePackageCoachChangeRecord(record *model.CoursePackageCoachChangeRecordModel) error {
	return db.Get().Table("course_package_coach_change_record").Create(record).Error
}

// GetCoursePackageCoachChangeRecordList 按 ID 倒序游标分页查询全量记录。
func (imp *CoursePackageCoachChangeRecordInterfaceImp) GetCoursePackageCoachChangeRecordList(lastID int64, limit int) ([]model.CoursePackageCoachChangeRecordModel, error) {
	var records []model.CoursePackageCoachChangeRecordModel
	var err error
	if lastID > 0 {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_coach_change_record
			WHERE id < ?
			ORDER BY id DESC
			LIMIT ?`, lastID, limit).Scan(&records).Error
	} else {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_coach_change_record
			ORDER BY id DESC
			LIMIT ?`, limit).Scan(&records).Error
	}
	return records, err
}

// GetCoursePackageCoachChangeRecordListByUid 按用户查询记录，按 ID 倒序游标分页。
func (imp *CoursePackageCoachChangeRecordInterfaceImp) GetCoursePackageCoachChangeRecordListByUid(uid, lastID int64, limit int) ([]model.CoursePackageCoachChangeRecordModel, error) {
	var records []model.CoursePackageCoachChangeRecordModel
	var err error
	if lastID > 0 {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_coach_change_record
			WHERE uid = ? AND id < ?
			ORDER BY id DESC
			LIMIT ?`, uid, lastID, limit).Scan(&records).Error
	} else {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_coach_change_record
			WHERE uid = ?
			ORDER BY id DESC
			LIMIT ?`, uid, limit).Scan(&records).Error
	}
	return records, err
}

// GetCoursePackageCoachChangeRecordListByPackageId 按课包查询记录，按 ID 倒序游标分页。
func (imp *CoursePackageCoachChangeRecordInterfaceImp) GetCoursePackageCoachChangeRecordListByPackageId(packageID string, lastID int64, limit int) ([]model.CoursePackageCoachChangeRecordModel, error) {
	var records []model.CoursePackageCoachChangeRecordModel
	var err error
	if lastID > 0 {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_coach_change_record
			WHERE package_id = ? AND id < ?
			ORDER BY id DESC
			LIMIT ?`, packageID, lastID, limit).Scan(&records).Error
	} else {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_coach_change_record
			WHERE package_id = ?
			ORDER BY id DESC
			LIMIT ?`, packageID, limit).Scan(&records).Error
	}
	return records, err
}
