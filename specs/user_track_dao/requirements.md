# Requirements Document

## Introduction

为客资基础档案 `user_track` 和客资节点 `user_track_node` 增加与现有 `db/dao` 一致的 DAO interface 及 GORM 实现，供业务层创建、查询和更新客资转化记录。

## Requirements

### Requirement 1 - 客资基础档案访问

**User Story:** 作为业务层，我需要通过稳定的 interface 读写客资基础档案。

#### Acceptance Criteria

1. When 业务层提交一个客资模型时，the DAO shall 向 `user_track` 表创建记录。
2. When 业务层按 `track_id` 查询时，the DAO shall 返回对应客资档案或 GORM 错误。
3. When 业务层按微信号查询时，the DAO shall 返回对应客资档案或 GORM 错误。
4. When 业务层按手机号查询时，the DAO shall 返回对应客资档案或 GORM 错误。
5. When 业务层传入 `track_id` 和更新字段时，the DAO shall 更新对应档案。
5. When 业务层分页查询未完成客资时，the DAO shall 使用 `lastCreatedTs + lastTrackID + limit` 按创建时间倒序游标分页返回数据。
6. While `stage` 为 1~6，when 业务层查询未完成客资时，the DAO shall 仅返回指定阶段数据。
7. While `stage` 为 0，when 业务层查询未完成客资时，the DAO shall 返回全部 1~6 阶段数据。
8. When 业务层分页查询全部客资时，the DAO shall 不添加 stage 过滤条件，并使用 `lastCreatedTs + lastTrackID + limit` 复合游标返回所有阶段数据。

### Requirement 2 - 客资节点访问

**User Story:** 作为业务层，我需要记录和查询客资流转节点。

#### Acceptance Criteria

1. When 业务层提交一个节点模型时，the DAO shall 向 `user_track_node` 表创建记录。
2. When 业务层按 `track_id` 查询节点列表时，the DAO shall 按阶段和节点创建时间升序返回全部节点。
3. When 业务层按 `track_id` 和 `stage` 查询时，the DAO shall 返回对应节点或 GORM 错误。
4. When 业务层传入 `track_id`、`stage` 和更新字段时，the DAO shall 更新对应节点。
5. When 业务层组装节点上的课包信息时，the node model shall 提供体验课包和付费课包两个独立结构体字段，并且 GORM shall 忽略这两个非数据库字段。

### Requirement 3 - 跟进记录 JSON

**User Story:** 作为业务层，我需要便利地追加和读取任意阶段的跟进记录。

#### Acceptance Criteria

1. When 业务层向任意阶段追加一条跟进记录时，the DAO shall 将记录作为 JSON 数组元素保存到节点 `data` 字段。
2. When 旧 `data` 为空字符串时，the DAO shall 将其视为空列表后追加。
3. When 旧 `data` 不是合法 JSON 数组时，the DAO shall 返回错误且不覆盖原数据。
4. While 多个请求同时向同一节点追加跟进记录，when DAO 执行追加时，the DAO shall 避免因读取-修改-写回竞争丢失已提交的记录。

## Proposed Scope

- 新增独立的 `interface_user_track_management.go` 和 `dao_user_track_management.go`。
- 沿用项目的 GORM v1 和全局 `db.Get()` 连接方式。
- 本次不添加建表 SQL、HTTP API 或 service 层。
