# 课包退费历史记录设计（已确认）

## 实现范围

沿用现有 Go、GORM v1 和 MySQL 结构，新增以下文件，不修改已有课包和白名单接口：

- `db/model/course_package_refund_record.go`
- `db/dao/interface_course_package_refund_record.go`
- `db/dao/dao_course_package_refund_record.go`
- `db/schema/course_package_refund_record.sql`

## 模型与表

模型命名 `CoursePackageRefundRecordModel`，表名 `course_package_refund_record`。字段使用蛇形 JSON 名称，与数据库列名一致。

| Go 字段 | 类型 | 数据库列 / 类型 | 含义 |
| --- | --- | --- | --- |
| ID | int64 | id BIGINT | 自增主键 |
| Uid | int64 | uid BIGINT | 用户 ID |
| PackageID | string | package_id VARCHAR(64) | 课包 ID |
| RefundLessonCnt | int | refund_lesson_cnt INT | 本次退费或清零课时数 |
| RefundAmount | int64 | refund_amount BIGINT | 本次退款金额，整数元；0 表示未登记退款金额 |
| RefundTs | int64 | refund_ts BIGINT | 退费或归零操作时间，Unix 秒 |
| CreatedBy | string | created_by VARCHAR(128) | 操作人，由业务层提供 |
| Remark | string | remark TEXT | 操作原因或备注 |
| CreatedTs | int64 | created_ts BIGINT | 记录创建时间，Unix 秒，由业务层提供 |

金额使用 int64，与用户提供的请求结构一致；不改变现有课包模型的金额类型。记录为操作当时的快照，查询不依赖课包当前状态。不设置外键或用户、课包唯一约束，允许同一用户、课包存在多条历史。

表使用 InnoDB、utf8mb4；主键自增，其他数值列 NOT NULL DEFAULT 0，VARCHAR 列 NOT NULL DEFAULT ''，备注 TEXT 可空。

索引：

- `PRIMARY KEY (id)`：全量游标分页。
- `idx_uid_id (uid, id)`：按用户游标分页。
- `idx_package_id_id (package_id, id)`：按课包游标分页。

## DAO 接口

```go
type CoursePackageRefundRecordInterface interface {
    AddCoursePackageRefundRecord(record *model.CoursePackageRefundRecordModel) error
    GetCoursePackageRefundRecordList(lastID int64, limit int) ([]model.CoursePackageRefundRecordModel, error)
    GetCoursePackageRefundRecordListByUid(uid, lastID int64, limit int) ([]model.CoursePackageRefundRecordModel, error)
    GetCoursePackageRefundRecordListByPackageId(packageID string, lastID int64, limit int) ([]model.CoursePackageRefundRecordModel, error)
}
```

使用 `CoursePackageRefundRecordInterfaceImp` 实现，导出 `ImpCoursePackageRefundRecord` 实例，并添加接口实现的编译期断言。

新增使用 `db.Get().Table(...).Create(record)`，自增 ID 回填到 record。查询使用参数化 Raw SQL 和 Scan，空列表不视为错误，数据库错误原样返回。lastID <= 0 为首页，后续增加 `id < ?`，统一 `ORDER BY id DESC LIMIT ?`。过滤值、游标和条数均使用占位符。

DAO 不校验业务参数，不自动填写时间，不提供修改或删除历史的接口。调用方负责正数 limit、金额/课时有效性、鉴权、课包归属及重复请求控制。

## 业务接入边界

本次不修改用户提供的 HTTP handler，不发起退款，不直接执行建表 SQL，也不自动补历史。

归零时记录 refund_lesson_cnt 为清零次数；refund_amount 记录本次登记金额而非课包累计金额。记录中不保存管理员密码或认证头。

独立的新增 DAO 与现有 UpdateCoursePackage 各自使用 db.Get()，先后调用不具备跨调用事务保证。正式接入时须另行将课包条件更新和记录插入置于同一事务，处理并发扣课和重复请求；本次不声称已实现该原子操作。

## 验证

- gofmt 格式化新增 Go 文件，git diff --check 检查补丁。
- 执行 `go test -vet=off ./db/model ./db/dao` 检查编译；已有 db/init.go 的 vet 问题不在本次修改范围。
- 核对结构体、SQL 列名、金额类型和查询索引一致性。
- 未连接测试数据库时，不声称已经完成真实数据库分页或插入测试。
