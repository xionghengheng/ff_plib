package model

// CoursePackageTransferRecordModel 课包及支付记录转移成功后的历史快照。
type CoursePackageTransferRecordModel struct {
	ID           int64  `json:"id"`
	SourceUID    int64  `json:"source_uid"`
	SourceNick   string `json:"source_nick"`
	SourcePhone  string `json:"source_phone"`
	TargetUID    int64  `json:"target_uid"`
	TargetNick   string `json:"target_nick"`
	TargetPhone  string `json:"target_phone"`
	OldPackageID string `json:"old_package_id"`
	NewPackageID string `json:"new_package_id"`
	PackageName  string `json:"package_name"`
	OrderID      string `json:"order_id"`
	TotalCnt     int    `json:"total_cnt"`
	RemainCnt    int    `json:"remain_cnt"`
	CoachID      int    `json:"coach_id"`
	CoachName    string `json:"coach_name"`
	TransferTs   int64  `json:"transfer_ts"`
	CreatedBy    string `json:"created_by"`
}
