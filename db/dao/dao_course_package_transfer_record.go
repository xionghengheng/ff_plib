package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

const coursePackageTransferRecordTable = "course_package_transfer_record"

func (imp *CoursePackageTransferRecordInterfaceImp) AddCoursePackageTransferRecord(record *model.CoursePackageTransferRecordModel) error {
	return addCoursePackageTransferRecord(db.Get(), record)
}

func addCoursePackageTransferRecord(conn *gorm.DB, record *model.CoursePackageTransferRecordModel) error {
	return conn.Table(coursePackageTransferRecordTable).Create(record).Error
}

func (imp *CoursePackageTransferRecordInterfaceImp) GetCoursePackageTransferRecordList(lastTransferTs, lastID int64, limit int) ([]model.CoursePackageTransferRecordModel, error) {
	return getCoursePackageTransferRecordList(db.Get(), lastTransferTs, lastID, limit)
}

func getCoursePackageTransferRecordList(conn *gorm.DB, lastTransferTs, lastID int64, limit int) ([]model.CoursePackageTransferRecordModel, error) {
	records := make([]model.CoursePackageTransferRecordModel, 0)
	query := conn.Table(coursePackageTransferRecordTable)
	if lastID > 0 {
		query = query.Where("transfer_ts < ? OR (transfer_ts = ? AND id < ?)", lastTransferTs, lastTransferTs, lastID)
	}
	err := query.Order("transfer_ts DESC, id DESC").Limit(limit).Find(&records).Error
	return records, err
}
