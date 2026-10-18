-- 业务上层按 uid、is_used 查询未使用授权，存在时阻断新增。
CREATE TABLE `course_package_purchase_whitelist` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '授权记录 ID',
  `uid` BIGINT NOT NULL COMMENT '获得优惠的用户 ID，必须大于 0',
  `min_course_cnt` INT NOT NULL COMMENT '本次授权允许的最低购买次数，必须大于 0',
  `is_used` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否已使用：0未使用，1已使用',
  `order_id` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '使用本次优惠的支付订单 ID',
  `used_ts` BIGINT NOT NULL DEFAULT 0 COMMENT '使用时间，Unix 秒',
  `created_by` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '创建人（顾问/运营）',
  `remark` TEXT COMMENT '授权原因或备注',
  `created_ts` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间，Unix 秒',
  `updated_ts` BIGINT NOT NULL DEFAULT 0 COMMENT '更新时间，Unix 秒',
  PRIMARY KEY (`id`),
  KEY `idx_uid_is_used` (`uid`, `is_used`),
  KEY `idx_uid_id` (`uid`, `id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='课包购买最低次数白名单授权记录';
