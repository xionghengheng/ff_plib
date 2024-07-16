package model

type CoachPageBannerModel struct {
	Id       int32  `json:"id"`
	PicUrl   string `json:"pic_url"`  // 图片链接
	Priority int    `json:"priority"` // 展示优先级，值越大，展示位置越靠前
}
