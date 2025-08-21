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

-- Create projects table
CREATE TABLE IF NOT EXISTS projects (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    proj_link VARCHAR(255),
    title VARCHAR(255),
    type VARCHAR(100),
    cover VARCHAR(255),
    logo VARCHAR(255),
    profile_name VARCHAR(255),
    profile_pic VARCHAR(255),
    about_project TEXT,
    technologies TEXT,
    linkedin_link VARCHAR(255),
    telegram_link VARCHAR(255),
    x_link VARCHAR(255),
    youtube_link VARCHAR(255),
    github_link VARCHAR(255),
    insta_link VARCHAR(255),
    INDEX idx_projects_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create crews table
CREATE TABLE IF NOT EXISTS crews (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    username VARCHAR(255),
    role VARCHAR(100),
    about TEXT,
    urlphoto VARCHAR(255),
    INDEX idx_crews_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;