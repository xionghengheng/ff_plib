# Implementation Plan

- [x] 1. 增加客资管理 DAO interface
  - 定义 `UserTrackManagementInterface` 的 9 个方法。
  - 增加 `UserTrackManagementInterfaceImp` 和全局实例 `ImpUserTrackManagement`。
  - _Requirement: 1, 2, 3_

- [x] 2. 实现客资档案和节点 DAO
  - 实现 `user_track` 的创建、单条查询和更新。
  - 实现 `user_track_node` 的创建、单条/列表查询和更新。
  - _Requirement: 1, 2_

- [x] 3. 实现不限阶段的跟进记录追加
  - 使用事务和行锁读取目标节点。
  - 解析、追加、序列化 JSON 数组并更新时间戳。
  - 确保节点不存在或 JSON 非法时事务回滚。
  - _Requirement: 3_

- [x] 4. 格式化与验证
  - 运行 `gofmt`、`go test ./db/...` 和 `go test ./...`。
  - 运行 `git diff --check` 检查格式问题。
  - 注：默认 `go test` 被现有 `db/init.go` 的 vet 告警拦截；使用 `-vet=off` 已完成全仓编译回归。
  - _Requirement: 1, 2, 3_

- [x] 5. 增加未完成客资分页与阶段过滤
  - 使用 `lastCreatedTs + lastTrackID + limit` 复合游标，按 `created_ts DESC, track_id DESC` 分页查询 `stage < 7` 的客资。
  - 支持 `stage=0` 查询全部未完成阶段和 `stage=1~6` 精确过滤。
  - _Requirement: 1_

- [x] 6. 调整客资节点列表查询
  - 不分页，按 `stage ASC, node_create_ts ASC` 查询指定客资的全部节点。
  - _Requirement: 2_

- [x] 7. 增加全阶段客资列表
  - 不过滤 stage，使用 `lastCreatedTs + lastTrackID + limit` 复合游标分页查询全部客资。
  - _Requirement: 1_
