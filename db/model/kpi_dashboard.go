package model

// KPIDailyStat 的 DayIndex 是距查询月初的自然日偏移（上海时区）。
type KPIDailyStat struct {
	DayIndex     int   `json:"day_index"`
	Leads        int64 `json:"leads"`
	Appointments int64 `json:"appointments"`
	NewPaidUsers int64 `json:"new_paid_users"`
	Lessons      int64 `json:"lessons"`
	GMVFen       int64 `json:"gmv_fen"`
	SoldLessons  int64 `json:"sold_lessons"`
}
type KPITrackStat struct {
	AdvisorName  string `json:"advisor_name"`
	Leads        int64  `json:"leads"`
	Appointments int64  `json:"appointments"`
	Visits       int64  `json:"visits"`
	Deals        int64  `json:"deals"`
}
type KPIBusinessStats struct {
	Daily                   []KPIDailyStat `json:"daily"`
	RefundCount             int64          `json:"refund_count"`
	RefundFen               int64          `json:"refund_fen"`
	ZeroLessonUsers         *int64         `json:"zero_lesson_users"`
	FirstLessonOverdueUsers *int64         `json:"first_lesson_overdue_users"`
}
