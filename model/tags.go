package model

type Tags struct {
	Id   int    `gorm:"type:init;primary_key"`
	Name string `gorm:"type:varchar(255)"`
}
