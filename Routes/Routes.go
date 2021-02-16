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

		grpsc.GET("get/:id", Controllers.GetSpacecraftById)
		
		grpsc.POST("create", Controllers.CreateSpacecraft)

		grpsc.PUT("update/:id", Controllers.UpdateSpacecraft)

		grpsc.DELETE("delete/:id", Controllers.DeleteSpacecraft)
	}
	return r
}
