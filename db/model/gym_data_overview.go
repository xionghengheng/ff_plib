package model

// GymDataOverviewModel 门店端数据概览（每小时离线构造，按门店+月份存储）
// GymID 为自训(通卡)门店id；私教数据按其关联的私教门店统计，自训数据按本门店统计
type GymDataOverviewModel struct {
	GymID      int   `gorm:"primary_key" json:"gym_id"`       // 自训(通卡)门店id
	MonthBegTs int64 `gorm:"primary_key" json:"month_beg_ts"` // 统计对应月份1号0点时间戳

	PrivateBookCnt     int `json:"private_book_cnt"`     // 私教-预约课次数（已预约）
	PrivateCompleteCnt int `json:"private_complete_cnt"` // 私教-已完成课次数（已完成）
	PrivatePendingCnt  int `json:"private_pending_cnt"`  // 私教-待核对次数（已旷课）

	SelfBookCnt     int `json:"self_book_cnt"`     // 自训-预约人次（已预约）
	SelfCompleteCnt int `json:"self_complete_cnt"` // 自训-已完成人次（已完成）
	SelfPendingCnt  int `json:"self_pending_cnt"`  // 自训-待核对人次（已旷课）

	UpdateTs int64 `json:"update_ts"` // 最近一次离线构造时间
}
