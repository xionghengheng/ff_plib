package model

type InvitationCodeModel struct {
	ID             int    `json:"id"`
	CompanyName    string `json:"company_name"`
	InvitationCode string `json:"invitation_code"`
	CreateTS       int64  `json:"create_ts"`
	IsUsed         bool   `json:"is_used"`
	UsedUid        int64  `json:"used_uid"`
	UsedTs         int64  `json:"used_ts"`
}
