package util

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/orm"
)

// getEnvOrDefault gets environment variable value or returns default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func Init() {
	// Get database configuration from environment variables
	dbUser := getEnvOrDefault("DB_USER", "root")
	dbPassword := getEnvOrDefault("DB_PASSWORD", "admin@123456")
	dbHost := getEnvOrDefault("DB_HOST", "127.0.0.1")
	dbPort := getEnvOrDefault("DB_PORT", "3306")
	dbName := getEnvOrDefault("DB_NAME", "ggTargetingEngine")

	// Register MySQL driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// Register the database
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	err := orm.RegisterDataBase("default", "mysql", dbConn)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	fmt.Println("Connected to MySQL database successfully!")
}
