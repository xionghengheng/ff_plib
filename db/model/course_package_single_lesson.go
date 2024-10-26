package model

// CoursePackageSingleModel 课包内单节课模型（用户维度资产，用户发起预约后，生成的单次课记录）
type CoursePackageSingleLessonModel struct {
	LessonID                string `json:"lesson_id"`                   // 单节课的唯一标识符（用户id_场地id_课程id_教练id_发起预约的时间戳）
	PackageID               string `json:"package_id"`                  // 关联的课包的唯一标识符
	CreateTs                int64  `json:"create_ts"`                   // 记录生成时间，发起预约的时间
	ScheduleBegTs           int64  `json:"schedule_beg_ts"`             // 单节课的安排上课时间
	ScheduleEndTs           int64  `json:"schedule_end_ts"`             // 单节课的安排上课时间
	Status                  int    `json:"status"`                      // 单次课状态(已预约、已完成、已取消、已旷课)
	LessonName              string `json:"lesson_name"`                 // 单节课的名称
	Duration                int    `json:"duration"`                    // 单节课的时长，单位秒
	Uid                     int64  `json:"uid"`                         // 用户id
	CoachId                 int    `json:"coach_id"`                    // 教练id
	GymId                   int    `json:"gym_id"`                      // 场地id
	CourseID                int    `json:"course_id"`                   // 课程id
	AppointmentID           int    `json:"appointment_id"`              // 预约ID
	QrCodePic               []byte `json:"qr_code_pic"`                 // 核销小程序码
	CancelByCoach           bool   `json:"cancel_by_coach"`             // 是否是教练取消
	CancelByCoachDelCard    bool   `json:"cancel_by_coach_del_card"`    // 教练取消的情况，是否被叉掉课程卡，叉掉后卡片不常驻
	WriteOffMissedReturnCnt bool   `json:"write_off_missed_return_cnt"` // 是否发生了旷课归还次数
	SendMsgGoLesson         bool   `json:"send_msg_go_lesson"`          // 是否已发送上课前的提醒
	TrainContent            string `json:"train_content"`               // 训练内容（教练端设置）
	ScheduledByCoach        bool   `json:"scheduled_by_coach"`          // 是否为教练排课
	WriteOffTs              int64  `json:"write_off_ts"`                // 核销时间
	IsConfirm               bool   `json:"is_confirm"`                  // 是否已被教练确认

	//评论相关内容
	Overall              int    `json:"overall"`                // 整体
	Professional         int    `json:"professional"`           // 专业
	Environment          int    `json:"environment"`            // 环境
	Service              int    `json:"service"`                // 服务
	ContinueAttendLesson int    `json:"continue_attend_lesson"` // 是否愿意继续上课，愿意、待考虑、不愿意
	CommentContent       string `json:"comment_content"`        // 评价内容
	AnonymousComment     bool   `json:"anonymous_comment"`      // 是否匿名评价
	CommentTs            int64  `json:"comment_ts"`             // 提交评价的时间
}

// 正常流程：已预约->已完成
// 用户上课前主动取消：已预约->已取消
// 如果教练忘记核销或者用户没去：已预约->已旷课
const (
	En_LessonStatus_Scheduled int = iota + 1 // 已预约
	En_LessonStatusCompleted                 // 已完成
	En_LessonStatusCanceled                  // 已取消
	En_LessonStatusMissed                    // 已旷课
)

const (
	Enum_Continue_Attend_Lesson_Invalid   int = iota // 0 - 无效
	Enum_Continue_Attend_Lesson_Willing              // 1 - 愿意
	Enum_Continue_Attend_Lesson_Undecided            // 2 - 待考虑
	Enum_Continue_Attend_Lesson_Unwilling            // 3 - 不愿意
)
