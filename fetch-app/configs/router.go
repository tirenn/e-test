package configs

import (
	"os"
	"tirenn/test-efishery/fetch-app/domain"
	"tirenn/test-efishery/fetch-app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()

	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/api/v1/health"),
		gin.Recovery(),
	)
	router.Use(cors.Default())

	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "https://stein.efishery.com/v1"
	}

	// Initiate Domain
	fetch := domain.InitFetchAPI(baseUrl)

	v1 := router.Group("/api")
	{
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"status": "UP"})
		})

		fetchRouter := v1.Group("/fetch")
		{
			fetchRouter.GET("", middleware.AuthorizeJWT([]string{"staff", "admin"}), fetch.Get)

			agregateRouter := fetchRouter.Group("/aggregate")
			{
				agregateRouter.GET("", middleware.AuthorizeJWT([]string{"admin", "staff"}), fetch.GetAggregate)
			}

			tokenRouter := fetchRouter.Group("/token")
			{
				tokenRouter.GET(":token", fetch.GetJWT)
			}

		}

	}

	return router
}
