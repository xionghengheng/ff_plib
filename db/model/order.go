package model

type PaymentOrderModel struct {
	OrderID        string `json:"order_id"`        // 业务支付订单id（唯一id）
	WechatOrderID  string `json:"wechat_order_id"` // 微信支付订单号
	OrderTime      int64  `json:"order_time"`      // 下单时间（Unix 时间戳）
	OrderStatus    int    `json:"order_status"`    // 订单状态
	PaymentTime    int64  `json:"payment_time"`    // 支付成功时间（Unix 时间戳）
	PayerUID       int64  `json:"payer_uid"`       // 付款人用户uid
	PaymentAmount  int    `json:"payment_amount"`  // 付款金额（以最小货币单位存储，如分）
	DiscountAmount int    `json:"discount_amount"` // 优惠金额（以最小货币单位存储，如分）
	PaymentChannel string `json:"payment_channel"` // 付款渠道
	PurchaseType   int    `json:"purchase_type"`   // 购买类型（1=vip，2=vip续费，3=课包）
	PackageID      int    `json:"package_id"`      // 课包的唯一标识符（用户id_教练id_获取课包的时间戳）
	CancelTime     int64  `json:"cancel_time"`     // 取消时间

	//每一笔订单，需要记录微信的payment参数
	AppId     string `json:"app_id"`
	TimeStamp string `json:"time_stamp"`
	NonceStr  string `json:"nonce_str"`
	Package   string `json:"package"`
	SignType  string `json:"sign_type"`
	PaySign   string `json:"pay_sign"`

	//购买课包时下单的参数
	GymId     int `json:"gym_id"`     // 场地id
	CoachId   int `json:"coach_id"`   // 教练id
	CourseId  int `json:"course_id"`  // 课程id
	Price     int `json:"price"`      // 价格
	CourseCnt int `json:"course_cnt"` // 购买课程的次数
}

// PaymentStatus 枚举类型
const (
	Enum_Pay_Status_Pending   int = iota // 0 - 待支付
	Enum_Pay_Status_Paid                 // 1 - 已支付
	Enum_Pay_Status_Cancelled            // 2 - 已取消
	Enum_Pay_Status_Expired              // 3 - 已过期
)
