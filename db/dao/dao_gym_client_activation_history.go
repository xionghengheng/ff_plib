package dao

import (
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

const gymClientActivationHistoryTable = "gym_client_activation_history"

func (imp *GymClientActivationHistoryInterfaceImp) Create(record *model.GymClientActivationHistoryModel) error {
	return db.Get().Table(gymClientActivationHistoryTable).Create(record).Error
}

func (imp *GymClientActivationHistoryInterfaceImp) GetList(phone string, lastID int64, limit int) ([]model.GymClientActivationHistoryModel, error) {
	if limit <= 0 || limit > 100 {
		limit = 21
	}
	query := db.Get().Table(gymClientActivationHistoryTable)
	if phone != "" {
		query = query.Where("user_phone = ?", phone)
	}
	if lastID > 0 {
		query = query.Where("id < ?", lastID)
	}
	rows := make([]model.GymClientActivationHistoryModel, 0)
	err := query.Order("id DESC").Limit(limit).Find(&rows).Error
	return rows, err
}
