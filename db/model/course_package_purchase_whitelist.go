package model

// CoursePackagePurchaseWhitelistModel 课包购买最低次数白名单授权记录。
// 一条记录只允许享受一次优惠；同一用户上一条授权使用后，才能新增下一条。
// 对应 course_package_purchase_whitelist 表，未使用授权的唯一性由数据库保证。
// 使用授权时，应在支付发放课时的同一事务中按 id、uid、is_used = 0 条件更新，
// 并检查 RowsAffected == 1，避免同一授权重复使用。
// 下单时还需将授权绑定到具体订单，避免多个待支付订单同时使用同一授权。
type CoursePackagePurchaseWhitelistModel struct {
	ID           int64  `json:"id"`             // 授权记录 ID
	Uid          int64  `json:"uid"`            // 获得优惠的用户 ID，必须大于 0
	MinCourseCnt int    `json:"min_course_cnt"` // 本次授权允许的最低购买次数，必须大于 0；不改变课包价格
	IsUsed       bool   `json:"is_used"`        // 是否已使用，支付成功并发放课时后置为 true
	OrderID      string `json:"order_id"`       // 使用本次优惠的支付订单 ID，未使用时为空
	UsedTs       int64  `json:"used_ts"`        // 使用时间，Unix 秒；未使用时为 0
	CreatedBy    string `json:"created_by"`     // 创建人（顾问/运营）
	Remark       string `json:"remark"`         // 授权原因或备注
	CreatedTs    int64  `json:"created_ts"`     // 创建时间，Unix 秒
	UpdatedTs    int64  `json:"updated_ts"`     // 更新时间，Unix 秒
}
