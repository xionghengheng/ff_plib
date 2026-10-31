package dao

import (
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

// AddCoursePackageCoachChangeRecord 新增课包更换教练操作记录。
func (imp *CoursePackageCoachChangeRecordInterfaceImp) AddCoursePackageCoachChangeRecord(record *model.CoursePackageCoachChangeRecordModel) error {
	return db.Get().Table("course_package_coach_change_record").Create(record).Error
}

// GetCoursePackageCoachChangeRecordList 按变更时间、ID 倒序游标分页查询全量记录。
func (imp *CoursePackageCoachChangeRecordInterfaceImp) GetCoursePackageCoachChangeRecordList(lastChangeCoachTs, lastID int64, limit int) ([]model.CoursePackageCoachChangeRecordModel, error) {
	var records []model.CoursePackageCoachChangeRecordModel
	var err error
	if lastID > 0 {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_coach_change_record
			WHERE change_coach_ts < ? OR (change_coach_ts = ? AND id < ?)
			ORDER BY change_coach_ts DESC, id DESC
			LIMIT ?`, lastChangeCoachTs, lastChangeCoachTs, lastID, limit).Scan(&records).Error
	} else {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_coach_change_record
			ORDER BY change_coach_ts DESC, id DESC
			LIMIT ?`, limit).Scan(&records).Error
	}
	return records, err
}
