package dao

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
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

// UpdateTraineeComment 更新评价
func (imp *CoachClientTraineeCommentInterfaceImp) UpdateTraineeComment(lessonID string, mapUpdates map[string]interface{}) error {
	cli := db.Get()
	return cli.Table(coach_client_trainee_comment_table_name).Where("lesson_id = ?", lessonID).Updates(mapUpdates).Error
}


const coach_client_monthly_statistics_table_name = "coach_client_monthly_statistics"

func (imp *CoachClientMonthlyStatisticInterfaceImp) GetItem(coachId int, monthBegTs int64) (*model.CoachMonthlyStatisticModel, error) {
	var err error
	var stCoachMonthlyStatisticModel = new(model.CoachMonthlyStatisticModel)
	cli := db.Get()
	err = cli.Table(coach_client_monthly_statistics_table_name).Where("coach_id = ? AND month_beg_ts = ?", coachId, monthBegTs).First(stCoachMonthlyStatisticModel).Error
	return stCoachMonthlyStatisticModel, err
}

func (imp *CoachClientMonthlyStatisticInterfaceImp) GetAllItem() ([]model.CoachMonthlyStatisticModel, error) {
	var err error
	var vecCoachMonthlyStatisticModel []model.CoachMonthlyStatisticModel
	cli := db.Get()
	err = cli.Table(coach_client_monthly_statistics_table_name).Find(&vecCoachMonthlyStatisticModel).Error
	return vecCoachMonthlyStatisticModel, err
}


func (imp *CoachClientMonthlyStatisticInterfaceImp) AddItem(stCoachMonthlyStatisticModel *model.CoachMonthlyStatisticModel) error {
	cli := db.Get()
	return cli.Table(coach_client_monthly_statistics_table_name).Save(stCoachMonthlyStatisticModel).Error
}

const coach_client_trainee_nickname_table_name = "coach_client_trainee_nickname"

// GetTraineeNickname 查询某教练对某用户的自定义昵称，无记录返回 gorm.ErrRecordNotFound
func (imp *CoachClientTraineeNicknameInterfaceImp) GetTraineeNickname(coachId int, uid int64) (*model.CoachClientTraineeNicknameModel, error) {
	var err error
	var stNickname = new(model.CoachClientTraineeNicknameModel)
	cli := db.Get()
	err = cli.Table(coach_client_trainee_nickname_table_name).Where("coach_id = ? AND trainee_uid = ?", coachId, uid).First(stNickname).Error
	return stNickname, err
}

// GetTraineeNicknameListByCoachId 查询某教练对所有用户的自定义昵称
func (imp *CoachClientTraineeNicknameInterfaceImp) GetTraineeNicknameListByCoachId(coachId int) ([]model.CoachClientTraineeNicknameModel, error) {
	var err error
	var vecNickname []model.CoachClientTraineeNicknameModel
	cli := db.Get()
	err = cli.Table(coach_client_trainee_nickname_table_name).Where("coach_id = ?", coachId).Find(&vecNickname).Error
	return vecNickname, err
}

// GetAllTraineeNickname 查询全量的教练自定义用户昵称
func (imp *CoachClientTraineeNicknameInterfaceImp) GetAllTraineeNickname() ([]model.CoachClientTraineeNicknameModel, error) {
	var err error
	var vecNickname []model.CoachClientTraineeNicknameModel
	cli := db.Get()
	err = cli.Table(coach_client_trainee_nickname_table_name).Find(&vecNickname).Error
	return vecNickname, err
}

// UpsertTraineeNickname 每个教练对每个用户只有一条记录：已存在则更新昵称，不存在则新增
func (imp *CoachClientTraineeNicknameInterfaceImp) UpsertTraineeNickname(coachId int, uid int64, nickname string) error {
	cli := db.Get()
	nowTs := time.Now().Unix()

	var stExist model.CoachClientTraineeNicknameModel
	err := cli.Table(coach_client_trainee_nickname_table_name).Where("coach_id = ? AND trainee_uid = ?", coachId, uid).First(&stExist).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	// 已存在 -> 更新
	if err == nil {
		mapUpdates := map[string]interface{}{}
		mapUpdates["nickname"] = nickname
		mapUpdates["update_ts"] = nowTs
		return cli.Table(coach_client_trainee_nickname_table_name).
			Where("coach_id = ? AND trainee_uid = ?", coachId, uid).Updates(mapUpdates).Error
	}

	// 不存在 -> 新增
	stNew := &model.CoachClientTraineeNicknameModel{
		CoachId:    coachId,
		TraineeUid: uid,
		Nickname:   nickname,
		CreateTs:   nowTs,
		UpdateTs:   nowTs,
	}
	return cli.Table(coach_client_trainee_nickname_table_name).Create(stNew).Error
}
