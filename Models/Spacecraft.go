//Models/Spacecraft.go
package Models

import (
	"fleet-inventory/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllSpacecraft Fetch all spacecraft data
func GetAllSpacecraft(spacecraft *[]Spacecraft) (err error) {
	if err = Config.DB.Find(spacecraft).Error; err != nil {
		return err
	}
	return nil
}

//GetSpacecraftByID ... Fetch only one spacecraft by Id
func GetSpacecraftByID(spacecraft *Spacecraft, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(spacecraft).Error; err != nil {
		return err
	}
	return nil
}

//CreateSpacecraft ... Insert New data
func CreateSpacecraft(spacecraft *Spacecraft) (err error) {
	if err = Config.DB.Create(spacecraft).Error; err != nil {
		return err
	}
	return nil
}

//UpdateSpacecraft ... Update spacecraft
func UpdateSpacecraft(spacecraft *Spacecraft, id string) (err error) {
	fmt.Println(spacecraft)
	Config.DB.Save(spacecraft)
	return nil
}

//DeleteSpacecraft ... Delete spacecraft
func DeleteSpacecraft(spacecraft *Spacecraft, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(spacecraft)
	return nil
}
