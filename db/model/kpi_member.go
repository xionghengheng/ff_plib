package model

const (
	KPIMemberTypePreSales   = 1 // 售前顾问
	KPIMemberTypeAfterSales = 2 // 售后交付
)

// KPIMemberModel KPI 人员配置，对应 kpi_member 表。
type KPIMemberModel struct {
	ID         int64  `json:"id"`          // KPI 人员 ID
	MemberName string `json:"member_name"` // 人员姓名
	MemberType int    `json:"member_type"` // 人员类型：1=售前顾问，2=售后交付
	CreatedTs  int64  `json:"created_ts"`  // 创建时间，Unix 秒
	UpdatedTs  int64  `json:"updated_ts"`  // 最近修改时间，Unix 秒
}
