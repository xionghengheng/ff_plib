package pass_card_model

// UserID 对应 JSON 中的 {"uid": 123}
type UserID struct {
	Uid int64 `json:"uid"`
}

type PassCardAppointmentModel struct {
	AppointmentID int `json:"appointment_id"` // 预约ID
	GymId         int `json:"gym_id"`         // 健身房场地ID
	// ✨ 关键修改在这里 ✨
	// Uids 字段的类型现在直接是 []UserID
	// 它精确地映射了 [{"uid":123}, {"uid":567}] 这个JSON数组
	Uids            []UserID `json:"uids" gorm:"type:json"`
	AppointmentDate int64    `json:"appointment_date"` // 预约日期（当天0点时间戳）
	StartTime       int64    `json:"start_time"`       // 起始时间
	EndTime         int64    `json:"end_time"`         // 结束时间
	Status          int      `json:"status"`           // 门店端设置的可用状态，参考 Enum_PassCardAppointment_Status
	CreateTs        int64    `json:"create_ts"`        // 创建时间
	UpdateTs        int64    `json:"update_ts"`        // 更新时间
	MaxBookCnt      uint32   `json:"max_book_cnt"`     // 最大可约人数
}

// 门店端设置的可用状态
const (
	Enum_PassCardAppointment_Status_Invalid     int = iota // 0 - 无效
	Enum_PassCardAppointment_Status_Available              // 1 - 可预约
	Enum_PassCardAppointment_Status_UnAvailable            // 2 - 不可预约
)
