package api

import (
	"demo_hubu_backend/config"
	"demo_hubu_backend/controller"
	"demo_hubu_backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func User(r *gin.Engine, db *gorm.DB) {
	middleware.SetDB(db)
	r.POST("/api/login", func(c *gin.Context) {
		controller.Login(c, db)
	})

	r.POST("/api/register", func(c *gin.Context) {
		controller.Register(c, db)
	})

	r.GET("/api/logout", func(c *gin.Context) {
		controller.Logout(c, db)
	})
}

func ApprovalResource(r *gin.Engine, db *gorm.DB) {
	middleware.SetDB(db)
	api := r
	// 需要认证的路由
	protected := api.Group("/", middleware.AuthMiddleware())
	{
		// 查询和创建产品的路由
		createAndRead := protected.Group("/", middleware.RoleMiddleware(config.RoleOfApplicant))
		{
			createAndRead.GET("/api/content/getList", func(c *gin.Context) {
				controller.GetContentList(c, db)
			})
			createAndRead.POST("/api/content/create", func(c *gin.Context) {
				controller.ContentCreate(c, db)
			})
			createAndRead.GET("/api/resource/getInfo", func(c *gin.Context) {
				controller.GetComprehensiveViewList(c, db)
			})
		}

		// 查询和更新产品的路由
		readAndUpdate := protected.Group("/", middleware.RoleMiddleware(config.RoleOfApprover))
		{
			readAndUpdate.POST("/api/content/update", func(c *gin.Context) {
				controller.ContentUpdate(c, db)
			})
		}
	}
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 配置 CORS 中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // 允许的来源
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Authorization"}
	config.AllowCredentials = true // 允许发送 cookies

	r.Use(cors.New(config))
	// User Controller Api 注册
	User(r, db)

	// ApprovalResource Controller Api 注册
	ApprovalResource(r, db)

	return r
}
