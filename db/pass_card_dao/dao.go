package pass_card_dao

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/pass_card_model"
	"time"
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

func (imp *PassCardAppointmentInterfaceImp) SetAppointmentBookedNew(uid int64, appointmentID int, courseId int, gymId int) (error, pass_card_model.PassCardAppointmentModel) {
	var err error
	cli := db.Get().Table(pass_card_gym_appointments_tableName)

	// 先获取再更新的原子操作
	var appointmentModel pass_card_model.PassCardAppointmentModel
	err = cli.Transaction(func(tx *gorm.DB) error {
		// 获取用户记录
		if err := tx.First(&appointmentModel, "appointment_id = ?", appointmentID).Error; err != nil {
			fmt.Printf("get err, uid:%d appointmentID:%d\n", uid, appointmentID)
			return err
		}

		if appointmentModel.Status == pass_card_model.Enum_PassCardAppointment_Status_UnAvailable {
			return errors.New("book status unavailable")
		}
		if uint32(len(appointmentModel.BookedUids)) >= appointmentModel.MaxBookCnt {
			return errors.New("book maxcnt reaclimit")
		}
		for _, v := range appointmentModel.BookedUids {
			if v.Uid == uid {
				return errors.New("user already book")
			}
		}
		appointmentModel.BookedUids = append(appointmentModel.BookedUids, pass_card_model.UserID{Uid: uid})

		// 更新用户记录
		mapUpdates := map[string]interface{}{}
		if uint32(len(appointmentModel.BookedUids)) >= appointmentModel.MaxBookCnt {
			mapUpdates["status"] = pass_card_model.Enum_PassCardAppointment_Status_UnAvailable
		}
		mapUpdates["booked_uids"] = appointmentModel.BookedUids
		mapUpdates["update_ts"] = time.Now().Unix()

		if uint32(len(appointmentModel.BookedUids)) >= appointmentModel.MaxBookCnt {
			appointmentModel.Status = pass_card_model.Enum_PassCardAppointment_Status_UnAvailable
		}
		appointmentModel.UpdateTs = time.Now().Unix()
		appointmentModel.GymId = gymId

		// 更新用户数据，使用 Update 方法
		if err := tx.Model(&pass_card_model.PassCardAppointmentModel{}).Where("appointment_id = ?", appointmentID).Updates(mapUpdates).Error; err != nil {
			fmt.Printf("update err, uid:%d appointmentID:%d mapUpdates:%+v\n", uid, appointmentID, mapUpdates)
			tx.Rollback()
			return err
		}
		return nil
	})
	return err, appointmentModel
}
