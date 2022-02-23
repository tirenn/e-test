package configs

import (
	"tirenn/test-efishery/auth-app/domain"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	router := gin.New()

	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/api/v1/health"),
		gin.Recovery(),
	)
	router.Use(cors.Default())

	// Initiate Domain
	auth := domain.InitAuthAPI(db)

	v1 := router.Group("/api")
	{
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"status": "UP"})
		})

		authRouter := v1.Group("/auth")
		{
			authRouter.POST("login", auth.Login)
			authRouter.POST("signup", auth.Create)
			authRouter.GET("token/:token", auth.GetJWT)

		}
	}

	return router
}
