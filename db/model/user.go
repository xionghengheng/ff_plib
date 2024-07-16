package model

// UserModel 用户模型
// UserInfo 用于匹配 user_info 表的字段
type UserInfoModel struct {
	UserID                  int64   `json:"user_id"`                          //用户uid
	WechatID                string  `json:"wechat_id"`                        //微信openid
	PhoneNumber             *string `json:"phone_number" gorm:"default:null"` //手机号
	Nick                    string  `json:"nick,omitempty"`                   //昵称
	HeadPic                 string  `json:"head_pic"`                         //头像
	Gender                  int     `json:"gender"`                           //"0=男", "1=女", "2=other"
	Age                     int     `json:"age"`                              //年龄
	Weight                  int     `json:"weight"`                           //体重
	Height                  int     `json:"height"`                           //身高
	FitnessExperience       int     `json:"fitness_experience"`               //健身经验
	FitnessGoal             int     `json:"fitness_goal"`                     //健身目标
	DesiredWeight           int     `json:"desired_weight"`                   //期望体重
	TimeFrame               int     `json:"time_frame"`                       //期望多快达到
	PreferredBodyPart       string  `json:"preferred_body_part"`              //最期望增强部位
	WeeklyExerciseFrequency int     `json:"weekly_exercise_frequency"`        //每周运动次数
	PreferredPriceRange     int     `json:"preferred_price_range"`            //偏好价格档位
	PreferredLocationID     int     `json:"preferred_location_id"`            //偏好健身房场地ID
	VipType                 int     `json:"vip_type"`                         //vip订阅类型 0=非会员 1=体验会员 2=付费年费会员
	VipExpiredTs            int64   `json:"vip_expired_ts"`                   //vip过期时间
	IsCoach                 bool    `json:"is_coach"`                         //是否教练
	CoachId                 int     `json:"coach_id"`                         //如果是教练，关联的教练id
}

const (
	Enum_VipType_Non      int = iota // 0 = 非会员
	Enum_VipType_Trial               // 1 = 体验会员
	Enum_VipType_PaidYear            // 2 = 付费年费会员
)
