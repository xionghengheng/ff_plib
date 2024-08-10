package dao

import "github.com/xionghengheng/ff_plib/db/model"

// ReportInterface 统一上报
type ReportInterface interface {
	DoReport(stReportItem model.ReportModel) (error)
}

// ReportInterfaceImp
type ReportInterfaceImp struct{}

// Imp 实现实例
var ImpReport ReportInterface = &ReportInterfaceImp{}
