package dao

import "github.com/xionghengheng/ff_plib/db/model"

// AfterSalesUserMaintenanceInterface 售后用户维护档案数据访问接口。
type AfterSalesUserMaintenanceInterface interface {
	// AddAfterSalesUserMaintenance 添加售后用户维护档案。
	AddAfterSalesUserMaintenance(maintenance *model.AfterSalesUserMaintenanceModel) error

	// GetAfterSalesUserMaintenanceList 按创建时间倒序游标分页查询档案列表。
	// 首页传 lastCreatedTs=0、lastID=0；后续传上一页最后一条记录的 created_ts 和 id。
	GetAfterSalesUserMaintenanceList(lastCreatedTs, lastID int64, limit int) ([]model.AfterSalesUserMaintenanceModel, error)

	// GetAfterSalesUserMaintenanceListByLatestPackagePurchaseTs 按最新课包购买时间倒序游标分页查询档案列表。
	// 首页传 lastLatestPackagePurchaseTs=0、lastID=0；后续传上一页最后一条记录的 latest_package_purchase_ts 和 id。
	GetAfterSalesUserMaintenanceListByLatestPackagePurchaseTs(lastLatestPackagePurchaseTs, lastID int64, limit int) ([]model.AfterSalesUserMaintenanceModel, error)

	// GetAfterSalesUserMaintenanceByPhone 根据手机号查询单条档案。
	GetAfterSalesUserMaintenanceByPhone(userPhone string) (*model.AfterSalesUserMaintenanceModel, error)

	// UpdateAfterSalesUserMaintenance 根据主键 ID 更新指定字段。
	UpdateAfterSalesUserMaintenance(id int64, updates map[string]interface{}) error
}

// AfterSalesUserMaintenanceInterfaceImp 售后用户维护档案数据访问实现。
type AfterSalesUserMaintenanceInterfaceImp struct{}

// ImpAfterSalesUserMaintenance 售后用户维护档案 DAO 实例。
var ImpAfterSalesUserMaintenance AfterSalesUserMaintenanceInterface = &AfterSalesUserMaintenanceInterfaceImp{}

var _ AfterSalesUserMaintenanceInterface = (*AfterSalesUserMaintenanceInterfaceImp)(nil)
