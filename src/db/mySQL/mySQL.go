package mySQL

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const username = "root"
const password = "12345678"

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	_, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
