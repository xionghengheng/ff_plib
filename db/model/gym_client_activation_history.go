package model

// GymClientActivationHistoryModel 保存管理员成功开通门店端时的快照。
// 由业务层完成开通后填充快照，DAO 仅负责历史记录的存取。
type GymClientActivationHistoryModel struct {
	ID              int64  `json:"id" gorm:"primary_key"`
	UserPhone       string `json:"user_phone"`
	UID             int64  `json:"uid"`
	UserNick        string `json:"user_nick"`
	GymID           int    `json:"gym_id"`
	GymName         string `json:"gym_name"`
	PreviousGymID   int    `json:"previous_gym_id"`
	WasCoach        bool   `json:"was_coach"`
	PreviousCoachID int    `json:"previous_coach_id"`
	CreatedBy       string `json:"created_by"`
	CreatedTs       int64  `json:"created_ts"`
}
