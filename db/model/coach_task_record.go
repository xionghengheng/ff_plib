package model

// CoachTaskRecordModel 教练任务完成记录
// 对应表 coach_task_record，记录教练对任务的完成状态
type CoachTaskRecordModel struct {
	ID             int64 `json:"id"`              // 主键ID
	CoachId        int   `json:"coach_id"`        // 教练id
	TaskID         int   `json:"task_id"`         // 任务id
	Status         int   `json:"status"`          // 完成状态，0=未完成 1=已完成
	BrowseDuration int   `json:"browse_duration"` // 实际浏览时长，单位秒
	FinishTs       int64 `json:"finish_ts"`       // 完成时间
	CreateTs       int64 `json:"create_ts"`       // 创建时间
	UpdateTs       int64 `json:"update_ts"`       // 更新时间
}

const (
	Enum_CoachTaskRecord_Status_Unfinished int = iota // 0 - 未完成
	Enum_CoachTaskRecord_Status_Finished              // 1 - 已完成
)
