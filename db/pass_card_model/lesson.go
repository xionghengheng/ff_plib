package pass_card_model

// 课包内单节课模型（用户维度资产，用户发起预约后，生成的单次课记录）
type LessonModel struct {
	LessonID      string `json:"lesson_id"`       // 单节课的唯一标识符（用户id_场地id_课程id_教练id_发起预约的时间戳）
	CreateTs      int64  `json:"create_ts"`       // 记录生成时间，发起预约的时间
	ScheduleBegTs int64  `json:"schedule_beg_ts"` // 单节课的安排上课时间
	ScheduleEndTs int64  `json:"schedule_end_ts"` // 单节课的安排上课时间
	Status        int    `json:"status"`          // 单次课状态(已预约、已完成、已取消)
	LessonName    string `json:"lesson_name"`     // 单节课的名称
	Duration      int    `json:"duration"`        // 单节课的时长，单位秒
	Uid           int64  `json:"uid"`             // 用户id
	GymId         int    `json:"gym_id"`          // 场地id
	AppointmentID int    `json:"appointment_id"`  // 预约ID
	QrCodePic     []byte `json:"qr_code_pic"`     // 核销小程序码
	WriteOffTs    int64  `json:"write_off_ts"`    // 核销时间
}

// 正常流程：已预约->已完成
// 用户上课前主动取消：已预约->已取消
// 如果教练忘记核销或者用户没去：已预约->已旷课
const (
	En_LessonStatus_Scheduled int = iota + 1 // 已预约
	En_LessonStatusCompleted                 // 已完成
	En_LessonStatusCanceled                  // 已取消
)
