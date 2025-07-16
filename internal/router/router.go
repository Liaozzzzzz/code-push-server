package router

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/liaozzzzzz/code-push-server/internal/config"
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

	// 在调试模式下添加pprof路由
	if config.C.General.Debug {
		debugGroup := r.Group("/debug/pprof")
		{
			debugGroup.GET("/", gin.WrapF(pprof.Index))
			debugGroup.GET("/cmdline", gin.WrapF(pprof.Cmdline))
			debugGroup.GET("/profile", gin.WrapF(pprof.Profile))
			debugGroup.POST("/symbol", gin.WrapF(pprof.Symbol))
			debugGroup.GET("/symbol", gin.WrapF(pprof.Symbol))
			debugGroup.GET("/trace", gin.WrapF(pprof.Trace))
			debugGroup.GET("/allocs", gin.WrapF(pprof.Handler("allocs").ServeHTTP))
			debugGroup.GET("/block", gin.WrapF(pprof.Handler("block").ServeHTTP))
			debugGroup.GET("/goroutine", gin.WrapF(pprof.Handler("goroutine").ServeHTTP))
			debugGroup.GET("/heap", gin.WrapF(pprof.Handler("heap").ServeHTTP))
			debugGroup.GET("/mutex", gin.WrapF(pprof.Handler("mutex").ServeHTTP))
			debugGroup.GET("/threadcreate", gin.WrapF(pprof.Handler("threadcreate").ServeHTTP))
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 创建控制器实例
	// userController := controller.NewUserController()
	loginController := controller.NewLoginController()
	deptController := controller.NewDeptController()

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

			// 部门管理
			dept := authenticated.Group("/dept")
			{

				dept.GET("/tree", deptController.SelectDeptTree)
				dept.POST("/create", deptController.Create)
				dept.PUT("/update", deptController.Update)
				dept.DELETE("/delete", deptController.Delete)
			}

			// 用户管理路由
			// users := authenticated.Group("/users")
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
