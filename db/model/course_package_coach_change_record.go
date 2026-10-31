package model

// CoursePackageCoachChangeRecordModel 课包更换教练操作记录。
// 每次操作独立新增，不覆盖历史；原教练 ID 应取自更换前的课包。
// 参数有效性、操作权限及时间字段由业务上层负责。
type CoursePackageCoachChangeRecordModel struct {
	ID            int64  `json:"id"`              // 自增记录 ID
	Uid           int64  `json:"uid"`             // 用户 ID
	PackageID     string `json:"package_id"`      // 课包 ID
	OldCoachId    int    `json:"old_coach_id"`    // 原教练 ID
	NewCoachId    int    `json:"new_coach_id"`    // 新教练 ID
	ChangeCoachTs int64  `json:"change_coach_ts"` // 更换教练时间，Unix 秒
	CreatedBy     string `json:"created_by"`      // 操作人
	Remark        string `json:"remark"`          // 操作原因或备注
	CreatedTs     int64  `json:"created_ts"`      // 记录创建时间，Unix 秒
}
