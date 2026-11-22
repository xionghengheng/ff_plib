package dao

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

const gymClientActivationTable = "gym_client_activation"

func (imp *GymClientActivationInterfaceImp) Activate(phone string, gymID int, requestID, operator string) (*model.GymClientActivationModel, error) {
	return activateGymClient(db.Get(), phone, gymID, requestID, operator)
}

func activateGymClient(cli *gorm.DB, phone string, gymID int, requestID, operator string) (*model.GymClientActivationModel, error) {
	if phone == "" || gymID <= 0 || requestID == "" || operator == "" {
		return nil, fmt.Errorf("missing gym client activation parameters")
	}
	tx := cli.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer tx.Rollback()

	// 串行处理同一用户的开通，防止并发覆盖和重复提交。查询两条以拒绝歧义手机号。
	var users []model.UserInfoModel
	if err := tx.Table(user_tableName).Where("phone_number = ?", phone).
		Set("gorm:query_option", "FOR UPDATE").Limit(2).Find(&users).Error; err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, ErrGymClientUserNotFound
	}
	if len(users) != 1 {
		return nil, ErrGymClientAmbiguousPhone
	}
	user := users[0]

	var existing model.GymClientActivationModel
	err := tx.Table(gymClientActivationTable).Where("request_id = ?", requestID).First(&existing).Error
	if err == nil {
		if existing.UserPhone != phone || existing.GymID != gymID || existing.CreatedBy != operator {
			return nil, ErrGymClientRequestConflict
		}
		return &existing, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	var gym model.GymInfoModel
	err = tx.Table(gym_tableName).Where("gym_id = ?", gymID).First(&gym).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrGymClientGymNotFound
	}
	if err != nil {
		return nil, err
	}
	if err := ensurePassCardGym(tx, gymID); err != nil {
		return nil, err
	}
	if err := tx.Table(user_tableName).Model(&model.UserInfoModel{}).Where("user_id = ?", user.UserID).
		Updates(map[string]interface{}{"bind_pass_card_gym_id": gymID, "is_coach": false, "coach_id": 0}).Error; err != nil {
		return nil, err
	}
	record := &model.GymClientActivationModel{
		RequestID: requestID, UserPhone: phone, UID: user.UserID, UserNick: user.Nick,
		GymID: gymID, GymName: gym.LocName, PreviousGymID: user.BindPassCardGymId,
		WasCoach: user.IsCoach, PreviousCoachID: user.CoachId, CreatedBy: operator, CreatedTs: time.Now().Unix(),
	}
	if err := tx.Table(gymClientActivationTable).Create(record).Error; err != nil {
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (imp *GymClientActivationInterfaceImp) GetList(phone string, lastID int64, limit int) ([]model.GymClientActivationModel, error) {
	if limit <= 0 || limit > 100 {
		limit = 21
	}
	query := db.Get().Table(gymClientActivationTable)
	if phone != "" {
		query = query.Where("user_phone = ?", phone)
	}
	if lastID > 0 {
		query = query.Where("id < ?", lastID)
	}
	rows := make([]model.GymClientActivationModel, 0)
	err := query.Order("id DESC").Limit(limit).Find(&rows).Error
	return rows, err
}
