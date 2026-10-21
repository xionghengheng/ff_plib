package model

// CoursePackagePurchaseWhitelistModel 课包购买最低次数白名单授权记录。
// 一条记录只允许享受一次优惠；同一用户上一条授权使用后，才能新增下一条。
// 对应 course_package_purchase_whitelist 表，新增前由业务上层查询未使用授权并决定是否阻断。
// 支持先按手机号建档，尚未关联用户时 Uid 为 0，后续由业务上层更新关联。
// 支付成功后，由业务上层更新使用状态、订单 ID 及使用时间。
type CoursePackagePurchaseWhitelistModel struct {
	ID           int64  `json:"id"`             // 授权记录 ID
	Uid          int64  `json:"uid"`            // 获得优惠的用户 ID，0 表示尚未关联用户
	UserPhone    string `json:"user_phone"`     // 获得优惠的用户手机号
	MinCourseCnt int    `json:"min_course_cnt"` // 本次授权允许的最低购买次数，必须大于 0；不改变课包价格
	IsUsed       bool   `json:"is_used"`        // 是否已使用，支付成功并发放课时后置为 true
	OrderID      string `json:"order_id"`       // 使用本次优惠的支付订单 ID，未使用时为空
	PackageID    string `json:"package_id"`     // 关联课包 ID
	UsedTs       int64  `json:"used_ts"`        // 使用时间，Unix 秒；未使用时为 0
	CreatedBy    string `json:"created_by"`     // 创建人（顾问/运营）
	Remark       string `json:"remark"`         // 授权原因或备注
	CreatedTs    int64  `json:"created_ts"`     // 创建时间，Unix 秒
	UpdatedTs    int64  `json:"updated_ts"`     // 更新时间，Unix 秒
}
