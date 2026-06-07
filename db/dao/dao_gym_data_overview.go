package dao

import (
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

const gym_data_overview_table_name = "gym_data_overview"

func (imp *GymDataOverviewInterfaceImp) GetItem(gymId int, monthBegTs int64) (*model.GymDataOverviewModel, error) {
	var err error
	var stGymDataOverviewModel = new(model.GymDataOverviewModel)
	cli := db.Get()
	err = cli.Table(gym_data_overview_table_name).Where("gym_id = ? AND month_beg_ts = ?", gymId, monthBegTs).First(stGymDataOverviewModel).Error
	return stGymDataOverviewModel, err
}

func (imp *GymDataOverviewInterfaceImp) AddItem(stGymDataOverviewModel *model.GymDataOverviewModel) error {
	cli := db.Get()
	return cli.Table(gym_data_overview_table_name).Save(stGymDataOverviewModel).Error
}
