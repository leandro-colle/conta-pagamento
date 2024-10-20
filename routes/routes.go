package routes

import (
	"conta-pagamento/controllers"
	"conta-pagamento/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequest() {
	routerMap := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	routerMap.GET("/account", controllers.GetAll)
	routerMap.GET("/account/:id", controllers.GetById)
	routerMap.POST("/account", controllers.Create)
	routerMap.DELETE("/account/:id", controllers.Delete)
	routerMap.PATCH("/account/:id", controllers.Update)
	routerMap.GET("/account/agencynumber/:agencynumber", controllers.GetByAgencyNumber)

	routerMap.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routerMap.Run(":5000")
}
