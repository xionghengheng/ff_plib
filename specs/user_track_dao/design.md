# Technical Design

## Architecture

在现有 `db/dao` 包中增加客资管理专用 interface 和实现，不改变全局数据库初始化方式。

- `interface_user_track_management.go`：定义 `UserTrackManagementInterface`、空实现类型及全局接口实例。
- `dao_user_track_management.go`：基于 GORM v1 实现数据访问。
- `model/user_track_management.go`：继续承载表模型和跟进 JSON 类型。

## Interface

```go
type UserTrackManagementInterface interface {
	CreateUserTrack(track *model.UserTrackModel) error
	GetUserTrack(trackID int64) (*model.UserTrackModel, error)
	GetUserTrackByWechatNo(wechatNo string) (*model.UserTrackModel, error)
	UpdateUserTrack(trackID int64, updates map[string]interface{}) error
	GetUnfinishedUserTrackList(stage int, lastCreatedTs int64, lastTrackID int64, limit int) ([]model.UserTrackModel, error)
	GetAllUserTrackList(lastCreatedTs int64, lastTrackID int64, limit int) ([]model.UserTrackModel, error)

	CreateUserTrackNode(node *model.UserTrackNodeModel) error
	GetUserTrackNode(trackID int64, stage int) (*model.UserTrackNodeModel, error)
	GetUserTrackNodeList(trackID int64) ([]model.UserTrackNodeModel, error)
	UpdateUserTrackNode(trackID int64, stage int, updates map[string]interface{}) error
	AppendFollowUpRecord(trackID int64, stage int, record model.FollowUpRecord) error
}
```

## Table Mapping

| Model | Table | Primary lookup |
| --- | --- | --- |
| `UserTrackModel` | `user_track` | `track_id` |
| `UserTrackNodeModel` | `user_track_node` | `track_id + stage` |

DAO 通过 `Table(...)` 显式指定表名，与现有代码风格一致。

## Follow-up Append Flow

`AppendFollowUpRecord` 使用数据库事务：

1. 根据 `track_id + stage` 使用 `SELECT ... FOR UPDATE` 锁定节点行。
2. 使用 `GetFollowUpRecords` 解析当前 `data`。
3. 追加新记录并序列化为 JSON 数组。
4. 更新 `data` 和 `node_updated_ts`，其中 `node_updated_ts` 使用 `record.FollowTs`。
5. 任一步失败则回滚。

不对 `stage` 值做 02/06 限制；只要目标节点存在即可追加。

## Query and Update Behavior

- `GetUserTrackNodeList` 不分页，按 `stage ASC, node_create_ts ASC` 返回指定客资的全部节点。
- `GetUnfinishedUserTrackList` 仅查询 `stage < 7` 的客资，按 `created_ts DESC, track_id DESC` 排序，并使用 `lastCreatedTs/lastTrackID/limit` 复合游标分页；`lastCreatedTs<=0` 取第一页。
- `GetAllUserTrackList` 不过滤 stage，按 `created_ts DESC, track_id DESC` 排序，并使用相同的复合游标分页。
- 未完成列表中 `stage<=0` 表示全部阶段，`stage>0` 表示精确过滤。
- 单条查询无数据时直接返回 `gorm.ErrRecordNotFound`。
- Update 方法使用 `Model(...).Where(...).Updates(...)`，不自动改写调用方传入的字段。
- Create 方法使用 `Create`，不使用 `Save` 的 upsert 语义。

## Verification

- 保留跟进 JSON 解析的单元测试。
- 运行 `gofmt`、`go test ./db/...` 和 `go test ./...` 验证编译及回归。
- 当前项目未引入 SQL mock 依赖，本次不为 DAO 额外增加第三方测试库。
