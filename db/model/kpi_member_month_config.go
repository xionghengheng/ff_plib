package model

// KPIMemberMonthConfigModel 每人每月客资及售后交付指标配置。
// 对应 kpi_member_month_config 表，按 stat_month + member_id 唯一。
// 可空指标 nil 表示未设置，指向 0 表示已设置为 0；百分比单位为 %（如 85.50）。
// 人员姓名和类型为该月快照，由业务层填写，不随人员配置自动更新。
type KPIMemberMonthConfigModel struct {
	ID                   int64    `json:"id"`
	MemberID             int64    `json:"member_id"`                                                                       // 关联 kpi_member.id
	StatMonth            string   `json:"stat_month"`                                                                      // 业务月份，YYYYMM
	MemberName           string   `json:"member_name"`                                                                     // 当月人员姓名快照
	MemberType           int      `json:"member_type"`                                                                     // 当月人员类型快照
	InitialLeadCnt       *int     `json:"initial_lead_cnt"`                                                                // 起始客资数，人
	ValidLeadCnt         *int     `json:"valid_lead_cnt"`                                                                  // 有效客资数，人
	SilentOver7dPct      *float64 `json:"silent_over_7d_pct" gorm:"column:silent_over_7d_pct;type:decimal(5,2)"`           // >7 天沉默用户占比
	FirstLessonOnTimePct *float64 `json:"first_lesson_on_time_pct" gorm:"type:decimal(5,2)"`                               // 首课按时率
	LessonConsumptionPct *float64 `json:"lesson_consumption_pct" gorm:"type:decimal(5,2)"`                                 // 课消率
	FourLessonsIn30dPct  *float64 `json:"four_lessons_in_30d_pct" gorm:"column:four_lessons_in_30d_pct;type:decimal(5,2)"` // 30 天内上课 4 次占比
	RenewalPct           *float64 `json:"renewal_pct" gorm:"type:decimal(5,2)"`                                            // 续费率
	CreatedTs            int64    `json:"created_ts"`                                                                      // 创建时间，Unix 秒
	UpdatedTs            int64    `json:"updated_ts"`                                                                      // 最近修改时间，Unix 秒
}
