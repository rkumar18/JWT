package lib

import("github.com/jinzhu/gorm")

type userData struct{
	gorm.Model
	Username string
	Name string
	Contact int64
	Email string
	Password string
}
type Postgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Db_Name  string `yaml:"dbname"`

}



type Service struct{
	Port string `yaml:"port"`
}