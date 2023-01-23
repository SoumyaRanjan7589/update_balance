package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// declar global variable
var Database *gorm.DB
var userIDSN = "root:Soumya@123@tcp(localhost:3306)/sys"
var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(userIDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("connection faild")
	}
	Database.AutoMigrate(&Employee{})

}
