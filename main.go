//main.go
package main

import (
	"fleet-inventory/Config"
	"fleet-inventory/Models"
	"fleet-inventory/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Spacecraft{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
