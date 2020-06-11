package user

import (
	"GoBP/config"
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Username string
	Password string
	Name string
}

var db *gorm.DB

func Init(){
	db = config.Open()
	db.AutoMigrate(&Users{})
}
func GetUser(id int) (Users){
	Init()
	var user Users

	var err = db.Close()
	if err != nil {
		panic(err)
	}
	return user

}

