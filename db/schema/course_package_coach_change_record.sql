-- 每次更换教练独立新增记录，保留操作前后的教练 ID，不覆盖历史。
CREATE TABLE `course_package_coach_change_record` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '自增记录 ID',
  `uid` BIGINT NOT NULL DEFAULT 0 COMMENT '用户 ID',
  `package_id` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '课包 ID',
  `old_coach_id` INT NOT NULL DEFAULT 0 COMMENT '原教练 ID',
  `new_coach_id` INT NOT NULL DEFAULT 0 COMMENT '新教练 ID',
  `change_coach_ts` BIGINT NOT NULL DEFAULT 0 COMMENT '更换教练时间，Unix 秒',
  `created_by` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '操作人',
  `remark` TEXT COMMENT '操作原因或备注',
  `created_ts` BIGINT NOT NULL DEFAULT 0 COMMENT '记录创建时间，Unix 秒',
  PRIMARY KEY (`id`),
  KEY `idx_uid_id` (`uid`, `id`),
  KEY `idx_package_id_id` (`package_id`, `id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='课包更换教练操作记录';
