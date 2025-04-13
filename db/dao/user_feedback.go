package dao

import (
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
	"time"
)

const user_feedback_tableName = "user_feedback"

// CreateFeedback 创建反馈记录
func (d *UserFeedBackInterfaceInterfaceImp) CreateFeedback(feedback *model.UserFeedbackModel) error {
	cli := db.Get()
	return cli.Table(user_feedback_tableName).Save(feedback).Error
}

// GetFeedbackList 获取用户反馈列表
func (d *UserFeedBackInterfaceInterfaceImp) GetFeedbackList(userID int64, page, pageSize int) ([]model.UserFeedbackModel, error) {
	var feedbacks []model.UserFeedbackModel
	cli := db.Get()
	err := cli.Table(user_feedback_tableName).Where("user_id = ?", userID).
		Order("create_ts desc").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&feedbacks).Error
	return feedbacks, err
}

// GetFeedbackByID 根据ID获取反馈
func (d *UserFeedBackInterfaceInterfaceImp) GetFeedbackByID(id int64) (*model.UserFeedbackModel, error) {
	var feedback model.UserFeedbackModel
	cli := db.Get()
	err := cli.Table(user_feedback_tableName).Where("id = ?", id).First(&feedback).Error
	if err != nil {
		return nil, err
	}
	return &feedback, nil
}

// UpdateFeedbackStatus 更新反馈状态
func (d *UserFeedBackInterfaceInterfaceImp) UpdateFeedbackStatus(id int64, status int, reply string) error {
	updates := map[string]interface{}{
		"status": status,
	}
	if reply != "" {
		updates["reply"] = reply
		updates["reply_ts"] = time.Now().Unix()
	}
	cli := db.Get()
	return cli.Table(user_feedback_tableName).Model(&model.UserFeedbackModel{}).Where("id = ?", id).Updates(updates).Error
}
