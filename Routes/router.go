package Routes

import (
	"testapi/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/data-list", Controllers.RequestData)

	return router

}
