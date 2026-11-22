-- 门店端成功开通历史记录。
CREATE TABLE IF NOT EXISTS gym_client_activation_history (
    id BIGINT NOT NULL AUTO_INCREMENT,
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
    KEY idx_phone_id (user_phone, id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
