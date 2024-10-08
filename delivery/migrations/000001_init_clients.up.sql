DROP TABLE IF EXISTS `clients`;

CREATE TABLE IF NOT EXISTS `clients` (
    `id`         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name`       VARCHAR(255) NOT NULL UNIQUE,
    `api_key`    VARCHAR(255) NOT NULL UNIQUE,
    `is_active`  BOOLEAN DEFAULT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
