package systems

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb() (*gorm.DB, error) {
	connectionString := "root:@tcp(127.0.0.1:3306)/heroes?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	return db, err
}
