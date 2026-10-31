package dao

import "github.com/xionghengheng/ff_plib/db/model"

// CoursePackageCoachChangeRecordInterface 课包更换教练操作记录数据访问接口。
// 参数有效性、权限、课包归属、教练存在性及重复请求控制由调用方负责。
type CoursePackageCoachChangeRecordInterface interface {
	// AddCoursePackageCoachChangeRecord 新增操作记录，自增 ID 回填到 record。
	// 原教练 ID 和时间字段由调用方提供；本方法不修改课包。
	// 与现有课包更新接口先后调用不具备跨调用事务保证。
	AddCoursePackageCoachChangeRecord(record *model.CoursePackageCoachChangeRecordModel) error

	// GetCoursePackageCoachChangeRecordList 按变更时间倒序分页查询全量记录，同一时间按 ID 倒序。
	// 首页传 (0, 0)，后续传上一页最后一条的 ChangeCoachTs、ID；limit 必须大于 0。
	GetCoursePackageCoachChangeRecordList(lastChangeCoachTs, lastID int64, limit int) ([]model.CoursePackageCoachChangeRecordModel, error)
}

// CoursePackageCoachChangeRecordInterfaceImp 课包更换教练操作记录数据访问实现。
type CoursePackageCoachChangeRecordInterfaceImp struct{}

// ImpCoursePackageCoachChangeRecord 课包更换教练操作记录 DAO 实例。
var ImpCoursePackageCoachChangeRecord CoursePackageCoachChangeRecordInterface = &CoursePackageCoachChangeRecordInterfaceImp{}

var _ CoursePackageCoachChangeRecordInterface = (*CoursePackageCoachChangeRecordInterfaceImp)(nil)
