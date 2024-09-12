package dao

import (
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

const report_table_expourse_name = "report_table_expourse"
const report_table_write_name = "report_table_write"
const report_table_coach_client_expourse_name = "report_table_coach_client_expourse"
const report_table_coach_client_write_name = "report_table_coach_client_write"

func (imp *ReportInterfaceImp) DoReport(stReportItem model.ReportModel) (error) {
	cli := db.Get()
	if stReportItem.IsCoach{
		if stReportItem.ActionID == 101{
			return cli.Table(report_table_coach_client_expourse_name).Save(stReportItem).Error
		}else if stReportItem.ActionID == 102{
			return cli.Table(report_table_coach_client_write_name).Save(stReportItem).Error
		}
	}else{
		if stReportItem.ActionID == 101{
			return cli.Table(report_table_expourse_name).Save(stReportItem).Error
		}else if stReportItem.ActionID == 102{
			return cli.Table(report_table_write_name).Save(stReportItem).Error
		}
	}
	return nil
}

func (imp *ReportInterfaceImp) GetPageReport(strPageId string, begTs int64, endTs int64) ([]model.ReportModel, error) {
	var err error
	var vecReportModel []model.ReportModel
	cli := db.Get()
	err = cli.Table(report_table_expourse_name).Where("page_id = ? AND report_time >= ? AND report_time <= ?", strPageId, begTs, endTs).Order("report_time DESC").Find(&vecReportModel).Error
	return vecReportModel, err
}

func (imp *ReportInterfaceImp) GetButtonReport(strButtonId string, begTs int64, endTs int64) ([]model.ReportModel, error) {
	var err error
	var vecReportModel []model.ReportModel
	cli := db.Get()
	err = cli.Table(report_table_write_name).Where("item_id = ? AND report_time >= ? AND report_time <= ?", strButtonId, begTs, endTs).Order("report_time DESC").Find(&vecReportModel).Error
	return vecReportModel, err
}
