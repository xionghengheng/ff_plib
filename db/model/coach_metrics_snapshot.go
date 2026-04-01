package model

type CoachMetricsSnapshotModel struct {
	ID                        int64 `json:"id"`                           // 主键
	CoachID                   int   `json:"coach_id"`                     // 教练ID
	MetricsAsOfDate           int   `json:"metrics_as_of_date"`           // 数据口径日期，格式 yyyymmdd
	FirstConversionCount      int   `json:"first_conversion_count"`       // 购课人数（实际首次转化数）
	TrialTotal                int   `json:"trial_total"`                  // 实际试课总数
	SecondConversionCount     int   `json:"second_conversion_count"`      // 二次转化数
	Redeem30dCount            int   `json:"redeem30d_count"`              // 近30天核销课程数
	TimeoutRedeemCount        int   `json:"timeout_redeem_count"`         // 超时核销课程数
	CourseTotalCount          int   `json:"course_total_count"`           // 课程总数
	RescheduleCount           int   `json:"reschedule_count"`             // 改课次数
	BookedCourseTotal         int   `json:"booked_course_total"`          // 预约课程总数
	TrainingFormFilledCount   int   `json:"training_form_filled_count"`   // 填写训练总数
	VerifiedCourseActualCount int   `json:"verified_course_actual_count"` // 实际核销课程数
	CreatedAt                 int64 `json:"created_at"`                   // 创建时间，Unix秒
	UpdatedAt                 int64 `json:"updated_at"`                   // 更新时间，Unix秒
}
