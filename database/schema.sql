-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    profile_image VARCHAR(255),
    company_name VARCHAR(255),
    company_email VARCHAR(255),
    company_address TEXT,
    company_phone VARCHAR(20),
    position VARCHAR(255),
    referral_source VARCHAR(255),
    role VARCHAR(10) DEFAULT 'user',
    refresh_token VARCHAR(255),
    resume_file VARCHAR(255),
    UNIQUE INDEX idx_users_email (email),
    INDEX idx_users_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;