package dao

import (
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

const direct_trial_booking_table_name = "direct_trial_booking"

func (imp *DirectTrialBookingInterfaceImp) GetByRequestID(requestID string) (*model.DirectTrialBookingHistoryModel, error) {
	var record model.DirectTrialBookingHistoryModel
	err := db.Get().Table(direct_trial_booking_table_name).Where("request_id = ?", requestID).First(&record).Error
	return &record, err
}
func (imp *DirectTrialBookingInterfaceImp) Create(record *model.DirectTrialBookingHistoryModel) error {
	return db.Get().Table(direct_trial_booking_table_name).Create(record).Error
}
func (imp *DirectTrialBookingInterfaceImp) UpdateResult(id int64, status, lessonID, message string, now int64) error {
	return db.Get().Table(direct_trial_booking_table_name).Model(&model.DirectTrialBookingHistoryModel{}).Where("id = ? AND status <> ?", id, "success").Updates(map[string]interface{}{
		"status": status, "lesson_id": lessonID, "error_msg": message, "updated_ts": now,
	}).Error
}
func (imp *DirectTrialBookingInterfaceImp) GetList(phone string, lastID int64, limit int) ([]model.DirectTrialBookingHistoryModel, error) {
	query := db.Get().Table(direct_trial_booking_table_name).Model(&model.DirectTrialBookingHistoryModel{})
	if phone != "" {
		query = query.Where("user_phone = ?", phone)
	}
	if lastID > 0 {
		query = query.Where("id < ?", lastID)
	}
	rows := make([]model.DirectTrialBookingHistoryModel, 0)
	err := query.Order("id DESC").Limit(limit).Find(&rows).Error
	return rows, err
}
