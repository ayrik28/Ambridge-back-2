package database

import (
	"log"
	"os"
	"time"

	"ambridge-backend/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	// تنظیمات را بارگذاری می‌کند
	config.LoadConfig()

	// اتصال به پایگاه داده با استفاده از تنظیمات
	dsn := config.GetMySQLDSN()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection established")
}

// AutoMigrate runs database migrations for all models
func AutoMigrate() error {
	log.Println("Running database migrations...")

	userTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100),
		surname VARCHAR(100),
		email VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		profile_image VARCHAR(255),
		company_name VARCHAR(100),
		company_email VARCHAR(255),
		company_address TEXT,
		company_phone VARCHAR(20),
		position VARCHAR(100),
		referral_source VARCHAR(100),
		role VARCHAR(20) DEFAULT 'user',
		refresh_token VARCHAR(255),
		resume_file VARCHAR(255),
		created_at DATETIME(3) NULL,
		updated_at DATETIME(3) NULL,
		deleted_at DATETIME(3) NULL,
		UNIQUE INDEX idx_users_email (email),
		INDEX idx_users_role (role),
		INDEX idx_users_deleted_at (deleted_at)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`

	// Execute the SQL statement for users table
	if err := DB.Exec(userTableSQL).Error; err != nil {
		log.Fatalf("Failed to create users table: %v", err)
		return err
	}

	projectTableSQL := `
	CREATE TABLE IF NOT EXISTS projects (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		projlink VARCHAR(255),
		title VARCHAR(255),
		type VARCHAR(100),
		cover VARCHAR(255),
		logo VARCHAR(255),
		profilename VARCHAR(255),
		profilepic VARCHAR(255),
		aboutproject TEXT,
		technologies TEXT,
		linkedin_link VARCHAR(255),
		telegram_link VARCHAR(255),
		x_link VARCHAR(255),
		youtube_link VARCHAR(255),
		github_link VARCHAR(255),
		insta_link VARCHAR(255),
		created_at DATETIME(3) NULL,
		updated_at DATETIME(3) NULL,
		deleted_at DATETIME(3) NULL,
		INDEX idx_projects_deleted_at (deleted_at)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`

	// Execute the SQL statement for projects table
	if err := DB.Exec(projectTableSQL).Error; err != nil {
		log.Fatalf("Failed to create projects table: %v", err)
		return err
	}

	crewTableSQL := `
	CREATE TABLE IF NOT EXISTS crews (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255),
		role VARCHAR(100),
		about TEXT,
		urlphoto VARCHAR(255),
		created_at DATETIME(3) NULL,
		updated_at DATETIME(3) NULL,
		deleted_at DATETIME(3) NULL,
		INDEX idx_crews_deleted_at (deleted_at)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`

	// Execute the SQL statement for crews table
	if err := DB.Exec(crewTableSQL).Error; err != nil {
		log.Fatalf("Failed to create crews table: %v", err)
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}
