package dao

import (
	"errors"

	"github.com/xionghengheng/ff_plib/db/model"
)

var (
	ErrGymClientUserNotFound    = errors.New("该手机号尚未注册，请先注册用户账号")
	ErrGymClientAmbiguousPhone  = errors.New("该手机号对应多个账号，请先核对用户信息")
	ErrGymClientGymNotFound     = errors.New("所选门店不存在，请刷新门店列表")
	ErrGymClientRequestConflict = errors.New("请求标识已用于其他开通操作")
)

type GymClientActivationInterface interface {
	// Activate 原子地补齐通卡门店、绑定账号、清除教练身份并保存成功历史。
	// 相同 requestID、手机号、门店、操作者的重试返回原记录，不重新绑定。
	Activate(phone string, gymID int, requestID, operator string) (*model.GymClientActivationModel, error)
	GetList(phone string, lastID int64, limit int) ([]model.GymClientActivationModel, error)
}

type GymClientActivationInterfaceImp struct{}

var ImpGymClientActivation GymClientActivationInterface = &GymClientActivationInterfaceImp{}
