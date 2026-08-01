package dao

import (
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

const physical_assessment_report_table_name = "physical_assessment_reports"

// 由后台到特定时间节点，自动触发，会创建对应的报告，记录解锁时间，变更状态为已解锁，未提交
// CreatePhysicalAssessmentReport 创建单个阶段的报告记录。（待解锁 --> 已解锁，未提交）
func (imp *PhysicalAssessmentReportInterfaceImp) CreatePhysicalAssessmentReport(report *model.PhysicalAssessmentReportModel) error {
	return db.Get().Table(physical_assessment_report_table_name).Create(report).Error
}

// 由前端提交变更
// SubmitPhysicalAssessmentReport 允许未提交和已提交报告保存更新；待解锁报告不可提交。
func (imp *PhysicalAssessmentReportInterfaceImp) SubmitPhysicalAssessmentReport(reportID string, submitTs int64, mapUpdates map[string]interface{}) error {
	if mapUpdates == nil {
		mapUpdates = make(map[string]interface{})
	}
	mapUpdates["status"] = model.Enum_PhysicalReportStatus_Submitted
	mapUpdates["submit_ts"] = submitTs
	mapUpdates["update_ts"] = submitTs

	return db.Get().Table(physical_assessment_report_table_name).
		Model(&model.PhysicalAssessmentReportModel{}).
		Where("report_id = ? AND status IN (?, ?)", reportID,
			model.Enum_PhysicalReportStatus_Unsubmitted,
			model.Enum_PhysicalReportStatus_Submitted).
		Updates(mapUpdates).Error
}

// GetPhysicalAssessmentReport 获取课包指定阶段的报告。
func (imp *PhysicalAssessmentReportInterfaceImp) GetPhysicalAssessmentReport(packageID string, reportType int) (*model.PhysicalAssessmentReportModel, error) {
	report := new(model.PhysicalAssessmentReportModel)
	err := db.Get().Table(physical_assessment_report_table_name).
		Where("package_id = ? AND report_type = ?", packageID, reportType).
		First(report).Error
	return report, err
}

// GetPhysicalAssessmentReportByReportID 根据报告 ID 获取报告详情。
func (imp *PhysicalAssessmentReportInterfaceImp) GetPhysicalAssessmentReportByReportID(reportID string) (*model.PhysicalAssessmentReportModel, error) {
	report := new(model.PhysicalAssessmentReportModel)
	err := db.Get().Table(physical_assessment_report_table_name).
		Where("report_id = ?", reportID).
		First(report).Error
	return report, err
}

// GetPhysicalAssessmentReportListByPackageID 获取课包下全部阶段的报告。
func (imp *PhysicalAssessmentReportInterfaceImp) GetPhysicalAssessmentReportListByPackageID(packageID string) ([]model.PhysicalAssessmentReportModel, error) {
	var reports []model.PhysicalAssessmentReportModel
	err := db.Get().Table(physical_assessment_report_table_name).
		Where("package_id = ?", packageID).
		Order("report_type ASC").
		Find(&reports).Error
	return reports, err
}

// UnlockPhysicalAssessmentReport 仅允许待解锁变为未提交，确保解锁状态不会回退。
func (imp *PhysicalAssessmentReportInterfaceImp) UnlockPhysicalAssessmentReport(packageID string, reportType int, unlockTs int64) error {
	mapUpdates := map[string]interface{}{
		"status":    model.Enum_PhysicalReportStatus_Unsubmitted,
		"unlock_ts": unlockTs,
		"update_ts": unlockTs,
	}
	return db.Get().Table(physical_assessment_report_table_name).
		Model(&model.PhysicalAssessmentReportModel{}).
		Where("package_id = ? AND report_type = ? AND status = ?", packageID, reportType, model.Enum_PhysicalReportStatus_Locked).
		Updates(mapUpdates).Error
}
