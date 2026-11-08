package model

// KPIMonthTargetModel KPI 月度目标配置，对应 kpi_month_target 表，stat_month 唯一。
type KPIMonthTargetModel struct {
	ID                         int64  `json:"id"`
	StatMonth                  string `json:"stat_month"`             // 业务月份，YYYYMM
	LessonBaselineTarget       int    `json:"lesson_baseline_target"` // 月消课底线目标，节
	LessonSprintTarget         int    `json:"lesson_sprint_target"`   // 月消课冲刺目标，节
	NewPaidUserTarget          int    `json:"new_paid_user_target"`   // 月新增付费学员目标，人
	ScheduledTrialLessonTarget int    `json:"scheduled_trial_lesson_target"`
	GMVTargetFen               int64  `json:"gmv_target_fen"` // 月成交金额目标，分
	CreatedBy                  string `json:"created_by"`     // 创建人稳定账号标识
	UpdatedBy                  string `json:"updated_by"`     // 最近修改人稳定账号标识
	CreatedTs                  int64  `json:"created_ts"`     // 创建时间，Unix 秒
	UpdatedTs                  int64  `json:"updated_ts"`     // 最近修改时间，Unix 秒
}
