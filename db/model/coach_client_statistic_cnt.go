package model

// 表结构
type CoachMonthlyStatisticModel struct {
	CoachID       int    `gorm:"primary_key" json:"coach_id"`     // 教练ID
	MonthBegTs    int64  `gorm:"primary_key" json:"month_beg_ts"` // 统计计数对应的月份
	PayUserCnt    uint32 `json:"pay_user_cnt"`                   // 付费用户数(去重的)
	LessonCnt     uint32 `json:"lesson_cnt"`                     // 上课数
	LessonUserCnt uint32 `json:"lesson_user_cnt"`                // 上课用户数(去重的)
	SaleRevenue   uint32 `json:"sale_revenue"`                   // 销售额(单位元)
}

