package Routes

import (
	"fleet-inventory/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	//group spacecraft
	grpsc := r.Group("/spacecraft")
	{
		grpsc.GET("list", Controllers.GetAllSpacecraft)

		grpsc.POST("create", Controllers.CreateSpacecraft)

		grpsc.GET("get/:name", Controllers.GetSpacecraftByName)

		grpsc.GET("get:class", Controllers.GetSpacecraftByClass)

		grpsc.GET("get/:status", Controllers.GetSpacecraftByStatus)

		grpsc.PUT("update/:id", Controllers.updateSpacecraft)

		grpsc.DELETE("delete/:id", Controllers.DeleteSpacecraft)
	}
	return r
}
