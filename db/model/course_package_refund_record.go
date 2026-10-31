package model

// CoursePackageRefundRecordModel 课包退费或课时归零的历史记录。
// 每次操作独立新增，不覆盖历史；仅记账，不触发任何线上退款。
// 参数有效性、操作权限及时间字段由业务上层负责。
type CoursePackageRefundRecordModel struct {
	ID              int64  `json:"id"`                // 自增记录 ID
	Uid             int64  `json:"uid"`               // 用户 ID
	PackageID       string `json:"package_id"`        // 课包 ID
	RefundLessonCnt int    `json:"refund_lesson_cnt"` // 本次退费或清零课时数
	RefundAmount    int64  `json:"refund_amount"`     // 本次退款金额，单位：元；0 表示未登记退款金额
	RefundTs        int64  `json:"refund_ts"`         // 退费或归零操作时间，Unix 秒
	CreatedBy       string `json:"created_by"`        // 操作人
	Remark          string `json:"remark"`            // 操作原因或备注
	CreatedTs       int64  `json:"created_ts"`        // 记录创建时间，Unix 秒
}
