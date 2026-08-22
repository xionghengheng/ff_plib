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
