package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

func (imp *KPIMemberInterfaceImp) AddKPIMember(record *model.KPIMemberModel) error {
	return db.Get().Table("kpi_member").Create(record).Error
}

func (imp *KPIMemberInterfaceImp) UpdateKPIMember(id int64, updates map[string]interface{}) error {
	return db.Get().Transaction(func(tx *gorm.DB) error {
		var member model.KPIMemberModel
		if err := tx.Table("kpi_member").Set("gorm:query_option", "FOR UPDATE").Where("id = ?", id).First(&member).Error; err != nil {
			return err
		}
		return tx.Table("kpi_member").Model(&model.KPIMemberModel{}).Where("id = ?", id).Updates(updates).Error
	})
}

func (imp *KPIMemberInterfaceImp) GetKPIMemberList(memberType int) ([]model.KPIMemberModel, error) {
	var records []model.KPIMemberModel
	var err error
	if memberType == 0 {
		err = db.Get().Raw(`SELECT * FROM kpi_member ORDER BY id ASC`).Scan(&records).Error
	} else {
		err = db.Get().Raw(`SELECT * FROM kpi_member WHERE member_type = ? ORDER BY id ASC`, memberType).Scan(&records).Error
	}
	return records, err
}
