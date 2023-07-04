package db

import (
	"gorm.io/gorm"
	"main/server/model"
)

func Execute(db *gorm.DB) {
	err := db.AutoMigrate(&model.DbVersion{})
	if err != nil {
		panic(err)
	}
}
