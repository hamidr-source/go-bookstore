package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "hamidr-source/09944947492h/simplerest?charsey=utf8&parseTime=True7loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

fun GetDB() *gorm.DB{
	return db
}