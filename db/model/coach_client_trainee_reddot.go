package model

type CoachClientTraineeReddotModel struct {
	Id         int32 `json:"id"`
	CoachId    int   `json:"coach_id"`
	TraineeUid int64 `json:"trainee_uid"`
	VisitTs    int64 `json:"visit_ts"`
}
