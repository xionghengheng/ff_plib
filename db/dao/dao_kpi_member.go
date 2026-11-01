package dao

import (
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

func (imp *KPIMemberInterfaceImp) AddKPIMember(record *model.KPIMemberModel) error {
	return db.Get().Table("kpi_member").Create(record).Error
}

func (imp *KPIMemberInterfaceImp) UpdateKPIMember(id int64, updates map[string]interface{}) error {
	return db.Get().Table("kpi_member").Model(&model.KPIMemberModel{}).
		Where("id = ?", id).Updates(updates).Error
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
