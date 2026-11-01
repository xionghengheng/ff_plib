# KPI 配置实现设计

三个 model、三组 interface/DAO 文件，分别对应 kpi_member、kpi_member_month_config、kpi_month_target。保持用户提供的表、列及唯一键定义，不执行或重建数据库表。

- BIGINT 使用 int64，INT/TINYINT 使用 int，CHAR/VARCHAR 使用 string。
- 可空客资数使用 *int，DECIMAL(5,2) 百分比使用 *float64；nil 对应 NULL，JSON 不加 omitempty。百分比用于配置展示，数据库按两位小数存储；不用于精确金额运算。
- 带数字的百分比字段显式标注 GORM column，避免默认命名与 DDL 不一致。GMVTargetFen 为 int64，单位分。
- 导出 ImpKPIMember、ImpKPIMemberMonthConfig、ImpKPIMonthTarget；均具有接口实现编译期断言。
- 人员列表 memberType=0 表示全部，否则按类型筛选，按 ID 升序；人员月配置列表按月份过滤、member_id 升序。这些是配置列表，暂不引入分页接口。
- 新增直接 Create，重复键错误由数据库返回，不做隐式 upsert；更新按唯一标识执行 Updates(map)，支持 0 和 nil，不自动填写时间、操作人或人员快照。
- 单点查询未找到返回 gorm.ErrRecordNotFound，列表为空不报错；数据库错误原样返回。
- 不增加删除、历史月份扫描、关联查询等额外接口；不改既有 DAO。

验证：gofmt、补丁检查、字段映射及 JSON NULL/0 测试、全仓库编译测试。不连接真实数据库，不声称验证了真实 SQL 读写。

