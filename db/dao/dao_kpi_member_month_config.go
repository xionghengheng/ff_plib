package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

func (imp *KPIMemberMonthConfigInterfaceImp) AddKPIMemberMonthConfig(record *model.KPIMemberMonthConfigModel) error {
	return db.Get().Table("kpi_member_month_config").Create(record).Error
}

func (imp *KPIMemberMonthConfigInterfaceImp) UpdateKPIMemberMonthConfig(statMonth string, memberID int64, updates map[string]interface{}) error {
	return db.Get().Table("kpi_member_month_config").Model(&model.KPIMemberMonthConfigModel{}).
		Where("stat_month = ? AND member_id = ?", statMonth, memberID).Updates(updates).Error
}

func (imp *KPIMemberMonthConfigInterfaceImp) GetKPIMemberMonthConfig(statMonth string, memberID int64) (*model.KPIMemberMonthConfigModel, error) {
	record := new(model.KPIMemberMonthConfigModel)
	result := db.Get().Raw(`SELECT * FROM kpi_member_month_config WHERE stat_month = ? AND member_id = ? LIMIT 1`, statMonth, memberID).Scan(record)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return record, nil
}

func (imp *KPIMemberMonthConfigInterfaceImp) GetKPIMemberMonthConfigList(statMonth string) ([]model.KPIMemberMonthConfigModel, error) {
	var records []model.KPIMemberMonthConfigModel
	err := db.Get().Raw(`SELECT * FROM kpi_member_month_config WHERE stat_month = ? ORDER BY member_id ASC`, statMonth).Scan(&records).Error
	return records, err
}
