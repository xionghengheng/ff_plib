package model

type ReportModel struct {
	UID        int64  `json:"uid"`
	ReportTime int64  `json:"report_time"` // 上报时间，由后台入库时添加
	IsCoach    bool   `json:"is_coach"`
	SessionID  string `json:"session_id"` // 访问id，uid + 用户进入小程序的时间戳，确保唯一性
	ItemID     string `json:"item_id"`    // 按钮id，每个点击对象分配一个，比如buy_vip
	ModuleID   string `json:"module_id"`  // 模块id，每个模块分配一个，比如course_info
	PageID     string `json:"page_id"`    // 页面id，每个页面分配一个，比如home
	Model      string `json:"model"`
	ActionID   int    `json:"action_id"` // 动作id，比如101代表曝光、102代表点击
	AppID      string `json:"app_id"`    // 应用id，全局唯一，比如funcoach
	Duration   int    `json:"duration"`  // 当前在小程序的停留时长，以用户进入小程序为基点计算
	BusiInfo   string `json:"busi_info"` // 额外信息，json字符串，比如一些非通用的附带状态等信息可以放到这里
	Brand      string `json:"brand"`
	EnvVersion string `json:"env_version"`
	Platform   string `json:"platform"`
	System     string `json:"system"`
	Version    string `json:"version"`
	StrExt1    string `json:"str_ext1"`
	StrExt2    string `json:"str_ext2"`
	StrExt3    string `json:"str_ext3"`
	Ext1       int    `json:"ext1"`
	Ext2       int    `json:"ext2"`
	Ext3       int    `json:"ext3"`
}
