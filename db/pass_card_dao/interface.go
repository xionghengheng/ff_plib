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

	// 获取某个门店的时间表
	GetAppointmentScheduleFromBegTs(gymId int, dayBegTs int64) ([]pass_card_model.PassCardAppointmentModel, error)

	// 通过预约id获取预约详情信息
	GetAppointmentById(appointmentID int) (pass_card_model.PassCardAppointmentModel, error)

	// 设置已预约状态
	SetAppointmentBooked(uid int64, appointmentID int, gymId int) (pass_card_model.PassCardAppointmentModel, error)

	// 查询场地某一天的预约时间表
	GetAppointmentScheduleOneDay(gymId int, dayBegTs int64) ([]pass_card_model.PassCardAppointmentModel, error)

	// 设置场地可预约时间
	SetAppointmentSchedule(stPassCardAppointmentModel pass_card_model.PassCardAppointmentModel) error

	// 用户取消约课
	CancelAppointmentBooked(uid int64, lessonID string, appointmentID int) error
}

type PassCardAppointmentInterfaceImp struct{}

var ImpAppointment PassCardAppointmentInterface = &PassCardAppointmentInterfaceImp{}

// 课数据模型接口
type PassCardLessonInterface interface {

	// 通过课程id获取课程详情
	GetSingleLessonById(uid int64, lessonId string) (pass_card_model.LessonModel, error)

	// 根据uid拉取课程列表，根据创建时间降序拉取
	GetLessonListByUid(uid int64, ceateTs int64, status int) ([]pass_card_model.LessonModel, error)

	// 根据门店id拉取课程列表，根据创建时间降序拉取
	GetLessonListByGymId(gymId int, ceateTs int64, status int) ([]pass_card_model.LessonModel, error)

	//// 根据uid拉取预约中的课程列表
	//GetScheduledLessonListByUid(uid int64, ceateTs int64) ([]pass_card_model.LessonModel, error)
	//
	//// 根据uid拉取已完成的课程列表
	//GetCompletedLessonListByUid(uid int64, ceateTs int64) ([]pass_card_model.LessonModel, error)
	//
	//// 根据uid拉取已取消的课程列表
	//GetCancelLessonListByUid(uid int64, ceateTs int64) ([]pass_card_model.LessonModel, error)

	// 根据uid，查询某一天的预约时间表
	GetLessonsOneDay(uid int64, dayBegTs int64) ([]pass_card_model.LessonModel, error)

	// 创建课程
	AddLesson(lesson *pass_card_model.LessonModel) error

	// 更新课程
	UpdateLesson(uid int64, lessonId string, mapUpdates map[string]interface{}) error

	// 获取已经过了结束时间但没有完成的课程，后台需要扫描设置为完成
	GetLessonListNotFinish(nowTs int64, limit int) ([]pass_card_model.LessonModel, error)
	GetLessonListNotFinishAndNotSendGoMsg(nowTs int64, limit int) ([]pass_card_model.LessonModel, error)
}

type PassCardLessonInterfaceImp struct{}

// Imp 实现实例
var ImpPassCardLesson PassCardLessonInterface = &PassCardLessonInterfaceImp{}
