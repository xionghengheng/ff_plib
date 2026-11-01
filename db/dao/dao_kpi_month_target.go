package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

func (imp *KPIMonthTargetInterfaceImp) AddKPIMonthTarget(record *model.KPIMonthTargetModel) error {
	return db.Get().Table("kpi_month_target").Create(record).Error
}

func (imp *KPIMonthTargetInterfaceImp) UpdateKPIMonthTarget(statMonth string, updates map[string]interface{}) error {
	return db.Get().Table("kpi_month_target").Model(&model.KPIMonthTargetModel{}).
		Where("stat_month = ?", statMonth).Updates(updates).Error
}

func (imp *KPIMonthTargetInterfaceImp) GetKPIMonthTarget(statMonth string) (*model.KPIMonthTargetModel, error) {
	record := new(model.KPIMonthTargetModel)
	result := db.Get().Raw(`SELECT * FROM kpi_month_target WHERE stat_month = ? LIMIT 1`, statMonth).Scan(record)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return record, nil
}
