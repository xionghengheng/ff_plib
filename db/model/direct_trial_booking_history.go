package model

// DirectTrialBookingHistoryModel 已购体验课直接排课操作历史记录。
// 每条记录对应一次排课操作，保存提交时的信息快照及排课结果，用于历史查询和幂等校验。
// Status 记录该次操作的处理结果；课程后续取消、核销等状态以课程记录为准。
type DirectTrialBookingHistoryModel struct {
	ID            int64  `json:"id" gorm:"primary_key"` // 排课操作记录主键，自增ID
	RequestID     string `json:"request_id"`            // 客户端生成的操作唯一标识，用于幂等校验，避免重复排课
	UserPhone     string `json:"user_phone"`            // 学员手机号，记录提交时的值
	UID           int64  `json:"uid"`                   // 学员用户ID，由手机号查询获得
	PackageID     string `json:"package_id"`            // 本次排课使用的已购体验课包ID
	CoachID       int    `json:"coach_id"`              // 课包关联的教练ID
	CoachName     string `json:"coach_name"`            // 排课时的教练名称快照，用于历史展示
	GymID         int    `json:"gym_id"`                // 课包关联的门店ID
	GymName       string `json:"gym_name"`              // 排课时的门店名称快照，用于历史展示
	AppointmentID int    `json:"appointment_id"`        // 预约使用的前30分钟课表槽位ID，课程占用连续两个半小时槽位
	LessonTimeBeg int64  `json:"lesson_time_beg"`       // 课程开始时间，Unix时间戳（秒）
	LessonTimeEnd int64  `json:"lesson_time_end"`       // 课程结束时间，Unix时间戳（秒），为开始时间加3600秒
	LessonID      string `json:"lesson_id"`             // 排课成功后生成的课程ID，未确认成功时为空
	Status        string `json:"status"`                // 排课操作状态：processing-处理中、success-成功、failed-失败、unknown-结果待确认；不代表课程后续状态
	ErrorMsg      string `json:"error_msg"`             // 排课失败原因或结果待确认提示，成功时为空
	CreatedBy     string `json:"created_by"`            // 发起排课的顾问或管理员名称，由后台鉴权身份获取
	CreatedTs     int64  `json:"created_ts"`            // 操作记录创建时间，Unix时间戳（秒）
	UpdatedTs     int64  `json:"updated_ts"`            // 操作记录最后更新时间，Unix时间戳（秒）
}
