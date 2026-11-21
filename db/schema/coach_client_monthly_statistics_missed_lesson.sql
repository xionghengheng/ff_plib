-- 先执行本迁移，再部署使用 MissedLessonCnt 的模型和扫描服务。
-- 历史月份默认填 0；月度扫描会刷新当前月份的实际旷课数。
ALTER TABLE `coach_client_monthly_statistics`
    ADD COLUMN `missed_lesson_cnt` INT UNSIGNED NOT NULL DEFAULT 0
    COMMENT '当月旷课数（付费课、体验课、赠课，按预约上课时间统计）'
    AFTER `lesson_cnt`;
