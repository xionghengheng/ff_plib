package model

// CoursePackageModel 课包模型（用户维度资产，用户可以得到体验课程课包，或者通过付费后购买得到课包）
type CoursePackageModel struct {
	PackageID          string `json:"package_id"`            // 课包的唯一标识符（用户id_获取课包的时间戳）
	PackageName        string `json:"package_name"`          // 课包的名称
	Uid                int64  `json:"uid"`                   // 用户id
	GymId              int    `json:"gym_id"`                // 场地id
	CourseId           int    `json:"course_id"`             // 课程id
	CoachId            int    `json:"coach_id"`              // 教练id
	Ts                 int64  `json:"ts"`                    // 获得课包的时间戳
	PackageType        int    `json:"package_type"`          // 课包类型(1=体验免费课包 2=付费)
	TotalCnt           int    `json:"total_cnt"`             // 课包中总的课程次数
	RemainCnt          int    `json:"remain_cnt"`            // 课包中剩余的课程次数
	Price              int    `json:"price"`                 // 价格
	CloseBarRemain2    bool   `json:"close_bar_remain2"`     // 是否关闭了剩余2节体验课底部bar展示
	CloseBarRemain1    bool   `json:"close_bar_remain1"`     // 是否关闭了剩余1节体验课底部bar展示
	LastLessonTs       int64  `json:"last_lesson_ts"`        // 最后一次上课时间
	UniidList          string `json:"uniid_list"`            // 扣减or增加次数的唯一订单号
	ChangeCoachTs      int64  `json:"change_coach_ts"`       // 更换教练的时间戳
	RefundTs           int64  `json:"refund_ts"`             // 发生退款的时间
	RefundLessonCnt    int    `json:"refund_lesson_cnt"`     // 退款课程数
	SendMsgTrailExpire bool   `json:"send_msg_trail_expire"` // 是否已发送消息提醒用户体验课快过期
	FirstTrialCoachId  int    `json:"first_trial_coach_id"`  // 第一次生成体验课包时候的教练id (分词要以下划线隔开。。。Error 1054: Unknown column 'first_trial_coach_id' in 'field list')
}

const (
	Enum_PackageType_TrialFree   int = iota + 1 // 1 = 体验免费课包
	Enum_PackageType_PaidPackage                // 2 = 付费课包
)
