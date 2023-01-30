package postgre

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"v1/src/errors"
)

func main() {
	var dsn = "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	errors.CheckErr(err)
}
