package dao

import "github.com/xionghengheng/ff_plib/db/model"

// DirectTrialBookingInterface 已购体验课直接排课历史数据访问接口。
type DirectTrialBookingInterface interface {
	GetByRequestID(requestID string) (*model.DirectTrialBookingHistoryModel, error)
	Create(record *model.DirectTrialBookingHistoryModel) error
	UpdateResult(id int64, status, lessonID, message string, now int64) error
	GetList(phone string, lastID int64, limit int) ([]model.DirectTrialBookingHistoryModel, error)
}

// DirectTrialBookingInterfaceImp 已购体验课排课历史数据访问实现。
type DirectTrialBookingInterfaceImp struct{}

var ImpDirectTrialBooking DirectTrialBookingInterface = &DirectTrialBookingInterfaceImp{}

var _ DirectTrialBookingInterface = (*DirectTrialBookingInterfaceImp)(nil)
