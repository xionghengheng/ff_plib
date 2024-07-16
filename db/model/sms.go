package model

// SmsModel 验证码短信模型
type SmsVerificationCodeModel struct {
	ID       int    `json:"id"`       // 验证码记录的唯一ID
	Unid     string `json:"unid"`     // 唯一标识符
	Code     string `json:"code"`     // 验证码
	Createts int64  `json:"createts"` // 创建时间
	Used     int    `json:"used"`     // 是否已使用的标志
}
