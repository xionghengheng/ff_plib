package model

type TaskModel struct {
	TaskID          int    `json:"task_id,omitempty"`
	DisplayOrder    int    `json:"display_order"`    // 展示排序
	TaskName        string `json:"task_name"`        // 任务标题名称（任务描述）
	TaskSubName     string `json:"task_sub_name"`    // 任务副标题名称（任务描述）
	TaskLevel       string `json:"task_level"`       // 任务等级，P0-P3
	TaskStatus      int    `json:"task_status"`      // 任务状态，0=下架 1=上架
	DocumentContent string `json:"document_content"` // 文档内容
	TaskType        int    `json:"task_type"`        // 任务类型，1=纯文本浏览 2=确认式
	BrowseDuration  int    `json:"browse_duration"`  // 浏览时长，单位秒
}

const (
	Enum_Task_Status_Offline int = iota // 0 - 下架
	Enum_Task_Status_Online             // 1 - 上架
)

const (
	Enum_Task_Type_TextBrowse int = iota + 1 // 1 - 纯文本浏览
	Enum_Task_Type_Confirm                   // 2 - 确认式
)

const (
	Enum_Task_Level_P0 = "P0"
	Enum_Task_Level_P1 = "P1"
	Enum_Task_Level_P2 = "P2"
	Enum_Task_Level_P3 = "P3"
)
