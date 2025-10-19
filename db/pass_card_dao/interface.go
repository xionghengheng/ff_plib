package pass_card_dao

import (
	"github.com/xionghengheng/ff_plib/db/pass_card_model"
)

// 通卡场地信息
type PassCardGymInterface interface {
	// 获取全量场地列表
	GetGymList() ([]pass_card_model.PassCardGymInfoModel, error)

	// 根据场地id获取场地详情
	GetGymInfoByGymId(gymId int) (pass_card_model.PassCardGymInfoModel, error)
}

type PassCardGymInterfaceImp struct{}

var ImpGym PassCardGymInterface = &PassCardGymInterfaceImp{}

// 通卡预约信息
type PassCardAppointmentInterface interface {
	// 通过预约id获取预约详情信息
	GetAppointmentById(appointmentID int) (pass_card_model.PassCardAppointmentModel, error)

	// 设置已预约状态
	SetAppointmentBooked(uid int64, appointmentID int, gymId int) (error, pass_card_model.PassCardAppointmentModel)

	// 查询场地某一天的预约时间表
	GetAppointmentScheduleOneDay(gymId int, dayBegTs int64) ([]pass_card_model.PassCardAppointmentModel, error)

	// 设置场地可预约时间
	SetAppointmentSchedule(stPassCardAppointmentModel pass_card_model.PassCardAppointmentModel) error
}

type PassCardAppointmentInterfaceImp struct{}

var ImpAppointment PassCardAppointmentInterface = &PassCardAppointmentInterfaceImp{}
