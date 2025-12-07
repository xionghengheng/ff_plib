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
	FitnessExperience       int     `json:"fitness_experience"`               //健身经验（初级=1，中级=2，高级=3）
	FitnessGoal             int     `json:"fitness_goal"`                     //健身目标（1=减脂减重 2=增肌增重 3=塑型体态）
	DesiredWeight           int     `json:"desired_weight"`                   //期望体重
	TimeFrame               int     `json:"time_frame"`                       //达成目标时间（1=慢一点但稳定 2=正常速度 3=越快真好）
	PreferredBodyPart       string  `json:"preferred_body_part"`              //最期望增强部位
	WeeklyExerciseFrequency int     `json:"weekly_exercise_frequency"`        //每周运动次数（频次1~2次/周=1，频次3~4次/周=2，频次5~7次/周=3）
	PreferredPriceRange     int     `json:"preferred_price_range"`            //偏好价格档位(对应的体验课程id)
	PreferredLocationID     int     `json:"preferred_location_id"`            //偏好健身房场地ID
	VipType                 int     `json:"vip_type"`                         //vip订阅类型 0=非会员 1=体验会员 2=付费年费会员
	VipExpiredTs            int64   `json:"vip_expired_ts"`                   //vip过期时间
	IsCoach                 bool    `json:"is_coach"`                         //是否教练
	CoachId                 int     `json:"coach_id"`                         //如果是教练，关联的教练id
	HeadPicSafeStatus       int     `json:"head_pic_safe_status"`             //头像审核结果(参考 Enum_HeadPic_Check)
	HeadPicWaitSafe         string  `json:"head_pic_wait_safe"`               //等待审核的头像
	HeadPicSafeTraceId      string  `json:"head_pic_safe_trace_id"`           //等待审核的traceid，用户和异步回调匹配
	RegistTs                int64   `json:"regist_ts"`                        //用户注册时间
	BeVipTs                 int64   `json:"be_vip_ts"`                        //成为订阅会员的时间
	LastLoginTs             int64   `json:"last_login_ts"`                    //上次登录时间（目前只记录教练的）
	HasShownNoStoreHint     bool    `json:"has_shown_no_store_hint"`          //是否展示过未覆盖门店提醒
	BeVipFromType           int     `json:"be_vip_from_type"`                 //成为订阅会员的方式
	VipPassCardType         int     `json:"vip_pass_card_type"`               //通卡会员类型 0=非通卡会员 1=体验通卡会员 2=付费通卡会员月卡
	VipPassCardExpiredTs    int64   `json:"vip_pass_card_expired_ts"`         //通卡会员过期时间
	BeVipPassCardTs         int64   `json:"be_vip_pass_card_ts"`              //成为通卡订阅会员的时间
	BindPassCardGymId       int     `json:"bind_pass_card_gym_id"`            //通卡场地端，用户
	RecPassCardGymId        int     `json:"rec_pass_card_gym_id"`             //通卡主页推荐的场地
	IsOfficialAssistant     bool    `json:"is_official_assistant"`            //是否官方小助理，展示特定的运营操作入口
	//IsFirstTrailBlock       bool    `json:"is_first_trail_block"`             //是否需要展示，首次约课加小助理微信的阻断
	RemoveTrailBlockTs int64 `json:"remove_trail_block_ts"` //顾问是设置解除加小助理微信的阻断的时间，大于0表示设置过解除阻断
}

const (
	Enum_VipType_Non      int = iota // 0 = 非会员
	Enum_VipType_Trial               // 1 = 体验会员
	Enum_VipType_PaidYear            // 2 = 付费年费会员
)

const (
	Enum_VipPassCardType_Non       int = iota // 0 = 非会员
	Enum_VipPassCardType_Trial                // 1 = 通卡体验会员
	Enum_VipPassCardType_PaidMonth            // 2 = 通卡付费会员月卡
)

const (
	Enum_HeadPic_Check_Non  int = iota // 0 = 无审核状态
	Enum_HeadPic_Check_Ing             // 1 = 审核中
	Enum_HeadPic_Check_Pass            // 2 = 审核通过
	Enum_HeadPic_Check_Deny            // 3 = 审核拒绝
)

const (
	Enum_BE_VIP_FROM_TYPE_Non          int = iota // 0 = 无审核状态
	Enum_BE_VIP_FROM_TYPE_BUY_FROM_REC            // 线下完成体验课，通过顾问或教练分享，直接购买课包
)
