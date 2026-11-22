-- 与 user_info、gym、pass_card_gym_info 位于同一数据库，相关表必须使用 InnoDB。
CREATE TABLE IF NOT EXISTS gym_client_activation (
    id BIGINT NOT NULL AUTO_INCREMENT,
    request_id VARCHAR(64) CHARACTER SET ascii COLLATE ascii_bin NOT NULL,
    user_phone VARCHAR(20) NOT NULL,
    uid BIGINT NOT NULL,
    user_nick VARCHAR(255) NOT NULL DEFAULT '',
    gym_id INT NOT NULL,
    gym_name VARCHAR(255) NOT NULL DEFAULT '',
    previous_gym_id INT NOT NULL DEFAULT 0,
    was_coach TINYINT(1) NOT NULL DEFAULT 0,
    previous_coach_id INT NOT NULL DEFAULT 0,
    created_by VARCHAR(255) NOT NULL,
    created_ts BIGINT NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY uk_request_id (request_id),
    KEY idx_phone_id (user_phone, id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
