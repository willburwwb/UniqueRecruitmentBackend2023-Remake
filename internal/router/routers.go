package router

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/controllers"
	"UniqueRecruitmentBackend/internal/tracer"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewRouter create backend http group routers
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(tracer.TracingMiddleware)

	//TODO(wwb)
	//Add access control middleware here
	//r.Use(middlewares.MemberMiddleware)

	if gin.Mode() == gin.DebugMode {
		r.Use(cors.Default())
	} else if gin.Mode() == gin.ReleaseMode {
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"https://join.hustunique.com", "https://hr.hustunique.com"}
		config.AllowMethods = []string{"GET", "POST", "DELETE", "UPDATE", "PUT", "OPTION"}
		r.Use(cors.New(config))
	}
	r.Use(sessions.Sessions("SSO_SESSION", global.SessStore))
	ping := r.Group("/ping")
	{
		ping.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "this is uniquestudio hr system",
			})
		})
	}
	recruitmentRouter := r.Group("/recruitments")
	{
		recruitmentRouter.GET("/:rid", controllers.GetRecruitmentById)
		recruitmentRouter.GET("/", controllers.GetAllRecruitment)
		recruitmentRouter.POST("/", controllers.CreateRecruitment)
		recruitmentRouter.PUT("/:rid/schedule", controllers.UpdateRecruitment)
		recruitmentRouter.PUT("/:rid/interviews/:name", controllers.SetRecruitmentInterviews)
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
	applicationRouter := r.Group("/applications")
	{
		applicationRouter.POST("/", controllers.CreateApplication)
		applicationRouter.GET("/:aid", controllers.GetApplicationById)
		applicationRouter.PUT("/:aid", controllers.UpdateApplicationById)
		applicationRouter.DELETE("/:aid", controllers.DeleteApplicationById)
		applicationRouter.PUT("/:aid/abandoned", controllers.AbandonApplicationById)
		applicationRouter.GET("/:aid/slots/:type", controllers.GetInterviewsSlots)
		applicationRouter.PUT("/:aid/slots/:type", controllers.SelectInterviewSlots)
		applicationRouter.GET("/:aid/resume", controllers.GetResumeById)
		applicationRouter.GET("/recruitment/:rid", controllers.GetApplicationByRecruitmentId)
		applicationRouter.PUT("/:aid/step", controllers.SetApplicationStepById)
		applicationRouter.PUT("/:aid/interview/:type", controllers.SetApplicationInterviewTimeById)
		applicationRouter.PUT("/interview/:type", controllers.SetApplicationInterviewTime)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.DELETE("/:cid", controllers.DeleteComment)
	}
	return r
}
