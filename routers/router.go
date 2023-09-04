package routers

import (
	"dousheng/controller"
	"github.com/gin-gonic/gin"
)

//func index(c *gin.Context) {
//	fmt.Println("a")
//}

func InitRouter() *gin.Engine {
	router := gin.Default()

	apiRouter := router.Group("/douyin")
	apiRouter.GET("/feed", controller.Feed)
	//apiRouter.GET("/user/", controller.User)

	userRouter := apiRouter.Group("/user")
	userRouter.GET("/", controller.User)
	userRouter.POST("/register/", controller.Register)
	userRouter.POST("/login/", controller.Login)

	publishRouter := apiRouter.Group("/publish")
	publishRouter.POST("/action/", controller.Upload)
	return router
}
