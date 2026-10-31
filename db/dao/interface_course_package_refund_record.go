package dao

import "github.com/xionghengheng/ff_plib/db/model"

// CoursePackageRefundRecordInterface 课包退费及课时归零历史数据访问接口。
// 参数有效性、权限、课包归属及重复请求控制由调用方负责。
type CoursePackageRefundRecordInterface interface {
	// AddCoursePackageRefundRecord 新增历史记录，自增 ID 回填到 record。
	// 时间字段由调用方提供；本方法不更新课包、不触发退款。
	// 与现有课包更新接口先后调用不具备跨调用事务保证。
	AddCoursePackageRefundRecord(record *model.CoursePackageRefundRecordModel) error

	// GetCoursePackageRefundRecordList 按退费时间倒序分页查询全量历史，同一时间按 ID 倒序。
	// 首页传 (0, 0)，后续传上一页最后一条的 RefundTs、ID；limit 必须大于 0。
	GetCoursePackageRefundRecordList(lastRefundTs, lastID int64, limit int) ([]model.CoursePackageRefundRecordModel, error)
}

// CoursePackageRefundRecordInterfaceImp 课包退费历史数据访问实现。
type CoursePackageRefundRecordInterfaceImp struct{}

// ImpCoursePackageRefundRecord 课包退费历史 DAO 实例。
var ImpCoursePackageRefundRecord CoursePackageRefundRecordInterface = &CoursePackageRefundRecordInterfaceImp{}

var _ CoursePackageRefundRecordInterface = (*CoursePackageRefundRecordInterfaceImp)(nil)
