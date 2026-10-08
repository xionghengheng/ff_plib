package dao

import (
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

// AddAfterSalesUserMaintenance 添加售后用户维护档案。
func (imp *AfterSalesUserMaintenanceInterfaceImp) AddAfterSalesUserMaintenance(maintenance *model.AfterSalesUserMaintenanceModel) error {
	return db.Get().Table("after_sales_user_maintenance").Create(maintenance).Error
}

// GetAfterSalesUserMaintenanceList 按创建时间倒序游标分页查询档案列表。
func (imp *AfterSalesUserMaintenanceInterfaceImp) GetAfterSalesUserMaintenanceList(lastCreatedTs, lastID int64, limit int) ([]model.AfterSalesUserMaintenanceModel, error) {
	var maintenanceList []model.AfterSalesUserMaintenanceModel
	var err error
	if lastCreatedTs > 0 {
		err = db.Get().Raw(`
			SELECT *
			FROM after_sales_user_maintenance
			WHERE created_ts < ? OR (created_ts = ? AND id < ?)
			ORDER BY created_ts DESC, id DESC
			LIMIT ?`, lastCreatedTs, lastCreatedTs, lastID, limit).
			Scan(&maintenanceList).Error
	} else {
		err = db.Get().Raw(`
			SELECT *
			FROM after_sales_user_maintenance
			ORDER BY created_ts DESC, id DESC
			LIMIT ?`, limit).
			Scan(&maintenanceList).Error
	}
	return maintenanceList, err
}

// GetAfterSalesUserMaintenanceListByLatestPackagePurchaseTs 按最新课包购买时间倒序游标分页查询档案列表。
func (imp *AfterSalesUserMaintenanceInterfaceImp) GetAfterSalesUserMaintenanceListByLatestPackagePurchaseTs(lastLatestPackagePurchaseTs, lastID int64, limit int) ([]model.AfterSalesUserMaintenanceModel, error) {
	var maintenanceList []model.AfterSalesUserMaintenanceModel
	var err error
	if lastID > 0 {
		err = db.Get().Raw(`
			SELECT *
			FROM after_sales_user_maintenance
			WHERE latest_package_purchase_ts < ?
			   OR (latest_package_purchase_ts = ? AND id < ?)
			ORDER BY latest_package_purchase_ts DESC, id DESC
			LIMIT ?`, lastLatestPackagePurchaseTs, lastLatestPackagePurchaseTs, lastID, limit).
			Scan(&maintenanceList).Error
	} else {
		err = db.Get().Raw(`
			SELECT *
			FROM after_sales_user_maintenance
			ORDER BY latest_package_purchase_ts DESC, id DESC
			LIMIT ?`, limit).
			Scan(&maintenanceList).Error
	}
	return maintenanceList, err
}

// GetAfterSalesUserMaintenanceByPhone 根据手机号查询单条档案。
func (imp *AfterSalesUserMaintenanceInterfaceImp) GetAfterSalesUserMaintenanceByPhone(userPhone string) (*model.AfterSalesUserMaintenanceModel, error) {
	maintenance := new(model.AfterSalesUserMaintenanceModel)
	err := db.Get().Raw(`
		SELECT *
		FROM after_sales_user_maintenance
		WHERE user_phone = ?
		LIMIT 1`, userPhone).
		Scan(maintenance).Error
	return maintenance, err
}

// UpdateAfterSalesUserMaintenance 根据主键 ID 更新指定字段。
func (imp *AfterSalesUserMaintenanceInterfaceImp) UpdateAfterSalesUserMaintenance(id int64, updates map[string]interface{}) error {
	return db.Get().Table("after_sales_user_maintenance").
		Model(&model.AfterSalesUserMaintenanceModel{}).
		Where("id = ?", id).
		Updates(updates).Error
}
