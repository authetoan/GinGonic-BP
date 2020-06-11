package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var DB *gorm.DB


func Open() *gorm.DB  {
	var err error
	DB , err := gorm.Open("mysql", "root"+ ":"+"root"+"@/"+"gobp"+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return DB
}

func Close() error {
	return DB.Close()
}
