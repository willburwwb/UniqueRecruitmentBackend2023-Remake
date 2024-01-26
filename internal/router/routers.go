package router

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/controllers"
	"UniqueRecruitmentBackend/internal/middlewares"
	"UniqueRecruitmentBackend/internal/tracer"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// NewRouter create backend http group routers
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(tracer.TracingMiddleware)

	if gin.Mode() == gin.DebugMode {
		r.Use(cors.Default())
	} else if gin.Mode() == gin.ReleaseMode {
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"https://join.hustunique.com", "https://hr.hustunique.com"}
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

	r.Use(middlewares.AuthMiddleware)
	r.Use(middlewares.GlobalRoleMiddleWare)

	recruitmentRouter := r.Group("/recruitments")
	{
		// public
		recruitmentRouter.GET("/:rid", controllers.GetRecruitmentById)
		recruitmentRouter.GET("/pending", controllers.GetPendingRecruitment)

		// member role
		recruitmentRouter.GET("/", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.GetAllRecruitment)
		recruitmentRouter.PUT("/:rid/interviews/:name", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.SetRecruitmentInterviews)
		recruitmentRouter.GET("/:rid/interviews/:name", controllers.GetRecruitmentInterviews)

		// admin role
		recruitmentRouter.POST("/", middlewares.CheckAdminRoleMiddleWare, controllers.CreateRecruitment)
		recruitmentRouter.PUT("/:rid/schedule", middlewares.CheckAdminRoleMiddleWare, controllers.UpdateRecruitment)
	}

	applicationRouter := r.Group("/applications")
	{
		// public
		applicationRouter.POST("/", controllers.CreateApplication)
		applicationRouter.GET("/:aid", controllers.GetApplication)
		applicationRouter.PUT("/:aid", controllers.UpdateApplication)
		applicationRouter.DELETE("/:aid", controllers.DeleteApplication)
		applicationRouter.GET("/:aid/slots/:type", controllers.GetInterviewsSlots)
		applicationRouter.GET("/:aid/resume", controllers.GetResume)
		applicationRouter.PUT("/:aid/interview/:type", controllers.SetApplicationInterviewTime)

		// member
		applicationRouter.PUT("/:aid/abandoned", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.AbandonApplication)
		applicationRouter.GET("/recruitment/:rid", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.GetAllApplications)
		applicationRouter.PUT("/:aid/slots/:type", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.SelectInterviewSlots)
		applicationRouter.PUT("/:aid/step", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.SetApplicationStep)
	}

	commentRouter := r.Group("/comments")
	{
		// member
		commentRouter.POST("/", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.CreateComment)
		commentRouter.DELETE("/:cid", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.DeleteComment)
	}

	smsRouter := r.Group("/sms")
	{
		// member
		smsRouter.POST("/", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.SendSMS)
	}

	userRouter := r.Group("/user")
	{
		// public
		userRouter.GET("/me", controllers.GetUserDetail)
	}

	return r
}
