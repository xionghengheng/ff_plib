package model

type CoachModel struct {
	CoachID       int    `json:"coach_id"`        //教练id
	GymID         int    `json:"gym_id"`          //健身房id
	CoachName     string `json:"coach_name"`      //教练名称
	Avatar        string `json:"avatar"`          //教练头像url
	Bio           string `json:"bio"`             //教练简介
	Rating        int    `json:"rating"`          //教练评分
	RecReason     string `json:"rec_reason"`      //教练推荐原因
	Priority      int    `json:"priority"`        //教练优先级
	CourseIdList  string `json:"course_id_list"`  //教练可上课程列表，英文逗号分割
	BannerPicList string `json:"banner_pic_list"` //教练详情页顶部banner图（英文逗号分割）
	GoodAt        string `json:"good_at"`         //教练擅长领域
	Phone         string `json:"phone"`           //手机号
	CircleAvatar  string `json:"circle_avatar"`   //教练圆形头像url
	JoinTs        int64  `json:"join_ts"`         //教练入驻时间（后台配置生效）
}
