package dao

import "github.com/xionghengheng/ff_plib/db/model"

// CoachClientHomePageBannerInterface 教练端主页顶部banner
type CoachClientHomePageBannerInterface interface {
	GetHomePageBannerList() ([]model.CoachClientHomePageBannerModel, error)
}

// CoachClientHomePageBannerInterfaceImp
type CoachClientHomePageBannerInterfaceImp struct{}

// Imp 实现实例
var ImpCoachClientHomePageBanner CoachClientHomePageBannerInterface = &CoachClientHomePageBannerInterfaceImp{}

// CoachClientTraineeReddotInterface 教练端体验用户新用户红点
type CoachClientTraineeReddotInterface interface {
	GetTraineeReddot(coachId int, uid int64) (*model.CoachClientTraineeReddotModel, error)
	AddTraineeReddotVisit(stTraineeReddotModel *model.CoachClientTraineeReddotModel) error
}

// CoachClientTraineeReddotInterfaceImp
type CoachClientTraineeReddotInterfaceImp struct{}

// Imp 实现实例
var ImpCoachClientTraineeReddot CoachClientTraineeReddotInterface = &CoachClientTraineeReddotInterfaceImp{}

// CoachClientTraineeCommentInterface
type CoachClientTraineeCommentInterface interface {
	GetTraineeCommentList(coachId int, limit int) ([]model.CoachClientTraineeCommentModel, error)
	AddTraineeComment(stTraineeReddotModel *model.CoachClientTraineeCommentModel) error
	// 更新评价
	UpdateTraineeComment(lessonID string, mapUpdates map[string]interface{}) error
}

// CoachClientTraineeReddotInterfaceImp
type CoachClientTraineeCommentInterfaceImp struct{}

// Imp 实现实例
var ImpCoachClientTraineeComment CoachClientTraineeCommentInterface = &CoachClientTraineeCommentInterfaceImp{}

// CoachClientMonthlyStatisticInterface
type CoachClientMonthlyStatisticInterface interface {
	GetAllItem() ([]model.CoachMonthlyStatisticModel, error)
	GetItem(coachId int, monthBegTs int64) (*model.CoachMonthlyStatisticModel, error)
	AddItem(stCoachMonthlyStatisticModel *model.CoachMonthlyStatisticModel) error
}

// CoachClientMonthlyStatisticInterfaceImp
type CoachClientMonthlyStatisticInterfaceImp struct{}

// Imp 实现实例
var ImpCoachClientMonthlyStatistic CoachClientMonthlyStatisticInterface = &CoachClientMonthlyStatisticInterfaceImp{}

// CoachClientTraineeNicknameInterface 教练对用户自定义昵称
type CoachClientTraineeNicknameInterface interface {
	GetTraineeNickname(coachId int, uid int64) (*model.CoachClientTraineeNicknameModel, error)
	// 不存在则新增，已存在则更新（每个教练对每个用户只有一条记录）
	UpsertTraineeNickname(coachId int, uid int64, nickname string) error
}

// CoachClientTraineeNicknameInterfaceImp
type CoachClientTraineeNicknameInterfaceImp struct{}

// Imp 实现实例
var ImpCoachClientTraineeNickname CoachClientTraineeNicknameInterface = &CoachClientTraineeNicknameInterfaceImp{}
