//Models/SpacecraftModel.go
package Models

import (
	"github.com/jinzhu/gorm"
  )

type Spacecraft struct {
	gorm.Model
	Name     string  `gorm:"size:255"` // set field size to 255
	Class    int     `gorm:"not null"`
	Crew     uint    `gorm:"not null"`
	Image    string  `gorm:"not null"`
	Value    float64 `gorm:"not null"`
	Status   uint8   `gorm:"not null"`
	Armament string  `gorm:"not null"`
}

func (b *Spacecraft) TableName() string {
	return "spacecraft"
}