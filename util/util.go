package util

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// DB connection constants
const (
	DBUser     = "root"
	DBPassword = "admin@123456"
	DBHost     = "127.0.0.1"
	DBPort     = "3306"
	DBName     = "ggTargetingEngine"
)

func Init() {

	// Register MySQL driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// Register the database
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		DBUser, DBPassword, DBHost, DBPort, DBName)

	err := orm.RegisterDataBase("default", "mysql", dbConn)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	fmt.Println("Connected to MySQL database successfully!")
}
