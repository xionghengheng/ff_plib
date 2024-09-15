package model

type CourseModel struct {
	CourseID     int    `json:"course_id"`
	Name         string `json:"name"`         //课程名称
	Introduction string `json:"introduction"` //课程介绍
	Price        int    `json:"price"`        //课程价格，单位元
	Duration     int    `json:"duration"`     //课程时长，单位分钟
	Image        string `json:"image"`        //课程图片-方行
	ImageCircle  string `json:"image_circle"` //课程图片-圆形
	MarketPrice  int    `json:"market_price"` //课程市场价格，单位元
	ChargeType   int    `json:"charge_type"`  //1=付费，2=免费体验课
}

const (
	Enum_Course_ChargeType_Paid      = iota + 1 // 1 付费
	Enum_Course_ChargeType_FreeTrial            // 2 免费体验课
)
