package dao

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

// ErrCoursePackagePurchaseWhitelistUnavailable 授权不存在、不属于当前用户或已使用。
var ErrCoursePackagePurchaseWhitelistUnavailable = errors.New("course package purchase whitelist is unavailable")

// AddCoursePackagePurchaseWhitelist 新增未使用的授权。
func (imp *CoursePackagePurchaseWhitelistInterfaceImp) AddCoursePackagePurchaseWhitelist(record *model.CoursePackagePurchaseWhitelistModel) error {
	return db.Get().Table("course_package_purchase_whitelist").Create(record).Error
}

// GetUnusedCoursePackagePurchaseWhitelistByUid 查询用户当前未使用的授权。
func (imp *CoursePackagePurchaseWhitelistInterfaceImp) GetUnusedCoursePackagePurchaseWhitelistByUid(uid int64) (*model.CoursePackagePurchaseWhitelistModel, error) {
	record := new(model.CoursePackagePurchaseWhitelistModel)
	result := db.Get().Raw(`
		SELECT *
		FROM course_package_purchase_whitelist
		WHERE unused_uid = ?
		LIMIT 1`, uid).Scan(record)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return record, nil
}

// GetCoursePackagePurchaseWhitelistListByUid 按 ID 倒序游标分页查询用户授权历史。
func (imp *CoursePackagePurchaseWhitelistInterfaceImp) GetCoursePackagePurchaseWhitelistListByUid(uid, lastID int64, limit int) ([]model.CoursePackagePurchaseWhitelistModel, error) {
	var records []model.CoursePackagePurchaseWhitelistModel
	var err error
	if lastID > 0 {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_purchase_whitelist
			WHERE uid = ? AND id < ?
			ORDER BY id DESC
			LIMIT ?`, uid, lastID, limit).Scan(&records).Error
	} else {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_purchase_whitelist
			WHERE uid = ?
			ORDER BY id DESC
			LIMIT ?`, uid, limit).Scan(&records).Error
	}
	return records, err
}

// UseCoursePackagePurchaseWhitelist 在调用方的支付事务中原子消费未使用授权。
func (imp *CoursePackagePurchaseWhitelistInterfaceImp) UseCoursePackagePurchaseWhitelist(tx *gorm.DB, id, uid int64, orderID string, usedTs int64) error {
	if tx.Error != nil {
		return tx.Error
	}
	result := tx.Exec(`
		UPDATE course_package_purchase_whitelist
		SET is_used = 1, order_id = ?, used_ts = ?, updated_ts = ?
		WHERE id = ? AND uid = ? AND is_used = 0`, orderID, usedTs, usedTs, id, uid)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return ErrCoursePackagePurchaseWhitelistUnavailable
	}
	return nil
}
