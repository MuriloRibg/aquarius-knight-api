package DataBase

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectToDataBase() {
	conectString := fmt.Sprintf("%s:%s@tcp(%s:3306)/aquarius?charset=utf8mb4&parseTime=True&loc=Local", "delta", senha, "aquarius-knight-db.mysql.database.azure.com")
	DB, err = gorm.Open(mysql.Open(conectString))
	if err != nil {
		panic(err.Error())
	}
}
