package dao

import "github.com/xionghengheng/ff_plib/db/model"

// GymDataOverviewInterface 门店端数据概览
type GymDataOverviewInterface interface {
	GetItem(gymId int, monthBegTs int64) (*model.GymDataOverviewModel, error)
	AddItem(stGymDataOverviewModel *model.GymDataOverviewModel) error
}

// GymDataOverviewInterfaceImp
type GymDataOverviewInterfaceImp struct{}

// Imp 实现实例
var ImpGymDataOverview GymDataOverviewInterface = &GymDataOverviewInterfaceImp{}
