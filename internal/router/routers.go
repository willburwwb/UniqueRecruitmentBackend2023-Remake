package router

import (
	"UniqueRecruitmentBackend/internal/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewRouter create backend http group routers
func NewRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.AuthMiddleware)
	if gin.Mode() == gin.DebugMode {
		r.Use(cors.Default())
	} else if gin.Mode() == gin.ReleaseMode {
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"https://join.hustunique.com", "https://hr.hustunique.com"}
		config.AllowMethods = []string{"GET", "POST", "DELETE", "UPDATE", "PUT", "OPTION"}
		r.Use(cors.New(config))
	}
	ping := r.Group("/ping")
	{
		ping.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "this is uniquestudio hr system",
			})
		})
	}
	//memberRouter := r.Group("/members")
	//{
	//	memberRouter.GET("/me")
	//	memberRouter.GET("/group")
	//	memberRouter.PUT("/me")
	//	memberRouter.PUT("/admin")
	//}
	//
	//candidateRouter := r.Group("/candidates")
	//{
	//	candidateRouter.POST("/")
	//	candidateRouter.GET("/me")
	//	candidateRouter.PUT("/me")
	//	candidateRouter.PUT("/me/password")
	//}
	//
	//applicationRouter := r.Group("/applications")
	//{
	//	applicationRouter.POST("/")
	//	applicationRouter.GET("/:aid")
	//	applicationRouter.PUT("/:aid")
	//	applicationRouter.DELETE("/:aid")
	//	applicationRouter.PUT("/:aid/abandoned")
	//	applicationRouter.GET("/:aid/slots/:type")
	//	applicationRouter.PUT("/:aid/slots/:type")
	//	applicationRouter.GET("/:aid/resume")
	//	applicationRouter.GET("/recruitment/:rid")
	//}
	return r
}
