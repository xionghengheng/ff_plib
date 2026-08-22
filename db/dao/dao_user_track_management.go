package dao

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/xionghengheng/ff_plib/db"
	"github.com/xionghengheng/ff_plib/db/model"
)

const (
	user_track_table_name      = "user_track"
	user_track_node_table_name = "user_track_node"
)

// CreateUserTrack 创建客资基础档案。
func (imp *UserTrackManagementInterfaceImp) CreateUserTrack(track *model.UserTrackModel) error {
	return db.Get().Table(user_track_table_name).Create(track).Error
}

// GetUserTrack 根据客资 ID 查询基础档案。
func (imp *UserTrackManagementInterfaceImp) GetUserTrack(trackID int64) (*model.UserTrackModel, error) {
	track := new(model.UserTrackModel)
	err := db.Get().Table(user_track_table_name).
		Where("track_id = ?", trackID).
		First(track).Error
	return track, err
}

// GetUserTrackByWechatNo 根据微信号查询基础档案。
func (imp *UserTrackManagementInterfaceImp) GetUserTrackByWechatNo(wechatNo string) (*model.UserTrackModel, error) {
	track := new(model.UserTrackModel)
	err := db.Get().Table(user_track_table_name).
		Where("wechat_no = ?", wechatNo).
		First(track).Error
	return track, err
}

// UpdateUserTrack 根据客资 ID 更新基础档案。
func (imp *UserTrackManagementInterfaceImp) UpdateUserTrack(trackID int64, updates map[string]interface{}) error {
	return db.Get().Table(user_track_table_name).
		Model(&model.UserTrackModel{}).
		Where("track_id = ?", trackID).
		Updates(updates).Error
}

// GetUnfinishedUserTrackList 按创建时间倒序游标分页查询未完成客资。
// 使用 track_id 作为相同 created_ts 时的稳定排序条件。
// stage <= 0 查询全部未完成阶段，lastCreatedTs <= 0 取第一页。
func (imp *UserTrackManagementInterfaceImp) GetUnfinishedUserTrackList(stage int, lastCreatedTs int64, lastTrackID int64, limit int) ([]model.UserTrackModel, error) {
	var tracks []model.UserTrackModel
	cli := db.Get()
	tx := cli.Table(user_track_table_name).
		Where("stage < ?", model.Enum_Track_Stage_DealDone)
	if stage > 0 {
		tx = tx.Where("stage = ?", stage)
	}
	if lastCreatedTs > 0 {
		tx = tx.Where(
			"(created_ts < ?) OR (created_ts = ? AND track_id < ?)",
			lastCreatedTs,
			lastCreatedTs,
			lastTrackID,
		)
	}
	err := tx.Order("created_ts DESC, track_id DESC").Limit(limit).Find(&tracks).Error
	return tracks, err
}

// GetAllUserTrackList 按创建时间倒序游标分页查询全部客资，不过滤 stage。
// 使用 track_id 作为相同 created_ts 时的稳定排序条件。
// lastCreatedTs <= 0 取第一页。
func (imp *UserTrackManagementInterfaceImp) GetAllUserTrackList(lastCreatedTs int64, lastTrackID int64, limit int) ([]model.UserTrackModel, error) {
	var tracks []model.UserTrackModel
	cli := db.Get()
	tx := cli.Table(user_track_table_name)
	if lastCreatedTs > 0 {
		tx = tx.Where(
			"(created_ts < ?) OR (created_ts = ? AND track_id < ?)",
			lastCreatedTs,
			lastCreatedTs,
			lastTrackID,
		)
	}
	err := tx.Order("created_ts DESC, track_id DESC").Limit(limit).Find(&tracks).Error
	return tracks, err
}

// CreateUserTrackNode 创建客资流转节点。
func (imp *UserTrackManagementInterfaceImp) CreateUserTrackNode(node *model.UserTrackNodeModel) error {
	return db.Get().Table(user_track_node_table_name).Create(node).Error
}

// GetUserTrackNode 查询客资指定阶段的节点。
func (imp *UserTrackManagementInterfaceImp) GetUserTrackNode(trackID int64, stage int) (*model.UserTrackNodeModel, error) {
	node := new(model.UserTrackNodeModel)
	err := db.Get().Table(user_track_node_table_name).
		Where("track_id = ? AND stage = ?", trackID, stage).
		First(node).Error
	return node, err
}

// GetUserTrackNodeList 查询客资的全部流转节点。
func (imp *UserTrackManagementInterfaceImp) GetUserTrackNodeList(trackID int64) ([]model.UserTrackNodeModel, error) {
	var nodes []model.UserTrackNodeModel
	err := db.Get().Table(user_track_node_table_name).
		Where("track_id = ?", trackID).
		Order("stage ASC, node_create_ts ASC").
		Find(&nodes).Error
	return nodes, err
}

// UpdateUserTrackNode 更新客资指定阶段的节点。
func (imp *UserTrackManagementInterfaceImp) UpdateUserTrackNode(trackID int64, stage int, updates map[string]interface{}) error {
	return db.Get().Table(user_track_node_table_name).
		Model(&model.UserTrackNodeModel{}).
		Where("track_id = ? AND stage = ?", trackID, stage).
		Updates(updates).Error
}

// AppendFollowUpRecord 在事务中锁定节点行并追加跟进记录，避免并发写入丢失。
func (imp *UserTrackManagementInterfaceImp) AppendFollowUpRecord(trackID int64, stage int, record model.FollowUpRecord) error {
	return db.Get().Transaction(func(tx *gorm.DB) error {
		var node model.UserTrackNodeModel
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Table(user_track_node_table_name).
			Where("track_id = ? AND stage = ?", trackID, stage).
			First(&node).Error; err != nil {
			return err
		}

		records, err := model.GetFollowUpRecords(node.Data)
		if err != nil {
			return err
		}
		records = append(records, record)

		data, err := json.Marshal(records)
		if err != nil {
			return err
		}

		return tx.Table(user_track_node_table_name).
			Model(&model.UserTrackNodeModel{}).
			Where("track_id = ? AND stage = ?", trackID, stage).
			Updates(map[string]interface{}{
				"data":            string(data),
				"node_updated_ts": record.FollowTs,
			}).Error
	})
}
