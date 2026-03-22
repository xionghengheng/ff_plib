package model

// CoursePackageRenewalModel 课包续费记录模型
type CoursePackageRenewalModel struct {
	ID            int64  `json:"id"`             // 主键
	PackageID     string `json:"package_id"`     // 课包ID，关联CoursePackageModel
	CourseId      int    `json:"course_id"`      // 课程ID
	CoachId       int    `json:"coach_id"`       // 教练ID
	CourseCnt     int    `json:"course_cnt"`     // 续费课程数量
	PaymentAmount int    `json:"payment_amount"` // 续费金额（分）
	OrderID       string `json:"order_id"`       // 订单ID
	RenewalTime   int64  `json:"renewal_time"`   // 续费时间
}
