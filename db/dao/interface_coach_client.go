package dao

import "FunFitnessTrainer/db/model"

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
}

// CoachClientTraineeReddotInterfaceImp
type CoachClientTraineeCommentInterfaceImp struct{}

// Imp 实现实例
var ImpCoachClientTraineeComment CoachClientTraineeCommentInterface = &CoachClientTraineeCommentInterfaceImp{}



// CoachClientMonthlyStatisticInterface
type CoachClientMonthlyStatisticInterface interface {
	GetItem(coachId int, monthBegTs int64) (*model.CoachMonthlyStatisticModel, error)
	AddItem(stCoachMonthlyStatisticModel *model.CoachMonthlyStatisticModel) error
}

// CoachClientMonthlyStatisticInterfaceImp
type CoachClientMonthlyStatisticInterfaceImp struct{}

// Imp 实现实例
var ImpCoachClientMonthlyStatistic CoachClientMonthlyStatisticInterface = &CoachClientMonthlyStatisticInterfaceImp{}
