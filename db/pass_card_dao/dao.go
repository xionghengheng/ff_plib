package pass_card_dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
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

func (imp *PassCardAppointmentInterfaceImp) SetAppointmentBooked(uid int64, appointmentID int, gymId int) (pass_card_model.PassCardAppointmentModel, error) {
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

		var vecBookedUserID []pass_card_model.UserID
		err := json.Unmarshal([]byte(appointmentModel.BookedUids), vecBookedUserID)
		if err != nil {
			return errors.New(fmt.Sprintf("json Unmarshal booked_uids err, uid:%d appointmentID:%d BookedUids:%+v\n", uid, appointmentID, appointmentModel.BookedUids))
		}

		if uint32(len(vecBookedUserID)) >= appointmentModel.MaxBookCnt {
			return errors.New("book maxcnt reaclimit")
		}
		for _, v := range vecBookedUserID {
			if v.Uid == uid {
				return errors.New("user already book")
			}
		}
		vecBookedUserID = append(vecBookedUserID, pass_card_model.UserID{Uid: uid})

		// 更新用户记录
		mapUpdates := map[string]interface{}{}
		if uint32(len(vecBookedUserID)) >= appointmentModel.MaxBookCnt {
			mapUpdates["status"] = pass_card_model.Enum_PassCardAppointment_Status_UnAvailable
		}

		bookedUidsJSON, err := json.Marshal(vecBookedUserID)
		if err != nil {
			return errors.New(fmt.Sprintf("json marshal booked_uids err, uid:%d appointmentID:%d\n", uid, appointmentID))
		}
		mapUpdates["booked_uids"] = string(bookedUidsJSON)
		mapUpdates["update_ts"] = time.Now().Unix()

		if uint32(len(vecBookedUserID)) >= appointmentModel.MaxBookCnt {
			appointmentModel.Status = pass_card_model.Enum_PassCardAppointment_Status_UnAvailable
		}
		appointmentModel.BookedUids = string(bookedUidsJSON)
		appointmentModel.UpdateTs = time.Now().Unix()

		// 更新用户数据，使用 Update 方法
		if err := tx.Model(&pass_card_model.PassCardAppointmentModel{}).Where("appointment_id = ?", appointmentID).Updates(mapUpdates).Error; err != nil {
			fmt.Printf("update err, err:%+v uid:%d appointmentID:%d mapUpdates:%+v\n", err, uid, appointmentID, mapUpdates)
			tx.Rollback()
			return err
		}
		fmt.Printf("update succ, uid:%d appointmentID:%d mapUpdates:%+v\n", uid, appointmentID, mapUpdates)
		return nil
	})
	return appointmentModel, err
}

// 用户取消约课，将课程变回可用状态，即所有用户都可预约
func (imp *PassCardAppointmentInterfaceImp) CancelAppointmentBooked(uid int64, lessonID string, appointmentID int) error {
	var err error
	cli := db.Get().Table(pass_card_gym_appointments_tableName)

	// 先获取再更新的原子操作
	err = cli.Transaction(func(tx *gorm.DB) error {
		var appointmentModel pass_card_model.PassCardAppointmentModel
		// 获取用户记录
		if err := tx.First(&appointmentModel, "appointment_id = ?", appointmentID).Error; err != nil {
			fmt.Printf("get err, uid:%d appointmentID:%d\n", uid, appointmentID)
			return err
		}

		var vecBookedUserID []pass_card_model.UserID
		err := json.Unmarshal([]byte(appointmentModel.BookedUids), vecBookedUserID)
		if err != nil {
			return errors.New(fmt.Sprintf("json Unmarshal booked_uids err, uid:%d appointmentID:%d BookedUids:%+v\n", uid, appointmentID, appointmentModel.BookedUids))
		}

		findIdx := -1
		for idx, v := range vecBookedUserID {
			if v.Uid == uid {
				findIdx = idx
			}
		}
		if findIdx == -1 {
			return errors.New("user already cancel")
		}
		vecBookedUserID = append(vecBookedUserID[:findIdx], vecBookedUserID[findIdx+1:]...)

		// 更新用户记录
		bookedUidsJSON, err := json.Marshal(vecBookedUserID)
		if err != nil {
			return errors.New(fmt.Sprintf("json marshal booked_uids err, uid:%d appointmentID:%d\n", uid, appointmentID))
		}

		mapUpdates := map[string]interface{}{}
		mapUpdates["status"] = model.Enum_Appointment_Status_Available
		mapUpdates["booked_uids"] = appointmentModel.BookedUids
		mapUpdates["update_ts"] = time.Now().Unix()
		appointmentModel.Status = model.Enum_Appointment_Status_Available
		appointmentModel.BookedUids = string(bookedUidsJSON)
		appointmentModel.UpdateTs = time.Now().Unix()

		// 更新用户数据，使用 Update 方法
		if err := tx.Model(&pass_card_model.PassCardAppointmentModel{}).Where("appointment_id = ?", appointmentID).Updates(mapUpdates).Error; err != nil {
			fmt.Printf("update err, uid:%d appointmentID:%d mapUpdates:%+v\n", uid, appointmentID, mapUpdates)
			tx.Rollback()
			return err
		}
		return nil
	})
	return err
}

const pass_card_lesson_tableName = "pass_card_lesson"

func (imp *PassCardLessonInterfaceImp) GetSingleLessonById(uid int64, lessonId string) (pass_card_model.LessonModel, error) {
	var err error
	var lesson pass_card_model.LessonModel
	cli := db.Get()
	err = cli.Table(pass_card_lesson_tableName).Where("uid = ? AND lesson_id = ?", uid, lessonId).Find(&lesson).Error
	return lesson, err
}

func (imp *PassCardLessonInterfaceImp) AddLesson(lesson *pass_card_model.LessonModel) error {
	return db.Get().Table(pass_card_lesson_tableName).Save(lesson).Error
}

func (imp *PassCardLessonInterfaceImp) UpdateLesson(uid int64, lessonId string, mapUpdates map[string]interface{}) error {
	cli := db.Get()
	return cli.Table(pass_card_lesson_tableName).Model(&pass_card_model.LessonModel{}).
		Where("uid = ? AND lesson_id = ?", uid, lessonId).Updates(mapUpdates).Error
}
