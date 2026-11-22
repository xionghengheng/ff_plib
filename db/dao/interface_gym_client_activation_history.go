package dao

import "github.com/xionghengheng/ff_plib/db/model"

// GymClientActivationHistoryInterface 门店端成功开通历史数据访问接口。
type GymClientActivationHistoryInterface interface {
	Create(record *model.GymClientActivationHistoryModel) error
	GetList(phone string, lastID int64, limit int) ([]model.GymClientActivationHistoryModel, error)
}

type GymClientActivationHistoryInterfaceImp struct{}

var ImpGymClientActivationHistory GymClientActivationHistoryInterface = &GymClientActivationHistoryInterfaceImp{}
