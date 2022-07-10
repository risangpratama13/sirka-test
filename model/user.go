package model

type User struct {
	Userid string `gorm:"type:varchar(50)"`
	Name   string `gorm:"type:varchar(50)"`
}
