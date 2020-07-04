package lib

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)


var Postgres_config Postgres
var Services_config Service

func DBConnection() *gorm.DB{
	DSN := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", Postgres_config.User, Postgres_config.Password, Postgres_config.Db_Name)
	db, _ := gorm.Open("postgres", DSN)
	db.AutoMigrate(&userData{})
	return db
}

func dbsignup(input *userData)userData{
	var result userData
	db := DBConnection()
	db.Where("Email=?",input.Email).Find(&result)
	if result.Email == ""{
		db.Create(&input)
	}
	defer db.Close()
	return result

}

func dblogin(input *userData)userData{
	var result userData
	db := DBConnection() 
	db.Where("Email=?",input.Email).Find(&result)
	defer db.Close()
	return result

}

func dballuser() []userData{
	var users []userData
	db := DBConnection()
	db.Find(&users)
	defer db.Close()
	return users
}