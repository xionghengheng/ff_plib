# 课包更换教练操作记录设计

沿用已确认的课包退费历史结构与 DAO 实现方式，表名为 `course_package_coach_change_record`，模型为 `CoursePackageCoachChangeRecordModel`。

- ID：int64 / BIGINT，自增主键。
- Uid：int64 / BIGINT，用户 ID。
- PackageID：string / VARCHAR(64)，课包 ID。
- OldCoachId、NewCoachId：int / INT，原教练及新教练 ID。
- ChangeCoachTs、CreatedTs：int64 / BIGINT，操作时间、记录创建时间，Unix 秒，由调用方提供。
- CreatedBy：string / VARCHAR(128)，操作人，不保存密码或认证头。
- Remark：string / TEXT，操作原因或备注。

JSON 名称与蛇形数据库列名一致。使用 InnoDB、utf8mb4，数值非空默认 0，VARCHAR 非空默认空串，TEXT 可空。索引为主键 `(id)`、`idx_change_coach_ts_id (change_coach_ts, id)`，不增加 UID 或课包唯一约束。

导出 `ImpCoursePackageCoachChangeRecord`，仅提供 AddCoursePackageCoachChangeRecord 和 GetCoursePackageCoachChangeRecordList(lastChangeCoachTs, lastID int64, limit int)。新增回填 ID；查询用 Raw/Scan，数据库错误原样返回，空结果不报错。首页传 (0, 0)，后续传上一页最后一条的 ChangeCoachTs、ID；lastID > 0 时使用 `change_coach_ts < ? OR (change_coach_ts = ? AND id < ?)`，统一 `ORDER BY change_coach_ts DESC, id DESC LIMIT ?`。同秒记录用 ID 排序避免翻页遗漏。

接入 handler 时应先读取课包、确认归属并取得原教练 ID；不能将请求中的新教练当成原教练。课包修改与记录插入应在同一事务中处理，并考虑并发修改和重复请求。此次独立 DAO 使用 db.Get()，不具备与现有课包更新方法跨调用的事务保证，不自动接入 handler。

验证采用 gofmt、补丁空白检查及 `go test -vet=off ./...` 编译检查；不连接数据库。
