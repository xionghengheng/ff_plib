package dao

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/xionghengheng/ff_plib/db/model"
)

func kpiMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	t.Helper()
	conn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		t.Fatal(err)
	}
	db.LogMode(false)
	t.Cleanup(func() {
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
		db.Close()
	})
	return db, mock
}
func TestSaveKPITargetAuditRollback(t *testing.T) {
	db, mock := kpiMockDB(t)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT .*kpi_month_target.*FOR UPDATE").WillReturnRows(sqlmock.NewRows([]string{"id", "stat_month", "created_by", "created_ts"}).AddRow(1, "202611", "original", 100))
	mock.ExpectExec("UPDATE .*kpi_month_target").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("INSERT INTO kpi_target_change_log").WillReturnError(errors.New("audit insert failed"))
	mock.ExpectRollback()
	record := &model.KPIMonthTargetModel{StatMonth: "202611", LessonBaselineTarget: 100, LessonSprintTarget: 120, NewPaidUserTarget: 10, GMVTargetFen: 12345, UpdatedBy: "admin", CreatedBy: "new", CreatedTs: 200, UpdatedTs: 200}
	if err := saveKPIMonthTarget(db, record); err == nil {
		t.Fatal("audit failure must roll back")
	}
	if record.CreatedBy != "original" || record.CreatedTs != 100 {
		t.Fatal("creation identity overwritten")
	}
}
func TestSaveKPITargetCreationCommitsAudit(t *testing.T) {
	db, mock := kpiMockDB(t)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT .*kpi_month_target.*FOR UPDATE").WillReturnRows(sqlmock.NewRows([]string{"id"}))
	mock.ExpectExec("INSERT INTO .*kpi_month_target").WillReturnResult(sqlmock.NewResult(8, 1))
	mock.ExpectExec("INSERT INTO kpi_target_change_log").WithArgs(int64(8), 1, "{}", sqlmock.AnyArg(), "admin", "admin", int64(200)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	record := &model.KPIMonthTargetModel{StatMonth: "202611", UpdatedBy: "admin", UpdatedTs: 200}
	if err := saveKPIMonthTarget(db, record); err != nil {
		t.Fatal(err)
	}
	if record.ID != 8 {
		t.Fatal("ID not filled")
	}
}
func TestMutateKPIMonthPreservesSnapshotAndNull(t *testing.T) {
	db, mock := kpiMockDB(t)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT .*kpi_member.*FOR UPDATE").WillReturnRows(sqlmock.NewRows([]string{"id", "member_name", "member_type"}).AddRow(1, "new name", 1))
	mock.ExpectQuery("SELECT .*kpi_member_month_config.*FOR UPDATE").WillReturnRows(sqlmock.NewRows([]string{"id", "member_id", "stat_month", "member_name", "member_type", "renewal_pct", "silent_over_7d_pct"}).AddRow(7, 1, "202611", "old name", 2, 33.33, 0))
	mock.ExpectExec("UPDATE .*kpi_member_month_config").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	result, err := mutateKPIMemberMonthConfig(db, "202611", 1, 200, func(r *model.KPIMemberMonthConfigModel) error {
		if r.MemberName != "old name" || r.MemberType != 2 {
			t.Fatal("snapshot changed")
		}
		r.RenewalPct = nil
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.RenewalPct != nil || result.SilentOver7dPct == nil || *result.SilentOver7dPct != 0 {
		t.Fatal("null/zero lost")
	}
}
func TestMutateKPIMonthRejectedPatchRollsBack(t *testing.T) {
	db, mock := kpiMockDB(t)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT .*kpi_member.*FOR UPDATE").WillReturnRows(sqlmock.NewRows([]string{"id", "member_name", "member_type"}).AddRow(1, "A", 1))
	mock.ExpectQuery("SELECT .*kpi_member_month_config.*FOR UPDATE").WillReturnRows(sqlmock.NewRows([]string{"id"}))
	mock.ExpectRollback()
	_, err := mutateKPIMemberMonthConfig(db, "202611", 1, 200, func(r *model.KPIMemberMonthConfigModel) error { return errors.New("invalid") })
	if err == nil {
		t.Fatal("invalid patch succeeded")
	}
}
