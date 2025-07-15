package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/controller"
	"github.com/liaozzzzzz/code-push-server/internal/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	// 创建Gin引擎
	r := gin.Default()

	// 添加全局中间件
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "ok",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 创建控制器实例
	// userController := controller.NewUserController()
	loginController := controller.NewLoginController()

	// API路由组
	api := r.Group("/api/v1")
	{
		// 登录相关路由
		api.POST("/login", loginController.Login)

		// 需要认证的路由
		authenticated := api.Group("")
		authenticated.Use(middleware.AuthRequired())
		{
			authenticated.POST("/logout", loginController.Logout)

			// 用户管理路由
			// users := api.Group("/users")
			// {
			// 	users.POST("", userController.Create)       // 创建用户
			// 	users.GET("", userController.List)          // 获取用户列表
			// 	users.GET("/:id", userController.GetByID)   // 获取用户详情
			// 	users.PUT("/:id", userController.Update)    // 更新用户
			// 	users.DELETE("/:id", userController.Delete) // 删除用户
			// }
		}
	}

	return r
}
