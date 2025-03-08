package model

type CoachModel struct {
	CoachID             int    `json:"coach_id"`              //教练id
	GymID               int    `json:"gym_id"`                //健身房id
	CoachName           string `json:"coach_name"`            //教练名称
	Avatar              string `json:"avatar"`                //教练头像url
	Bio                 string `json:"bio"`                   //教练简介
	Rating              int    `json:"rating"`                //教练评分
	RecReason           string `json:"rec_reason"`            //教练推荐原因
	Priority            int    `json:"priority"`              //教练优先级
	CourseIdList        string `json:"course_id_list"`        //教练可上课程列表，英文逗号分割
	BannerPicList       string `json:"banner_pic_list"`       //教练详情页顶部banner图（英文逗号分割）
	GoodAt              string `json:"good_at"`               //教练擅长领域
	Phone               string `json:"phone"`                 //手机号
	CircleAvatar        string `json:"circle_avatar"`         //教练圆形头像url
	JoinTs              int64  `json:"join_ts"`               //教练入驻时间（后台配置生效）
	BTestCoach          bool   `json:"b_test_coach"`          //是否测试教练，外网白名单测试专用
	GymIdList           string `json:"gym_list"`              //教练支持多门店（英文逗号分割）
	QualifyType         int    `json:"qualify_type"`          //教练资质类型 (参考枚举 Enum_Coach_QualifyType)
	SkillCertification  string `json:"skill_certification"`   //教练的技能认证（英文逗号分割）
	Style               string `json:"style"`                 //教练风格（英文逗号分割）
	YearsOfWork         string `json:"years_of_work"`         //从业时长
	TotalCompleteLesson string `json:"total_complete_lesson"` //累计上课节数
	CanShow             int    `json:"can_show"`              //是否可以在平台展示，0=可展示 1=不可展示
}

const (
	Enum_Coach_QualifyType_Basic        int = iota + 1 // 1 = 基础
	Enum_Coach_QualifyType_Intermediate                // 2 = 中级
	Enum_Coach_QualifyType_Advanced                    // 3 = 高级
	Enum_Coach_QualifyType_Senior                      // 4 = 资深
)
