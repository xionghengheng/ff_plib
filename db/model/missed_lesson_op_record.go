package model

// MissedLessonOpRecordModel 异常课程（已旷课）补核销/补取消操作历史
// 对应表 missed_lesson_op_record，每成功执行一次补核销/补取消，落一条记录
type MissedLessonOpRecordModel struct {
	ID           int64  `json:"id"`            // 主键ID
	OpType       int    `json:"op_type"`       // 操作类型：1-补核销，2-补取消
	OperatorName string `json:"operator_name"` // 操作人（顾问姓名，否则管理员用户名）

	Uid         int64  `json:"uid"`          // 用户id
	PhoneNumber string `json:"phone_number"` // 用户手机号
	UserName    string `json:"user_name"`    // 用户昵称
	LessonId    string `json:"lesson_id"`    // 单节课唯一标识符
	PackageId   string `json:"package_id"`   // 课包唯一标识符

	GymId       int    `json:"gym_id"`        // 门店id
	GymName     string `json:"gym_name"`      // 门店名称
	CoachId     int    `json:"coach_id"`      // 教练id
	CoachName   string `json:"coach_name"`    // 教练姓名
	LessonBegTs int64  `json:"lesson_beg_ts"` // 课程开始时间

	OrigStatus int `json:"orig_status"` // 原状态（补核销/补取消前，固定为已旷课）
	NewStatus  int `json:"new_status"`  // 新状态（补核销→已完成，补取消→已取消）

	AffectUserPackage bool `json:"affect_user_package"` // 是否影响用户课包（本次是否真的增减了课时）
	AffectSettlement  bool `json:"affect_settlement"`   // 是否影响结算

	Reason   string `json:"reason"`    // 补核销/补取消原因
	Remark   string `json:"remark"`    // 备注
	OpDevice string `json:"op_device"` // 操作设备（存档用，前端不展示）
	OpTs     int64  `json:"op_ts"`     // 操作时间
}

// 补核销/补取消操作类型枚举
const (
	Enum_MissedLessonOpType_WriteOff = 1 // 补核销
	Enum_MissedLessonOpType_Cancel   = 2 // 补取消
)
