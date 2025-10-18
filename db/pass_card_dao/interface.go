package pass_card_dao

import "github.com/xionghengheng/ff_plib/db/pass_card_model"

// GymInterface 用户数据模型接口
type PassCardGymInterface interface {
	GetGymList() ([]pass_card_model.PassCardGymInfoModel, error)
	GetGymInfoByGymId(gymId int) (pass_card_model.PassCardGymInfoModel, error)
}

type PassCardGymInterfaceImp struct{}

// Imp 实现实例
var ImpGym PassCardGymInterface = &PassCardGymInterfaceImp{}
