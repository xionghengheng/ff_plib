package model

// CoachClientTraineeNicknameModel 教练对用户自定义昵称（备注名）
// 每个教练对每个用户只能有一条记录，(coach_id, trainee_uid) 唯一
type CoachClientTraineeNicknameModel struct {
	Id         int32  `json:"id"`          // 自增主键
	CoachId    int    `json:"coach_id"`    // 教练id
	TraineeUid int64  `json:"trainee_uid"` // 用户id
	Nickname   string `json:"nickname"`    // 教练自定义的用户昵称
	CreateTs   int64  `json:"create_ts"`   // 创建时间
	UpdateTs   int64  `json:"update_ts"`   // 更新时间
}
