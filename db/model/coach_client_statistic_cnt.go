package model

// 表结构
type CoachMonthlyStatisticModel struct {
	CoachID           int    `gorm:"primary_key" json:"coach_id"`     // 教练ID
	MonthBegTs        int64  `gorm:"primary_key" json:"month_beg_ts"` // 统计计数对应的月份
	PayUserCnt        uint32 `json:"pay_user_cnt"`                    // 付费用户数(去重的)
	LessonCnt         uint32 `json:"lesson_cnt"`                      // 上课数
	LessonUserCnt     uint32 `json:"lesson_user_cnt"`                 // 上课用户数(去重的)
	SaleRevenue       uint32 `json:"sale_revenue"`                    // 销售额(单位元)
	TrialUserCnt      uint32 `json:"trial_user_cnt"`                  // 体验用户数-体验课包
	TrialLessonCnt    uint32 `json:"trial_lesson_cnt"`                // 体验课上课数-体验课包
	GiftLessonCnt     uint32 `json:"gift_lesson_cnt"`                 // 赠课次数-赠课课包
	UsedGiftLessonCnt uint32 `json:"used_gift_lesson_cnt"`            // 已上赠课数-赠课课包
}
