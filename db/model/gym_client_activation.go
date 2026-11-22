package model

// GymClientActivationModel 保存管理员成功开通门店端时的快照。
// 与用户绑定变更在同一事务内写入；失败操作不产生成功历史。
type GymClientActivationModel struct {
	ID              int64  `json:"id" gorm:"primary_key"`
	RequestID       string `json:"request_id"`
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
