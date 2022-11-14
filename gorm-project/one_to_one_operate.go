package main

import (
	"gorm-project/models/relate_tables"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	//username:password@tcp(ip:port)/dbname?charset=utf8&parseTime=True&loc=Local
	dsn := "root:xxj753896@tcp(localhost:3306)/gorm_project?charset=utf8mb3&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//自动迁移
	db.AutoMigrate(&relate_tables.UserProfile{}, &relate_tables.User{})

	userProfile := relate_tables.UserProfile{
		Pic:   "1.jpg",
		CPic:  "2.jpg",
		Phone: "150xxx",
		User: relate_tables.User{
			Id:   0,
			Name: "jamison",
			Age:  20,
			Addr: "xc",
		},
	}

	db.Create(&userProfile)

}