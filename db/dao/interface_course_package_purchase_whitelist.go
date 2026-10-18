package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/xionghengheng/ff_plib/db/model"
)

// CoursePackagePurchaseWhitelistInterface 课包购买最低次数白名单数据访问接口。
// 参数有效性及新增授权的初始状态由调用方校验。
type CoursePackagePurchaseWhitelistInterface interface {
	// AddCoursePackagePurchaseWhitelist 新增未使用的授权，自增 ID 回填到 record。
	// 同一用户已有未使用授权时，由 uk_unused_uid 唯一索引拒绝新增。
	AddCoursePackagePurchaseWhitelist(record *model.CoursePackagePurchaseWhitelistModel) error

	// GetUnusedCoursePackagePurchaseWhitelistByUid 查询用户当前未使用的授权。
	// 未找到时返回 gorm.ErrRecordNotFound。
	GetUnusedCoursePackagePurchaseWhitelistByUid(uid int64) (*model.CoursePackagePurchaseWhitelistModel, error)

	// GetCoursePackagePurchaseWhitelistListByUid 按 ID 倒序游标分页查询用户授权历史。
	// lastID <= 0 取首页，后续传上一页最后一条记录的 ID；limit 必须大于 0。
	GetCoursePackagePurchaseWhitelistListByUid(uid, lastID int64, limit int) ([]model.CoursePackagePurchaseWhitelistModel, error)

	// UseCoursePackagePurchaseWhitelist 在调用方传入的支付事务 tx 中消费授权。
	// 调用方需检查返回错误，并与订单支付、课时发放一起提交或回滚。
	// 授权不存在、不属于该用户或已使用时，返回 ErrCoursePackagePurchaseWhitelistUnavailable。
	UseCoursePackagePurchaseWhitelist(tx *gorm.DB, id, uid int64, orderID string, usedTs int64) error
}

// CoursePackagePurchaseWhitelistInterfaceImp 课包购买最低次数白名单数据访问实现。
type CoursePackagePurchaseWhitelistInterfaceImp struct{}

// ImpCoursePackagePurchaseWhitelist 课包购买最低次数白名单 DAO 实例。
var ImpCoursePackagePurchaseWhitelist CoursePackagePurchaseWhitelistInterface = &CoursePackagePurchaseWhitelistInterfaceImp{}

var _ CoursePackagePurchaseWhitelistInterface = (*CoursePackagePurchaseWhitelistInterfaceImp)(nil)
