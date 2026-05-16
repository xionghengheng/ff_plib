package model

// CoreOpDailyStatModel 核心操作按天统计表
// 表 core_op_daily_stat 需在 stat_date 上建 UNIQUE 索引，写入走 INSERT ... ON DUPLICATE KEY UPDATE 原子 +1
type CoreOpDailyStatModel struct {
	ID                           int64 `json:"id"`                               // 主键
	StatDate                     int   `json:"stat_date"`                        // 统计日期，格式 yyyymmdd，唯一
	UserBookCount                int   `json:"user_book_count"`                  // 用户主动约课次数
	CoachScheduleCount           int   `json:"coach_schedule_count"`             // 教练排课次数（新增可约时段）
	CoachSetUnavailableCount     int   `json:"coach_set_unavailable_count"`      // 教练设置不可用时间次数
	CoursePackageOrderCount      int   `json:"course_package_order_count"`       // 课包下单次数
	CoursePackagePaySuccessCount int   `json:"course_package_pay_success_count"` // 课包支付成功次数
	UserCancelBookCount          int   `json:"user_cancel_book_count"`           // 用户主动取消预约次数
	CoachCancelBookCount         int   `json:"coach_cancel_book_count"`          // 教练取消预约次数
	CourseCompleteWriteOffCount  int   `json:"course_complete_write_off_count"`  // 课程完成核销次数
	CreatedAt                    int64 `json:"created_at"`                       // 创建时间，Unix秒
	UpdatedAt                    int64 `json:"updated_at"`                       // 更新时间，Unix秒
}
