package model

// GiftCourseRecordModel 赠课记录（管理后台给用户赠送课时的操作流水）
// 对应表 gift_course_record，每成功赠送一次落一条记录；课包创建成功后回填 package_id
type GiftCourseRecordModel struct {
	ID           int64  `json:"id"`            // 主键ID
	OperatorName string `json:"operator_name"` // 操作人（顾问昵称，否则管理员用户名）
	OpTs         int64  `json:"op_ts"`         // 操作时间

	Uid         int64  `json:"uid"`          // 用户id
	PhoneNumber string `json:"phone_number"` // 用户手机号
	UserName    string `json:"user_name"`    // 用户昵称

	GiftType        int   `json:"gift_type"`         // 赠课类型（见 Enum_GiftType_*）
	GiftLessonCnt   int   `json:"gift_lesson_cnt"`   // 赠送节数
	ApplyCourseType int   `json:"apply_course_type"` // 适用课程类型（见 Enum_GiftApplyCourseType_*）
	ValidBegTs      int64 `json:"valid_beg_ts"`      // 有效期开始时间(预留，暂时先不做有效期处理)
	ValidEndTs      int64 `json:"valid_end_ts"`      // 有效期结束时间(预留，暂时先不做有效期处理)

	BCountGmv        bool `json:"b_count_gmv"`        // 是否计入GMV
	BCoachSettlement bool `json:"b_coach_settlement"` // 是否参与教练结算
	BGymSettlement   bool `json:"b_gym_settlement"`   // 是否参与门店结算

	GiftReason int    `json:"gift_reason"` // 赠课原因（见 Enum_GiftReason_*）
	Remark     string `json:"remark"`      // 备注

	// 赠课课包的教练/门店/课程信息（从用户选定的来源付费课包带出）
	SrcPackageId string `json:"src_package_id"` // 来源付费课包id
	GymId        int    `json:"gym_id"`         // 门店id
	GymName      string `json:"gym_name"`       // 门店名称
	CoachId      int    `json:"coach_id"`       // 教练id
	CoachName    string `json:"coach_name"`     // 教练姓名
	CourseId     int    `json:"course_id"`      // 课程id
	CourseName   string `json:"course_name"`    // 课程名称

	PackageId string `json:"package_id"` // 赠课生成的课包id（课包创建成功后回填）
}

// 赠课类型枚举
const (
	Enum_GiftType_PT1V1 = 1 // 1V1私教赠课
)

// 适用课程类型枚举
const (
	Enum_GiftApplyCourseType_PT1V1 = 1 // 1V1私教赠课
)

// 赠课原因枚举
const (
	Enum_GiftReason_Activity     = 1  // 活动赠课
	Enum_GiftReason_Compensation = 2  // 补偿赠课
	Enum_GiftReason_Referral     = 3  // 老带新
	Enum_GiftReason_Other        = 99 // 其他
)
