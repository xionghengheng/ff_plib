package model

// UserFeedbackModel 用户反馈模型
type UserFeedbackModel struct {
	ID       int64   `json:"id" gorm:"primary_key"` // 反馈ID
	Uid      int64   `json:"uid"`                   // 用户ID
	Phone    *string `json:"phone,omitempty"`       // 用户手机号
	Content  string  `json:"content"`               // 反馈内容
	CreateTs int64   `json:"create_ts"`             // 创建时间
	Status   int     `json:"status"`                // 反馈状态
	Reply    string  `json:"reply,omitempty"`       // 回复内容
	ReplyTs  int64   `json:"reply_ts,omitempty"`    // 回复时间
}

const (
	Enum_FeedbackStatus_Pending int = iota // 0 = 待处理
	Enum_FeedbackStatus_Replied            // 1 = 已回复
	Enum_FeedbackStatus_Closed             // 2 = 已关闭
)
