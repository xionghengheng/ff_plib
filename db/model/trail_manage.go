package model

// 体验课预生成管理模型
// 用于匹配 trail_manage 表的字段
type PreTrailManageModel struct {
	ID            int64  `json:"id"`              // 主键ID
	UserPhone     string `json:"user_phone"`      // 用户手机号（微信绑定）
	TrainingNeed  string `json:"training_need"`   // 训练需求（增肌/减脂/塑形/普拉提/拳击/体态/康复/上门/其他）
	GymID         int    `json:"gym_id"`          // 门店ID
	GymName       string `json:"gym_name"`        // 门店名称
	CoachID       int    `json:"coach_id"`        // 教练ID
	CoachName     string `json:"coach_name"`      // 教练名称
	LessonDate    int64  `json:"lesson_date"`     // 体验课日期
	LessonTimeBeg int64  `json:"lesson_time_beg"` // 体验课开始时间
	LessonTimeEnd int64  `json:"lesson_time_end"` // 体验课结束时间
	Price         int    `json:"price"`           // 体验课价格（单位：元）
	H5LinkToken   string `json:"h5_link_token"`   // 生成的H5链接token
	PackageId     string `json:"package_id"`      // 关联的课包id
	Remark        string `json:"remark"`          // 备注
	CreatedBy     string `json:"created_by"`      // 创建人（顾问）
	CreatedAt     int64  `json:"created_at"`      // 创建时间
	UpdatedAt     int64  `json:"updated_at"`      // 更新时间
}
