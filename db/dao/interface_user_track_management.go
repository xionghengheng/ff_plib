package dao

import "github.com/xionghengheng/ff_plib/db/model"

// UserTrackManagementInterface 客资档案与流转节点数据访问接口。
type UserTrackManagementInterface interface {
	// CreateUserTrack 创建客资基础档案。
	CreateUserTrack(track *model.UserTrackModel) error

	// GetUserTrack 根据客资 ID 查询基础档案。
	GetUserTrack(trackID int64) (*model.UserTrackModel, error)

	// GetUserTrackByWechatNo 根据微信号查询基础档案。
	GetUserTrackByWechatNo(wechatNo string) (*model.UserTrackModel, error)

	// GetUserTrackByPhone 根据手机号查询基础档案。
	GetUserTrackByPhone(phone string) (*model.UserTrackModel, error)

	// UpdateUserTrack 根据客资 ID 更新基础档案。
	UpdateUserTrack(trackID int64, updates map[string]interface{}) error

	// GetUnfinishedUserTrackList 按创建时间倒序游标分页查询未完成客资。
	// stage <= 0 查询全部未完成阶段，lastCreatedTs <= 0 取第一页。
	GetUnfinishedUserTrackList(stage int, lastCreatedTs int64, lastTrackID int64, limit int) ([]model.UserTrackModel, error)

	// GetAllUserTrackList 按创建时间倒序游标分页查询全部客资，不过滤 stage。
	// lastCreatedTs <= 0 取第一页。
	GetAllUserTrackList(lastCreatedTs int64, lastTrackID int64, limit int) ([]model.UserTrackModel, error)

	// CreateUserTrackNode 创建客资流转节点。
	CreateUserTrackNode(node *model.UserTrackNodeModel) error

	// GetUserTrackNode 查询客资指定阶段的节点。
	GetUserTrackNode(trackID int64, stage int) (*model.UserTrackNodeModel, error)

	// GetUserTrackNodeList 查询客资的全部流转节点。
	GetUserTrackNodeList(trackID int64) ([]model.UserTrackNodeModel, error)

	// UpdateUserTrackNode 更新客资指定阶段的节点。
	UpdateUserTrackNode(trackID int64, stage int, updates map[string]interface{}) error

	// AppendFollowUpRecord 向任意阶段节点追加一条跟进记录。
	AppendFollowUpRecord(trackID int64, stage int, record model.FollowUpRecord) error
}

// UserTrackManagementInterfaceImp 客资管理数据访问实现。
type UserTrackManagementInterfaceImp struct{}

// ImpUserTrackManagement 客资管理 DAO 实例。
var ImpUserTrackManagement UserTrackManagementInterface = &UserTrackManagementInterfaceImp{}

var _ UserTrackManagementInterface = (*UserTrackManagementInterfaceImp)(nil)
