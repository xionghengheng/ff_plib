package model

// GymClientActivationHistoryModel 保存管理员成功开通门店端时的快照。
// 由业务层完成开通后填充快照，DAO 仅负责历史记录的存取。
type GymClientActivationHistoryModel struct {
	ID                int64  `json:"id" gorm:"primary_key"` // 历史记录主键
	UserPhone         string `json:"user_phone"`            // 开通用户的手机号快照
	UID               int64  `json:"uid"`                   // 开通用户ID
	UserNick          string `json:"user_nick"`             // 开通用户的昵称快照
	GymID             int    `json:"gym_id"`                // 开通门店ID
	GymName           string `json:"gym_name"`              // 开通门店名称快照
	BindPassCardGymId int    `json:"bind_pass_card_gym_id"` // 用户绑定的通卡场地ID
	CreatedBy         string `json:"created_by"`            // 执行开通操作的管理员标识
	CreatedTs         int64  `json:"created_ts"`            // 记录创建时间，Unix秒级时间戳
}
