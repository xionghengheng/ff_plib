package model

type CoachAppointmentModel struct {
	AppointmentID    int   `json:"appointment_id"`     // 预约ID
	CoachID          int   `json:"coach_id"`           // 教练ID
	GymId            int   `json:"gym_id"`             // 健身房场地ID
	UserID           int64 `json:"user_id"`            // 用户ID(用户预约成功后更新)
	UserCourseID     int   `json:"user_course_id"`     // 课程ID(用户预约成功后更新)
	AppointmentDate  int64 `json:"appointment_date"`   // 预约日期（当天0点时间戳）
	StartTime        int64 `json:"start_time"`         // 起始时间
	EndTime          int64 `json:"end_time"`           // 结束时间
	Status           int   `json:"status"`             // 预约状态
	CreateTs         int64 `json:"create_ts"`          // 创建时间
	UpdateTs         int64 `json:"update_ts"`          // 更新时间
	ScheduledByCoach bool  `json:"scheduled_by_coach"` // 是否为教练排课
}

// AppointmentStatus 预约状态
const (
	Enum_Appointment_Status_Invalid     int = iota // 0 - 无效
	Enum_Appointment_Status_Available              // 1 - 可预约
	Enum_Appointment_Status_UnAvailable            // 2 - 不可预约
)
