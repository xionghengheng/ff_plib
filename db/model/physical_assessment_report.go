package model

import "encoding/json"

const (
	// 报告阶段：每个课包每个阶段只保留一份报告。
	Enum_PhysicalReportType_Initial int = iota + 1 // 初期：发起第一节课预约后解锁
	Enum_PhysicalReportType_Mid                    // 中期：已完成课时超过总课时 50% 后解锁
	Enum_PhysicalReportType_Final                  // 结课：课包剩余 3 节课时后解锁
)

const (
	// 报告后台核心状态，仅保留这三种状态。
	Enum_PhysicalReportStatus_Locked      int = iota + 1 // 待解锁
	Enum_PhysicalReportStatus_Unsubmitted                // 未提交（解锁态）
	Enum_PhysicalReportStatus_Submitted                  // 已提交（解锁态）
)

// PhysicalAssessmentReportModel 体测报告模型。
// 同一 PackageID + ReportType 只能有一条记录；UnlockTs 一旦写入不再回退。
type PhysicalAssessmentReportModel struct {
	ReportID              string `json:"report_id"`                // 报告唯一 ID
	PackageID             string `json:"package_id"`               // 关联课包 ID
	RelatedTrialPackageID string `json:"related_trial_package_id"` // 当前课包为首个付费课包时，关联的体验课包 ID；否则为空
	ReportType            int    `json:"report_type"`              // 初期 / 中期 / 结课
	Status                int    `json:"status"`                   // 待解锁 / 未提交 / 已提交

	Uid      int64 `json:"uid"`
	GymID    int   `json:"gym_id"`
	CourseID int   `json:"course_id"`
	CoachID  int   `json:"coach_id"`

	UnlockTs int64 `json:"unlock_ts"` // 首次满足解锁条件的时间
	SubmitTs int64 `json:"submit_ts"` // 提交时间，未提交时为 0
	CreateTs int64 `json:"create_ts"`
	UpdateTs int64 `json:"update_ts"`

	TrainingGoal string `json:"training_goal"`
	MeasureDate  int    `json:"measure_date"` // yyyyMMdd，例如 20240907
	MeasureTs    int64  `json:"measure_ts"`   // 测试时间 Unix 时间戳（秒）

	// 统一使用整数存储，避免浮点精度问题。
	WeightGram         int `json:"weight_gram"`         // 体重，单位kg
	BodyFatPercent     int `json:"body_fat_percent"`    // 体脂率，单位百分比
	Height             int `json:"height"`              // 身高，单位cm
	BMI                int `json:"bmi"`                 // BMI，单位百分比
	ChestCircumference int `json:"chest_circumference"` // 胸围，单位cm
	WaistCircumference int `json:"waist_circumference"` // 腰围，单位cm
	LegCircumference   int `json:"leg_circumference"`   // 腿围，单位cm
	ThighCircumference int `json:"thigh_circumference"` // 大腿围，单位cm

	// 图片 JSON 数组字符串，最多 6 张，例如：[{"url":"https://example.com/front.jpg"}]。
	PhotoURLs   string `json:"photo_urls"`
	CoachAdvice string `json:"coach_advice"` // 教练建议，提交时至少 10 个字符

	// 总结报告专有字段
	TrainingPerformanceEvaluation string `json:"training_performance_evaluation"` // 训练表现评价，提交时至少 20 个字符
	NextStageSuggestion           string `json:"next_stage_suggestion"`           // 下一阶段建议，提交时至少 20 个字符
	FollowUpPrecautions           string `json:"follow_up_precautions"`           // 后续注意事项；可选，填写时至少 20 个字符
}

// PhysicalAssessmentReportPhoto 对应图片 JSON 数组中的 {"url": "https://..."}。
type PhysicalAssessmentReportPhoto struct {
	URL string `json:"url"`
}

// GetPhysicalAssessmentReportPhotos 将图片 URL 的 JSON 数组字符串解析为图片列表。
func GetPhysicalAssessmentReportPhotos(strPhotoURLs string) ([]PhysicalAssessmentReportPhoto, error) {
	var photos []PhysicalAssessmentReportPhoto
	if len(strPhotoURLs) > 0 {
		if err := json.Unmarshal([]byte(strPhotoURLs), &photos); err != nil {
			return photos, err
		}
	}
	return photos, nil
}
