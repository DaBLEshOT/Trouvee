package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Treasure struct {
	gorm.Model
	Lat         float64 `json:"lat" form:"lat"`
	Lng         float64 `json:"lng" form:"lng"`
	Found       bool    `json:"found,default=false" form:"found,default=false"`
	Name        string  `json:"name,default=nil" form:"name,default=nil"`
	Description string  `json:"description,default=nil" form:"description,default=nil"`
}

func createDatabase() {
	db := openDb()
	db.AutoMigrate(&Treasure{})

	db.Create(&Treasure{
		Lat:         46.561916,
		Lng:         15.63886,
		Name:        "Test 1",
		Description: "",
	})
}

func openDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
