package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// MySQL Config
	MySQLHost     string
	MySQLPort     string
	MySQLUser     string
	MySQLPassword string
	MySQLDatabase string

	// JWT Config
	JWTSecret     string
	JWTExpiration int // in hours

	// Server Config
	ServerPort string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("File .env not found, using environment variables")
	}

	AppConfig = &Config{
		// MySQL Config
		MySQLHost:     getEnv("MYSQL_HOST", "localhost"),
		MySQLPort:     getEnv("MYSQL_PORT", "3306"),
		MySQLUser:     getEnv("MYSQL_USER", "ambridge_user"),
		MySQLPassword: getEnv("MYSQL_PASSWORD", "1362rh83835668@&$"),
		MySQLDatabase: getEnv("MYSQL_DATABASE", "ambridge_db"),

		// JWT Config
		JWTSecret:     getEnv("JWT_SECRET", "ambridge_secret_key_change_this_in_production"),
		JWTExpiration: getEnvAsInt("JWT_EXPIRATION", 24), // default 24 hours

		// Server Config
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// Database access functions
func GetMySQLDSN() string {
	return AppConfig.MySQLUser + ":" + AppConfig.MySQLPassword + "@tcp(" +
		AppConfig.MySQLHost + ":" + AppConfig.MySQLPort + ")/" +
		AppConfig.MySQLDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func GetMySQLHost() string {
	return AppConfig.MySQLHost
}

func GetMySQLPort() string {
	return AppConfig.MySQLPort
}

func GetMySQLUser() string {
	return AppConfig.MySQLUser
}

func GetMySQLPassword() string {
	return AppConfig.MySQLPassword
}

func GetMySQLDatabase() string {
	return AppConfig.MySQLDatabase
}

// JWT access functions
func GetJWTSecret() string {
	return AppConfig.JWTSecret
}

func GetJWTExpiration() int {
	return AppConfig.JWTExpiration
}

// Server access functions
func GetServerPort() string {
	return AppConfig.ServerPort
}
