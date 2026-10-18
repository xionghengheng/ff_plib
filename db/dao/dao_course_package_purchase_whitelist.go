package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

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
		WHERE uid = ? AND is_used = 0
		LIMIT 1`, uid).Scan(record)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return record, nil
}

// GetUnusedCoursePackagePurchaseWhitelistByPhone 按手机号查询未使用授权。
func (imp *CoursePackagePurchaseWhitelistInterfaceImp) GetUnusedCoursePackagePurchaseWhitelistByPhone(userPhone string) (*model.CoursePackagePurchaseWhitelistModel, error) {
	record := new(model.CoursePackagePurchaseWhitelistModel)
	result := db.Get().Raw(`
		SELECT *
		FROM course_package_purchase_whitelist
		WHERE user_phone = ? AND is_used = 0
		LIMIT 1`, userPhone).Scan(record)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return record, nil
}

// GetCoursePackagePurchaseWhitelistList 按 ID 倒序游标分页查询全量授权记录。
func (imp *CoursePackagePurchaseWhitelistInterfaceImp) GetCoursePackagePurchaseWhitelistList(lastID int64, limit int) ([]model.CoursePackagePurchaseWhitelistModel, error) {
	var records []model.CoursePackagePurchaseWhitelistModel
	var err error
	if lastID > 0 {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_purchase_whitelist
			WHERE id < ?
			ORDER BY id DESC
			LIMIT ?`, lastID, limit).Scan(&records).Error
	} else {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_purchase_whitelist
			ORDER BY id DESC
			LIMIT ?`, limit).Scan(&records).Error
	}
	return records, err
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

// GetCoursePackagePurchaseWhitelistListByPhone 按手机号、ID 倒序游标分页查询授权历史。
func (imp *CoursePackagePurchaseWhitelistInterfaceImp) GetCoursePackagePurchaseWhitelistListByPhone(userPhone string, lastID int64, limit int) ([]model.CoursePackagePurchaseWhitelistModel, error) {
	var records []model.CoursePackagePurchaseWhitelistModel
	var err error
	if lastID > 0 {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_purchase_whitelist
			WHERE user_phone = ? AND id < ?
			ORDER BY id DESC
			LIMIT ?`, userPhone, lastID, limit).Scan(&records).Error
	} else {
		err = db.Get().Raw(`
			SELECT *
			FROM course_package_purchase_whitelist
			WHERE user_phone = ?
			ORDER BY id DESC
			LIMIT ?`, userPhone, limit).Scan(&records).Error
	}
	return records, err
}

// UpdateCoursePackagePurchaseWhitelist 根据主键 ID 更新指定字段。
func (imp *CoursePackagePurchaseWhitelistInterfaceImp) UpdateCoursePackagePurchaseWhitelist(id int64, updates map[string]interface{}) error {
	return db.Get().Table("course_package_purchase_whitelist").
		Model(&model.CoursePackagePurchaseWhitelistModel{}).
		Where("id = ?", id).
		Updates(updates).Error
}
