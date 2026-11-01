package dao

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

func (imp *KPIMemberInterfaceImp) DeleteKPIMember(id int64) error {
	return db.Get().Transaction(func(tx *gorm.DB) error {
		var member model.KPIMemberModel
		if err := tx.Table("kpi_member").Set("gorm:query_option", "FOR UPDATE").Where("id = ?", id).First(&member).Error; err != nil {
			return err
		}
		return tx.Table("kpi_member").Where("id = ?", id).Delete(&model.KPIMemberModel{}).Error
	})
}

func (imp *KPIMonthTargetInterfaceImp) SaveKPIMonthTarget(record *model.KPIMonthTargetModel) error {
	return saveKPIMonthTarget(db.Get(), record)
}

func saveKPIMonthTarget(conn *gorm.DB, record *model.KPIMonthTargetModel) error {
	return conn.Transaction(func(tx *gorm.DB) error {
		var old model.KPIMonthTargetModel
		err := tx.Table("kpi_month_target").Set("gorm:query_option", "FOR UPDATE").Where("stat_month = ?", record.StatMonth).First(&old).Error
		operation, before := 1, []byte("{}")
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return err
		}
		if err == nil {
			operation = 2
			before, err = json.Marshal(old)
			if err != nil {
				return err
			}
			record.ID, record.CreatedBy, record.CreatedTs = old.ID, old.CreatedBy, old.CreatedTs
			err = tx.Table("kpi_month_target").Where("id = ?", old.ID).Save(record).Error
		} else {
			err = tx.Table("kpi_month_target").Create(record).Error
		}
		if err != nil {
			return err
		}
		after, err := json.Marshal(record)
		if err != nil {
			return err
		}
		// 明确使用Exec写入审计记录，避免GORM v1对map Create的不支持。
		return tx.Exec(`INSERT INTO kpi_target_change_log
   (target_id, operation_type, before_data, after_data, operator_account, operator_name, change_reason, created_ts)
   VALUES (?, ?, ?, ?, ?, ?, '', ?)`, record.ID, operation, string(before), string(after), record.UpdatedBy, record.UpdatedBy, record.UpdatedTs).Error
	})
}

func (imp *KPIMemberMonthConfigInterfaceImp) MutateKPIMemberMonthConfig(month string, memberID, now int64, mutate func(*model.KPIMemberMonthConfigModel) error) (*model.KPIMemberMonthConfigModel, error) {
	return mutateKPIMemberMonthConfig(db.Get(), month, memberID, now, mutate)
}

func mutateKPIMemberMonthConfig(conn *gorm.DB, month string, memberID, now int64, mutate func(*model.KPIMemberMonthConfigModel) error) (*model.KPIMemberMonthConfigModel, error) {
	var record model.KPIMemberMonthConfigModel
	err := conn.Transaction(func(tx *gorm.DB) error {
		var member model.KPIMemberModel
		if err := tx.Table("kpi_member").Set("gorm:query_option", "FOR UPDATE").Where("id = ?", memberID).First(&member).Error; err != nil {
			return err
		}
		err := tx.Table("kpi_member_month_config").Set("gorm:query_option", "FOR UPDATE").Where("stat_month = ? AND member_id = ?", month, memberID).First(&record).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return err
		}
		creating := gorm.IsRecordNotFoundError(err)
		if creating {
			record = model.KPIMemberMonthConfigModel{StatMonth: month, MemberID: memberID, MemberName: member.MemberName, MemberType: member.MemberType, CreatedTs: now}
		}
		if err := mutate(&record); err != nil {
			return err
		}
		record.UpdatedTs = now
		if creating {
			return tx.Table("kpi_member_month_config").Create(&record).Error
		}
		return tx.Table("kpi_member_month_config").Where("id = ?", record.ID).Save(&record).Error
	})
	return &record, err
}
