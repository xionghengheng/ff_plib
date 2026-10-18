package dao

import "github.com/xionghengheng/ff_plib/db/model"

// CoursePackagePurchaseWhitelistInterface 课包购买最低次数白名单数据访问接口。
// 参数有效性及新增授权的初始状态由调用方校验。
type CoursePackagePurchaseWhitelistInterface interface {
	// AddCoursePackagePurchaseWhitelist 新增未使用的授权，自增 ID 回填到 record。
	// 是否已有未使用授权及是否允许新增，由业务上层查询后决定。
	AddCoursePackagePurchaseWhitelist(record *model.CoursePackagePurchaseWhitelistModel) error

	// GetUnusedCoursePackagePurchaseWhitelistByUid 查询用户当前未使用的授权。
	// 未找到时返回 gorm.ErrRecordNotFound。
	GetUnusedCoursePackagePurchaseWhitelistByUid(uid int64) (*model.CoursePackagePurchaseWhitelistModel, error)

	// GetCoursePackagePurchaseWhitelistListByUid 按 ID 倒序游标分页查询用户授权历史。
	// lastID <= 0 取首页，后续传上一页最后一条记录的 ID；limit 必须大于 0。
	GetCoursePackagePurchaseWhitelistListByUid(uid, lastID int64, limit int) ([]model.CoursePackagePurchaseWhitelistModel, error)

	// UpdateCoursePackagePurchaseWhitelist 根据主键 ID 更新指定字段。
	// 支付成功后，由业务上层更新使用状态、订单 ID 及使用时间。
	UpdateCoursePackagePurchaseWhitelist(id int64, updates map[string]interface{}) error
}

// CoursePackagePurchaseWhitelistInterfaceImp 课包购买最低次数白名单数据访问实现。
type CoursePackagePurchaseWhitelistInterfaceImp struct{}

// ImpCoursePackagePurchaseWhitelist 课包购买最低次数白名单 DAO 实例。
var ImpCoursePackagePurchaseWhitelist CoursePackagePurchaseWhitelistInterface = &CoursePackagePurchaseWhitelistInterfaceImp{}

var _ CoursePackagePurchaseWhitelistInterface = (*CoursePackagePurchaseWhitelistInterfaceImp)(nil)
