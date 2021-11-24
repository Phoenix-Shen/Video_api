package common

import (
	"fmt"
	"video_api/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	userName := viper.GetString("datasource.userName")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		userName, password, host, port, database, charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database:" + err.Error())
	}
	db.AutoMigrate(&model.VideoEffect{})
	return db
}

func GetDB() *gorm.DB {
	return DB
}
