package model

type CoachClientTraineeCommentModel struct {
	LessonID      string `json:"lesson_id"`       // 单节课的唯一标识符（用户id_场地id_课程id_教练id_发起预约的时间戳）
	PackageID     string `json:"package_id"`      // 关联的课包的唯一标识符
	ScheduleBegTs int64  `json:"schedule_beg_ts"` // 单节课的安排上课时间
	Uid           int64  `json:"uid"`             // 用户id
	GymId         int    `json:"gym_id"`          // 场地id
	CoachId       int    `json:"coach_id"`        // 教练id
	CourseID      int    `json:"course_id"`       // 课程id

	//评论相关内容
	Overall              int    `json:"overall"`                // 整体
	Professional         int    `json:"professional"`           // 专业
	Environment          int    `json:"environment"`            // 环境
	Service              int    `json:"service"`                // 服务
	ContinueAttendLesson int    `json:"continue_attend_lesson"` // 是否愿意继续上课，愿意、待考虑、不愿意
	CommentContent       string `json:"comment_content"`        // 评价内容
	AnonymousComment     bool   `json:"anonymous_comment"`      // 是否匿名评价
	CommentTs            int64  `json:"comment_ts"`             // 提交评价的时间
	IsApproved           bool   `json:"is_approved"`            // 运营审核通过，才在教练端展示，默认是false，不展示
}
