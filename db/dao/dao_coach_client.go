package dao

import (
	"ff_plib/db"
	"ff_plib/db/model"
)

const coach_client_home_page_banner_table_name = "coach_client_home_page_banner"

func (imp *CoachClientHomePageBannerInterfaceImp) GetHomePageBannerList() ([]model.CoachClientHomePageBannerModel, error) {
	var err error
	var vecBanner []model.CoachClientHomePageBannerModel
	cli := db.Get()
	err = cli.Table(coach_client_home_page_banner_table_name).Order("priority DESC").Find(&vecBanner).Error
	return vecBanner, err
}

const coach_client_trainee_reddot_table_name = "coach_client_trainee_reddot"

func (imp *CoachClientTraineeReddotInterfaceImp) GetTraineeReddot(coachId int, uid int64) (*model.CoachClientTraineeReddotModel, error) {
	var err error
	var reddot = new(model.CoachClientTraineeReddotModel)
	cli := db.Get()
	err = cli.Table(coach_client_trainee_reddot_table_name).Where("coach_id = ? AND trainee_uid = ?", coachId, uid).First(reddot).Error
	return reddot, err
}

func (imp *CoachClientTraineeReddotInterfaceImp) AddTraineeReddotVisit(stTraineeReddotModel *model.CoachClientTraineeReddotModel) error {
	cli := db.Get()
	return cli.Table(coach_client_trainee_reddot_table_name).Save(stTraineeReddotModel).Error
}



const coach_client_trainee_comment_table_name = "coach_client_trainee_comment"

func (imp *CoachClientTraineeCommentInterfaceImp) GetTraineeCommentList(coachId int, limit int) ([]model.CoachClientTraineeCommentModel, error) {
	var err error
	var vecCommentList []model.CoachClientTraineeCommentModel
	cli := db.Get()
	err = cli.Table(coach_client_trainee_comment_table_name).Where("coach_id = ? AND is_approved = ?", coachId, true).Order("comment_ts DESC").Find(&vecCommentList).Limit(limit).Error
	return vecCommentList, err
}

func (imp *CoachClientTraineeCommentInterfaceImp) AddTraineeComment(stCoachClientTraineeCommentModel *model.CoachClientTraineeCommentModel) error {
	cli := db.Get()
	return cli.Table(coach_client_trainee_comment_table_name).Save(stCoachClientTraineeCommentModel).Error
}


const coach_client_monthly_statistics_table_name = "coach_client_monthly_statistics"

func (imp *CoachClientMonthlyStatisticInterfaceImp) GetItem(coachId int, monthBegTs int64) (*model.CoachMonthlyStatisticModel, error) {
	var err error
	var stCoachMonthlyStatisticModel = new(model.CoachMonthlyStatisticModel)
	cli := db.Get()
	err = cli.Table(coach_client_monthly_statistics_table_name).Where("coach_id = ? AND month_beg_ts = ?", coachId, monthBegTs).First(stCoachMonthlyStatisticModel).Error
	return stCoachMonthlyStatisticModel, err
}

func (imp *CoachClientMonthlyStatisticInterfaceImp) AddItem(stCoachMonthlyStatisticModel *model.CoachMonthlyStatisticModel) error {
	cli := db.Get()
	return cli.Table(coach_client_monthly_statistics_table_name).Save(stCoachMonthlyStatisticModel).Error
}
