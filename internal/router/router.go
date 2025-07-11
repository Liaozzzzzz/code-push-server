package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/controller"
	"github.com/liaozzzzzz/code-push-server/internal/middleware"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	// 创建Gin引擎
	r := gin.Default()

	// 添加全局中间件
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	// 创建控制器实例
	userController := controller.NewUserController()
	appController := controller.NewAppController()

	// API路由组
	api := r.Group("/api")
	{
		// 认证相关路由（无需认证）
		auth := api.Group("/auth")
		{
			auth.POST("/login", userController.Login)
		}

		// 用户管理路由
		users := api.Group("/users")
		{
			users.POST("", userController.Create)                 // 创建用户
			users.GET("", userController.List)                    // 获取用户列表
			users.GET("/:id", userController.GetByID)             // 获取用户详情
			users.PUT("/:id", userController.Update)              // 更新用户
			users.DELETE("/:id", userController.Delete)           // 删除用户
			users.PUT("/:id/status", userController.UpdateStatus) // 更新用户状态
		}

		// 应用管理路由
		apps := api.Group("/apps")
		{
			apps.GET("", appController.List)                            // 获取应用列表
			apps.GET("/:id", appController.GetByID)                     // 获取应用详情
			apps.GET("/bundle/:bundle_id", appController.GetByBundleID) // 根据Bundle ID获取应用
		}

		// 需要认证的路由
		authenticated := api.Group("")
		authenticated.Use(middleware.AuthRequired())
		{
			// 应用管理（需要认证）
			authenticated.POST("/apps", appController.Create)       // 创建应用
			authenticated.PUT("/apps/:id", appController.Update)    // 更新应用
			authenticated.DELETE("/apps/:id", appController.Delete) // 删除应用

			// 用户个人应用
			my := authenticated.Group("/my")
			{
				my.GET("/apps", appController.ListByUser) // 获取当前用户的应用列表
			}
		}
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "服务运行正常",
		})
	})

	return r
}
