-- 支持仅按手机号建档，未关联用户时 uid 为 0。
-- 业务上层按手机号或有效 uid 查询未使用授权，存在时阻断新增。
CREATE TABLE `course_package_purchase_whitelist` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '授权记录 ID',
  `uid` BIGINT NOT NULL DEFAULT 0 COMMENT '获得优惠的用户 ID，0表示尚未关联用户',
  `user_phone` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '获得优惠的用户手机号',
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
  KEY `idx_uid_id` (`uid`, `id`),
  KEY `idx_user_phone_is_used` (`user_phone`, `is_used`),
  KEY `idx_user_phone_id` (`user_phone`, `id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='课包购买最低次数白名单授权记录';
