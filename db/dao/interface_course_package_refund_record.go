package dao

import "github.com/xionghengheng/ff_plib/db/model"

// CoursePackageRefundRecordInterface 课包退费及课时归零历史数据访问接口。
// 参数有效性、权限、课包归属及重复请求控制由调用方负责。
type CoursePackageRefundRecordInterface interface {
	// AddCoursePackageRefundRecord 新增历史记录，自增 ID 回填到 record。
	// 时间字段由调用方提供；本方法不更新课包、不触发退款。
	// 与现有课包更新接口先后调用不具备跨调用事务保证。
	AddCoursePackageRefundRecord(record *model.CoursePackageRefundRecordModel) error

	// GetCoursePackageRefundRecordList 按 ID 倒序游标分页查询全量历史。
	// lastID <= 0 取首页，后续传上一页最后一条记录的 ID；limit 必须大于 0。
	GetCoursePackageRefundRecordList(lastID int64, limit int) ([]model.CoursePackageRefundRecordModel, error)

	// GetCoursePackageRefundRecordListByUid 按用户查询历史，按 ID 倒序游标分页。
	// lastID <= 0 取首页，后续传上一页最后一条记录的 ID；limit 必须大于 0。
	GetCoursePackageRefundRecordListByUid(uid, lastID int64, limit int) ([]model.CoursePackageRefundRecordModel, error)

	// GetCoursePackageRefundRecordListByPackageId 按课包查询历史，按 ID 倒序游标分页。
	// lastID <= 0 取首页，后续传上一页最后一条记录的 ID；limit 必须大于 0。
	GetCoursePackageRefundRecordListByPackageId(packageID string, lastID int64, limit int) ([]model.CoursePackageRefundRecordModel, error)
}

// CoursePackageRefundRecordInterfaceImp 课包退费历史数据访问实现。
type CoursePackageRefundRecordInterfaceImp struct{}

// ImpCoursePackageRefundRecord 课包退费历史 DAO 实例。
var ImpCoursePackageRefundRecord CoursePackageRefundRecordInterface = &CoursePackageRefundRecordInterfaceImp{}

var _ CoursePackageRefundRecordInterface = (*CoursePackageRefundRecordInterfaceImp)(nil)
