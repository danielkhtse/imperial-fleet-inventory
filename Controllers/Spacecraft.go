//Controllers/Spacecraft.go
package Controllers

import (
	"fleet-inventory/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetAllSpacecraft ... Get all Spacecraft
func GetAllSpacecraft(c *gin.Context) {
	var spacecraft []Models.Spacecraft
	err := Models.GetAllSpacecraft(&spacecraft)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, spacecraft)
	}
}

//GetSpacecraftByID ... Get the spacecraft by id
func GetSpacecraftByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var spacecraft Models.Spacecraft
	err := Models.GetSpacecraftByID(&spacecraft, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, spacecraft)
	}
}


//CreateSpacecraft ... Create Spacecraft
func CreateSpacecraft(c *gin.Context) {
	var spacecraft Models.Spacecraft
	c.BindJSON(&spacecraft)
	err := Models.CreateSpacecraft(&spacecraft)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		//c.JSON(http.StatusOK, spacecraft)
	}
}

//UpdateSpacecraft ... Update the spacecraft information
func UpdateSpacecraft(c *gin.Context) {
	var spacecraft Models.Spacecraft
	id := c.Params.ByClass("id")
	err := Models.GetSpacecraftByID(&spacecraft, id)
	if err != nil {
		c.JSON(http.StatusNotFound, spacecraft)
	}
	c.BindJSON(&spacecraft)
	err = Models.UpdateSpacecraft(&spacecraft, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		//c.JSON(http.StatusOK, spacecraft)
	}
}

//DeleteSpacecraft ... Delete the spacecraft
func DeleteSpacecraft(c *gin.Context) {
	var spacecraft Models.Spacecraft
	id := c.Params.ByName("id")
	err := Models.DeleteSpacecraft(&spacecraft, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		//c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
