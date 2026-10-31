-- 每次退费或课时归零独立新增记录，仅记账，不触发线上退款。
CREATE TABLE `course_package_refund_record` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '自增记录 ID',
  `uid` BIGINT NOT NULL DEFAULT 0 COMMENT '用户 ID',
  `package_id` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '课包 ID',
  `before_remain_cnt` INT NOT NULL DEFAULT 0 COMMENT '操作前剩余课时',
  `after_remain_cnt` INT NOT NULL DEFAULT 0 COMMENT '操作后剩余课时',
  `refund_lesson_cnt` INT NOT NULL DEFAULT 0 COMMENT '本次退费或清零课时数',
  `refund_amount` BIGINT NOT NULL DEFAULT 0 COMMENT '本次退款金额，单位元；0表示未登记退款金额',
  `refund_ts` BIGINT NOT NULL DEFAULT 0 COMMENT '退费或归零操作时间，Unix 秒',
  `created_by` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '操作人',
  `remark` TEXT COMMENT '操作原因或备注',
  `created_ts` BIGINT NOT NULL DEFAULT 0 COMMENT '记录创建时间，Unix 秒',
  PRIMARY KEY (`id`),
  KEY `idx_uid_id` (`uid`, `id`),
  KEY `idx_package_id_id` (`package_id`, `id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='课包退费及课时归零历史记录';
