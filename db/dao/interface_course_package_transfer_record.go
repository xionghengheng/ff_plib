package dao

import "github.com/xionghengheng/ff_plib/db/model"

// CoursePackageTransferRecordInterface 仅负责转移历史的写入和查询，不修改课包或支付订单。
type CoursePackageTransferRecordInterface interface {
	// AddCoursePackageTransferRecord 新增不可覆盖的成功记录，自增 ID 回填至 record。
	// 调用方在课包和支付记录更新成功后调用；不具有跨调用事务保证。
	AddCoursePackageTransferRecord(record *model.CoursePackageTransferRecordModel) error
	// GetCoursePackageTransferRecordList 按 transfer_ts、id 倒序分页。
	// 首页游标为 (0, 0)，后续传上一页末条记录的时间和 ID。
	GetCoursePackageTransferRecordList(lastTransferTs, lastID int64, limit int) ([]model.CoursePackageTransferRecordModel, error)
}

type CoursePackageTransferRecordInterfaceImp struct{}

var ImpCoursePackageTransferRecord CoursePackageTransferRecordInterface = &CoursePackageTransferRecordInterfaceImp{}

var _ CoursePackageTransferRecordInterface = (*CoursePackageTransferRecordInterfaceImp)(nil)
