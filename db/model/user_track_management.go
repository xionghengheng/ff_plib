package model

import "encoding/json"

// UserTrackModel 客资基础档案主表
// 用于匹配 user_track 表的字段
type UserTrackModel struct {
	TrackID       int64  `json:"track_id"`       // 用户ID（系统自动生成）
	WechatNick    string `json:"wechat_nick"`    // 微信昵称（必填）
	WechatNo      string `json:"wechat_no"`      // 微信号（必填）
	TrainingNeed  string `json:"training_need"`  // 训练需求（非必填）
	ClassArea     string `json:"class_area"`     // 上课区域（非必填）
	ClassTime     string `json:"class_time"`     // 上课时间（非必填）
	IntentLevel   int    `json:"intent_level"`   // 意向等级（非必填，参考 Enum_Intent_Level）
	UserPhone     string `json:"user_phone"`     // 手机号（非必填，生成体验课链接后自动拉取或手填）
	ProfileRemark string `json:"profile_remark"` // 建档的时候填写，档案备注
	Stage         int    `json:"stage"`          // 当前状态（系统，参考 Enum_Track_Stage）
	AdvisorName   string `json:"advisor_name"`   // 销售负责人名称
	CreatedTs     int64  `json:"created_ts"`     // 创建时间
	UpdatedTs     int64  `json:"updated_ts"`     // 更新时间
}

// UserTrackNodeModel 客资节点数据表
// 一个节点一行，记录到达该节点的时间 + 该节点的数据（含每次跟进明细）
// 用于匹配 user_track_node 表的字段
type UserTrackNodeModel struct {
	NodeID           int64                     `json:"node_id"`            // 主键
	TrackID          int64                     `json:"track_id"`           // 客资ID
	Stage            int                       `json:"stage"`              // 节点 1~7（参考 Enum_Track_Stage）
	Data             string                    `json:"data"`               // 该节点的数据 JSON（结构按 stage 定，见注释）
	TrialPackageInfo UserTrackTrialPackageInfo `json:"trial_package_info"` // 体验课包信息（非数据库字段）
	PaidPackageInfo  UserTrackPaidPackageInfo  `json:"paid_package_info"`  // 付费课包信息（非数据库字段）
	NodeCreateTs     int64                     `json:"node_create_ts"`     // 到达该节点时间
	NodeUpdatedTs    int64                     `json:"node_updated_ts"`    // 更新时间
}

// UserTrackTrialPackageInfo 体验课包信息，具体字段由业务方补充。
type UserTrackTrialPackageInfo struct {
	PreTrailID int64 `json:"id"` // 预体验课id

	// 预体验课支付成功后，生成的
	PackageID string `json:"package_id"`
	LessonID  string `json:"lesson_id"`
}

// UserTrackPaidPackageInfo 付费课包信息，如果课包id已填充，说明用户购买成功
type UserTrackPaidPackageInfo struct {
	PackageID string `json:"package_id"`
}

// FollowUpRecord 表示一条跟进记录，可用于任意阶段。
type FollowUpRecord struct {
	FollowTs     int64  `json:"follow_ts"`     // 跟进时间
	FollowRemark string `json:"follow_remark"` // 跟进备注
}

// FollowUpRecords 是存放在 UserTrackNodeModel.Data 中的跟进记录 JSON 数组。
// JSON 示例：[{"follow_ts":1720000000,"follow_remark":"用户还在考虑"}]
type FollowUpRecords []FollowUpRecord

// GetFollowUpRecords 将节点 Data 的 JSON 数组解析为跟进记录列表。
// Data 为空字符串时返回 nil 列表和 nil 错误。
func GetFollowUpRecords(strData string) (FollowUpRecords, error) {
	var records FollowUpRecords
	if len(strData) == 0 {
		return records, nil
	}

	if err := json.Unmarshal([]byte(strData), &records); err != nil {
		return records, err
	}
	return records, nil
}

// 客资扭转状态（7个阶段，编号与业务 01~07 对齐）
const (
	Enum_Track_Stage_FirstContact    int = iota + 1 // 01 首联完成并建档（自动）
	Enum_Track_Stage_RequirementComm                // 02 需求沟通（手动）
	Enum_Track_Stage_InfoCollected                  // 03 完成信息搜集，待排体验（手动/自动）
	Enum_Track_Stage_TrailScheduled                 // 04 已排体验（手动/自动）
	Enum_Track_Stage_TrailDone                      // 05 已核销完成体验（手动/自动）
	Enum_Track_Stage_PendingDeal                    // 06 待成交（自动）
	Enum_Track_Stage_DealDone                       // 07 已成交（自动）
)

// 意向等级
const (
	Enum_Intent_Level_High int = iota + 1 // 高：明确表达体验/近期购买意愿
	Enum_Intent_Level_Mid                 // 中：有需求，但仍在考虑
	Enum_Intent_Level_Low                 // 低：暂无明确行动意愿
)
