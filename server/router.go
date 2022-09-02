package server

import (
	"github.com/gin-gonic/gin"
)

// func NewRouter() *gin.Engine {
// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	router.Use(gin.Recovery())

// 	callController := new(controller.CallController)

// 	callController.InitializeCallController()

// 	rootPath := router.Group("/api")
// 	{
// 		versionPath := rootPath.Group("v1")
// 		{
// 			account := versionPath.Group("account")
// 			{
// 				account.POST(":account_id/Call/:call_id", callController.CreateCall)
// 				account.PUT(":account_id/Call/:call_id", callController.UpdateCall)
// 				account.GET(":account_id/Call/:call_id", callController.GetCall)
// 				account.DELETE(":account_id/Call/:call_id", callController.DeleteCall)
// 			}

// 			versionPath.GET("health", callController.GetHealth)
// 		}
// 	}
// 	return router
// }

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", GetHealth) //Ali SLB监控检查，Docker健康检查
	rootPath := router.Group("/api")
	{
		versionPath := rootPath.Group("v1")
		{
			versionPath.POST("CreateCall", CreateCall)
		}
	}

	return router
}

// GetHealth 健康检测
func GetHealth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Status": "0",
	})
}

//处理客户端HTTP请求
func (ctx *gin.Context) {
	//调用自研TTS
	shuke.Shake(ctx)
}

//checkHandler 异步查询接口
func checkHandler(ctx *gin.Context) {
	//调用自研TTS
	shuke.AsyncCheck(ctx)
}
