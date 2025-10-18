package pass_card_dao

import (
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/pass_card_model"
)

const pass_card_gym_tableName = "pass_card_gym_info"

func (imp *PassCardGymInterfaceImp) GetGymList() ([]pass_card_model.PassCardGymInfoModel, error) {
	var err error
	var vecGyms []pass_card_model.PassCardGymInfoModel
	cli := db.Get()
	err = cli.Table(pass_card_gym_tableName).Find(&vecGyms).Order("gym_id ASC").Error
	return vecGyms, err
}

func (imp *PassCardGymInterfaceImp) GetGymInfoByGymId(gymId int) (pass_card_model.PassCardGymInfoModel, error) {
	var err error
	var stGym pass_card_model.PassCardGymInfoModel
	cli := db.Get()
	err = cli.Table(pass_card_gym_tableName).Where("gym_id = ?", gymId).Find(&stGym).Error
	return stGym, err
}
