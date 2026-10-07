package model

import "encoding/json"

// AfterSalesUserMaintenanceModel 售后用户维护档案，这里的字段大部分是顾问手填写的
// 一个用户对应一条维护档案，Uid 和 UserPhone 在表中分别唯一。
type AfterSalesUserMaintenanceModel struct {
	// 关联标识
	ID int64 `json:"id"` // 主键 ID

	// 用户基础信息
	Uid               int64  `json:"uid"`                 // 用户uid（小程序）
	WechatNick        string `json:"wechat_nick"`         // 微信昵称
	WecomGroupName    string `json:"wecom_group_name"`    // 企微群名
	Salutation        string `json:"salutation"`          // 称呼
	UserPhone         string `json:"user_phone"`          // 手机号
	TrainingNeed      string `json:"training_need"`       // 训练需求
	Height            int    `json:"height_cm"`           // 身高，单位：厘米
	Weight            int    `json:"weight_gram"`         // 体重，单位：kg
	Age               int    `json:"age"`                 // 年龄
	MedicalHistory    string `json:"medical_history"`     // 历史病史
	RegularLessonTime string `json:"regular_lesson_time"` // 日常上课时间
	//MiniProgramNick   string `json:"mini_program_nick"`   // 用户昵称（小程序）（后台吐的时候，实时拉取）
	//LessonFrequencyStatus string `json:"lesson_frequency_status"` // 用户上课频次状态（后台吐的时候，实时拉取）
	//IsRenewed             bool   `json:"is_renewed"`              // 是否续费（后台吐的时候，实时拉取）
	//RenewalCount          int    `json:"renewal_count"`           // 续费次数（后台吐的时候，实时拉取）
	//TrialLessonTs         int64  `json:"trial_lesson_ts"`         // 试课时间，Unix 秒（后台吐的时候，实时拉取）

	// 课包信息，JSON 数组字符串；一个元素对应一个课包。
	PackageInfo string `json:"package_info"`

	// 补充说明与审计时间
	Remark    string `json:"remark"`     // 备注
	CreatedTs int64  `json:"created_ts"` // 创建时间，Unix 秒
	UpdatedTs int64  `json:"updated_ts"` // 更新时间，Unix 秒
}

// AfterSalesPackageInfo PackageInfo 数组中单个课包的数据结构。
type AfterSalesPackageInfo struct {
	PackageID string `json:"package_id"` // 已购课包 ID
	//TotalLessonCnt                int      `json:"total_lesson_cnt"`                 // 总课时（后台吐的时候，实时拉取）
	//PaidAmount                    int      `json:"paid_amount"`                      // 实付金额，单位：分（后台吐的时候，实时拉取）
	//CoachName                     string   `json:"coach_name"`                       // 教练姓名（后台吐的时候，实时拉取）
	//FirstFormalLessonTs           int64    `json:"first_formal_lesson_ts"`           // 首节正式课时间，Unix 秒（后台吐的时候，实时拉取）
	//RemainLessonCnt               int      `json:"remain_lesson_cnt"`                // 剩余课时（后台吐的时候，实时拉取）
	//CompletedLessonCnt            int      `json:"completed_lesson_cnt"`             // 已上课时（后台吐的时候，实时拉取）
	//LessonLocation                string   `json:"lesson_location"`                  // 上课地点（后台吐的时候，实时拉取）
	FirstTrainingPart             string   `json:"first_training_part"`              // 首次训练部位
	PhysicalAssessmentPhotos      []string `json:"physical_assessment_photos"`       // 体测照片
	FirstLessonFeedback           string   `json:"first_lesson_feedback"`            // 首节课反馈
	InitialPhysicalAssessmentData string   `json:"initial_physical_assessment_data"` // 首课体测数据
	CoachAssessmentSummary        string   `json:"coach_assessment_summary"`         // 教练评估总结
	EarlyMidFeedback              string   `json:"early_mid_feedback"`               // 初中期反馈
	MidFeedback                   string   `json:"mid_feedback"`                     // 中期反馈
	MidPhysicalAssessmentData     string   `json:"mid_physical_assessment_data"`     // 中期体测数据
	MidLateFeedback               string   `json:"mid_late_feedback"`                // 中后期反馈
	FinalPhysicalAssessmentData   string   `json:"final_physical_assessment_data"`   // 末期体测数据
	RenewalPreparationFeedback    string   `json:"renewal_preparation_feedback"`     // 续费铺垫反馈
	RenewalClosingFeedback        string   `json:"renewal_closing_feedback"`         // 续费收割反馈
}

// GetAfterSalesPackageInfos 将课包信息 JSON 数组字符串解析为课包列表。
func GetAfterSalesPackageInfos(strPackageInfo string) ([]AfterSalesPackageInfo, error) {
	var packageInfos []AfterSalesPackageInfo
	if len(strPackageInfo) > 0 {
		if err := json.Unmarshal([]byte(strPackageInfo), &packageInfos); err != nil {
			return packageInfos, err
		}
	}
	return packageInfos, nil
}

// MarshalAfterSalesPackageInfos 将课包列表序列化为 JSON 数组字符串。
func MarshalAfterSalesPackageInfos(packageInfos []AfterSalesPackageInfo) (string, error) {
	if packageInfos == nil {
		packageInfos = make([]AfterSalesPackageInfo, 0)
	}
	data, err := json.Marshal(packageInfos)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
