package dao

import "github.com/xionghengheng/ff_plib/db/model"

// PhysicalAssessmentReportInterface 体测报告数据访问接口。
type PhysicalAssessmentReportInterface interface {
	// CreatePhysicalAssessmentReport 创建单个阶段的报告记录。
	CreatePhysicalAssessmentReport(report *model.PhysicalAssessmentReportModel) error

	// GetPhysicalAssessmentReport 获取课包指定阶段的报告。
	GetPhysicalAssessmentReport(packageID string, reportType int) (*model.PhysicalAssessmentReportModel, error)

	// GetPhysicalAssessmentReportListByPackageID 获取课包下全部三个阶段的报告。
	GetPhysicalAssessmentReportListByPackageID(packageID string) ([]model.PhysicalAssessmentReportModel, error)

	// UnlockPhysicalAssessmentReport 将待解锁报告原子地变更为未提交；已解锁或已提交时不回退。
	UnlockPhysicalAssessmentReport(packageID string, reportType int, unlockTs int64) error

	// SubmitPhysicalAssessmentReport 保存报告内容；未提交和已提交的报告均允许更新。
	SubmitPhysicalAssessmentReport(reportID string, submitTs int64, mapUpdates map[string]interface{}) error
}

// PhysicalAssessmentReportInterfaceImp 体测报告数据访问实现。
type PhysicalAssessmentReportInterfaceImp struct{}

// ImpPhysicalAssessmentReport 体测报告 DAO 实例。
var ImpPhysicalAssessmentReport PhysicalAssessmentReportInterface = &PhysicalAssessmentReportInterfaceImp{}
