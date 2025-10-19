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

const pass_card_gym_appointments_tableName = "pass_card_gym_appointments"

func (imp *PassCardAppointmentInterfaceImp) GetAppointmentScheduleOneDay(gymId int, dayBegTs int64) ([]pass_card_model.PassCardAppointmentModel, error) {
	var err error
	var vecCoachAppointmentModel []pass_card_model.PassCardAppointmentModel
	cli := db.Get()
	err = cli.Table(pass_card_gym_appointments_tableName).Where("gym_id = ? AND appointment_date = ?", gymId, dayBegTs).Order("start_time ASC").Find(&vecCoachAppointmentModel).Error
	return vecCoachAppointmentModel, err
}

func (imp *PassCardAppointmentInterfaceImp) SetAppointmentSchedule(stPassCardAppointmentModel pass_card_model.PassCardAppointmentModel) error {
	return db.Get().Table(pass_card_gym_appointments_tableName).Save(stPassCardAppointmentModel).Error
}

func (imp *PassCardAppointmentInterfaceImp) GetAppointmentById(appointmentID int) (pass_card_model.PassCardAppointmentModel, error) {
	var err error
	var appointment pass_card_model.PassCardAppointmentModel
	cli := db.Get()
	err = cli.Table(pass_card_gym_appointments_tableName).Where("appointment_id = ?", appointmentID).First(&appointment).Error
	return appointment, err
}
