package dao

import (
	"github.com/xionghengheng/ff_plib/db/model"
)

// UserInterface 用户数据模型接口
type UserInterface interface {
	//根据uid获取用户信息
	GetUser(uid int64) (*model.UserInfoModel, error)

	//根据openid获取用户信息
	GetUserByOpenId(openid string) (*model.UserInfoModel, error)

	//根据wx安全回调的traceid获取用户信息
	GetUserByTraceId(traceid string) (*model.UserInfoModel, error)

	//插入用户信息
	UpsertUser(user *model.UserInfoModel) error

	//更新用户信息
	UpdateUserInfo(uid int64, mapUpdates map[string]interface{}) error

	//更新用户电话号码
	UpdateUserPhone(uid int64, phone string) error

	//删除用户信息（慎用，用户后台清空数据）
	RemoveUser(openid string) error

	//根据uid获取用户信息
	GetAllUser() ([]model.UserInfoModel, error)
}

// UserInterfaceImp 用户数据模型实现
type UserInterfaceImp struct{}

// Imp 实现实例
var ImpUser UserInterface = &UserInterfaceImp{}

// SmsInterface 用户数据模型接口
type SmsInterface interface {
	GetCode(uniqid string) (*model.SmsVerificationCodeModel, error)
	AddCode(code *model.SmsVerificationCodeModel) error
	UpdateCodeUserStatus(uniqid string) error
}

// SmsInterfaceImp 用户数据模型实现
type SmsInterfaceImp struct{}

// Imp 实现实例
var ImpSms SmsInterface = &SmsInterfaceImp{}

// GymInterface 用户数据模型接口
type GymInterface interface {
	GetGymList() ([]model.GymInfoModel, error)
	GetGymInfoByGymId(gymId int) (model.GymInfoModel, error)
}

// GymInterfaceImp 用户数据模型实现
type GymInterfaceImp struct{}

// Imp 实现实例
var ImpGym GymInterface = &GymInterfaceImp{}

// CourseInterface 用户数据模型接口
type CourseInterface interface {
	GetCourseList() ([]model.CourseModel, error)
	GetCourseById(id int) (*model.CourseModel, error)
}

// CourseInterfaceImp 用户数据模型实现
type CourseInterfaceImp struct{}

// Imp 实现实例
var ImpCourse CourseInterface = &CourseInterfaceImp{}

// CoachInterface 教练数据模型接口
type CoachInterface interface {

	//通过教练id获取教练详细信息
	GetCoachById(id int) (*model.CoachModel, error)

	GetCoachAll() ([]model.CoachModel, error)

	//根据场地id+优先级 获取教练列表
	GetCoachListByGymId(gymId int) ([]model.CoachModel, error)
}

// CoachInterfaceImp 教练数据模型实现
type CoachInterfaceImp struct{}

// Imp 实现实例
var ImpCoach CoachInterface = &CoachInterfaceImp{}

// CoursePackageInterface 课包数据模型接口
type CoursePackageInterface interface {

	//根据uid获取免费课包
	GetTrailCoursePackage(uid int64) (*model.CoursePackageModel, error)

	//根据uid获取付费课包
	GetPayCoursePackageList(uid int64) ([]model.CoursePackageModel, error)

	//按时间降序拉取某个uid的课包列表
	GetCoursePackageListByUid(uid int64) ([]model.CoursePackageModel, error)

	//根据课包id获取课包详情
	GetCoursePackageById(packageId string) (*model.CoursePackageModel, error)

	//按时间降序拉取某个教练的所有课课包列表
	GetAllCoursePackageListByCoachId(coachId int, limit int) ([]model.CoursePackageModel, error)

	//按时间降序拉取某个教练的体验课课包列表
	GetTrailCoursePackageListByCoachId(coachId int, limit int) ([]model.CoursePackageModel, error)

	//按时间降序拉取某个教练的付费课课包列表
	GetPayCoursePackageListByCoachId(coachId int, limit int) ([]model.CoursePackageModel, error)

	//按最后一节课上课的升序，拉取某个教练的所有课课包列表
	GetListByCoachIdAndLastFinishLessonTs(coachId int, limit int) ([]model.CoursePackageModel, error)

	//按时间降序拉取某个教练下，用户购买的的所有课包列表
	GetAllPackageListByCoachIdAndUid(coachId int, uid int64) ([]model.CoursePackageModel, error)

	//按时间降序拉取某个教练下，用户购买的的付费课包列表
	GetPayCoursePackageListByCoachIdAndUid(coachId int, uid int64) ([]model.CoursePackageModel, error)

	//预约成功后，扣减一节课时
	SubCourseCnt(packageId string) error

	//总数和剩余数都会添加
	AddCourseCnt(packageId string, cnt int) error

	//归还剩余计数（因为取消或旷课 归还剩余次数）
	AddRemainCourseCnt(packageId string, cnt int) error

	//发放课包到用户资产
	AddCoursePackage2Uid(stCoursePackageModel *model.CoursePackageModel) error

	//场地id、教练id和课程id一致，则直接走续费逻辑
	FindSamePackage(uid int64, gymId int, coachId int, courseId int) (string, error)

	//新用户可以更新体验课课包里的教练or场地【谨慎使用】
	UpdateCoursePackage(uid int64, packageId string, mapUpdates map[string]interface{}) error

	//获取所有课包，通过创建时间来分页
	GetAllCoursePackageList(ts int64) ([]model.CoursePackageModel, error)
}

// CoursePackageInterfaceImp 课包数据模型实现
type CoursePackageInterfaceImp struct{}

// Imp 实现实例
var ImpCoursePackage CoursePackageInterface = &CoursePackageInterfaceImp{}

// GetSingleLessonListByPackageId 课包单次课数据模型接口
type CoursePackageSingleLessonInterface interface {
	GetSingleLessonById(uid int64, lessonId string) (*model.CoursePackageSingleLessonModel, error)
	GetSingleLessonByAppointmentId(uid int64, appointmentID int) (*model.CoursePackageSingleLessonModel, error)
	GetSingleLessonListByPackageId(uid int64, packageId string) ([]model.CoursePackageSingleLessonModel, error)
	AddSingleLesson2Package(stCoursePackageSingleLessonModel *model.CoursePackageSingleLessonModel) error
	UpdateSingleLesson(uid int64, lessonId string, mapUpdates map[string]interface{}) error
	GetSingleLessonListNotFinish(nowTs int64, limit int) ([]model.CoursePackageSingleLessonModel, error)
	GetSingleLessonListMissed(limit int) ([]model.CoursePackageSingleLessonModel, error)
	GetTodaySingleLessonListNotSendMsgGoLesson(ts int64, limit int) ([]model.CoursePackageSingleLessonModel, error)
	GetCompletedSingleLessonListByCoachId(coachId int, uBegTs int64) ([]model.CoursePackageSingleLessonModel, error)

	//获取所有次课信息，通过创建时间来分页
	GetAllSingleLessonList(createTs int64) ([]model.CoursePackageSingleLessonModel, error)
}

// CoursePackageSingleLessonInterfaceImp 课包单次课数据模型实现
type CoursePackageSingleLessonInterfaceImp struct{}

// Imp 实现实例
var ImpCoursePackageSingleLesson CoursePackageSingleLessonInterface = &CoursePackageSingleLessonInterfaceImp{}

// PaymentOrderInterface 订单相关数据模型接口
type PaymentOrderInterface interface {

	//添加订单
	AddOrder(stPaymentOrderModel model.PaymentOrderModel) error

	//获取订单信息
	GetOrderById(orderId string, uid int64) (*model.PaymentOrderModel, error)

	//获取订单信息
	GetOrder(orderId string) (*model.PaymentOrderModel, error)

	//订单支付成功，更新订单数据
	UpdateOrderSucc(orderId string, uid int64, mapUpdates map[string]interface{}) error

	//获取订单列表，按时间降序排列
	GetOrderList(uid int64) ([]model.PaymentOrderModel, error)

	//从某个时间点开始，通过教练id，获取全量属于该教练的订单列表
	GetOrderListByCoachId(coachId int, begTs int64) ([]model.PaymentOrderModel, error)

	//通过课包id获取退款订单信息
	GetRefundOrderByPackageId(uid int64, packageId string) (*model.PaymentOrderModel, error)

	//通过课包id获取订单信息
	GetOrderByPackageId(uid int64, packageId string) ([]model.PaymentOrderModel, error)
}

// PaymentOrderInterfaceImp
type PaymentOrderInterfaceImp struct{}

// Imp 实现实例
var ImpPaymentOrder PaymentOrderInterface = &PaymentOrderInterfaceImp{}

// AppointmentInterface 预约课程的数据模型接口
type AppointmentInterface interface {

	//查询教练预约时间表，从某一天的零点时间戳开始
	GetAppointmentScheduleFromBegTs(gymId int, coachid int, dayBegTs int64) ([]model.CoachAppointmentModel, error)

	//查询教练预约时间表，从某一天的零点时间戳开始(已经有用户预约的)
	GetAppointmentScheduleHasUidFromBegTs(gymId int, coachid int, dayBegTs int64) ([]model.CoachAppointmentModel, error)


	//查询教练某一天的预约时间表
	GetAppointmentScheduleOneDay(gymId int, coachid int, dayBegTs int64) ([]model.CoachAppointmentModel, error)



	//查询用户预约时间记录
	GetUserAppointmentRecordOneDay(uid int64, dayBegTs int64) ([]model.CoachAppointmentModel, error)
	GetUserAppointmentRecordFromBegTs(uid int64, dayBegTs int64, limit int) ([]model.CoachAppointmentModel, error)
	GetUserAppointmentRecord(uid int64, limit int) ([]model.CoachAppointmentModel, error)

	//通过预约id获取预约详情信息
	GetAppointmentById(appointmentID int) (*model.CoachAppointmentModel, error)

	DelAppointmentByCoach(appointmentID int, coachId int) (error)

	GetAppointmentByBegTsAndEndTs(gymId int, coachid int, begTs int64, endTs int64) (*model.CoachAppointmentModel, error)

	//用户发起约课
	SetAppointmentBooked(uid int64, appointmentID int, courseId int) (error, model.CoachAppointmentModel)

	//用户取消约课
	CancelAppointmentBooked(uid int64, appointmentID int) error

	//教练端，设置可预约时间
	SetAppointmentSchedule(stCoachAppointmentModel model.CoachAppointmentModel) error
}

// AppointmentInterfaceImp
type AppointmentInterfaceImp struct{}

// Imp 实现实例
var ImpAppointment AppointmentInterface = &AppointmentInterfaceImp{}

// InvitationCodeInterface 预约课程的数据模型接口
type InvitationCodeInterface interface {

	//获取邀请码
	GetCode(code string) (*model.InvitationCodeModel, error)

	//更新邀请码
	UpdateCode(code string, uid int64) error

	//生成邀请码
	AddCode(stInvitationCodeModel *model.InvitationCodeModel) error
}

// InvitationCodeInterfaceImp
type InvitationCodeInterfaceImp struct{}

// Imp 实现实例
var ImpInvitationCode InvitationCodeInterface = &InvitationCodeInterfaceImp{}

// CoachPageBannerInterface 私教tab页顶部banner
type CoachPageBannerInterface interface {
	GetBannerList() ([]model.CoachPageBannerModel, error)
}

// CoachPageBannerInterfaceImp
type CoachPageBannerInterfaceImp struct{}

// Imp 实现实例
var ImpCoachPageBanner CoachPageBannerInterface = &CoachPageBannerInterfaceImp{}
