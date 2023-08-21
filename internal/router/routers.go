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

	//TODO(wwb)
	//Add access control middleware here
	//r.Use(middlewares.MemberMiddleware)

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

	r.Use(middlewares.LocalAuthMiddleware)
	r.Use(middlewares.GlobalRoleMiddleWare)
	recruitmentRouter := r.Group("/recruitments")
	{
		// public
		recruitmentRouter.GET("/:rid", controllers.GetRecruitmentById)
		recruitmentRouter.GET("/pending", controllers.GetPendingRecruitment)

		// member role
		recruitmentRouter.GET("/", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.GetAllRecruitment)

		recruitmentRouter.PUT("/:rid/interviews/:name", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.SetRecruitmentInterviews)

		// admin role
		recruitmentRouter.POST("/", middlewares.CheckAdminRoleMiddleWare, controllers.CreateRecruitment)
		recruitmentRouter.PUT("/:rid/schedule", middlewares.CheckAdminRoleMiddleWare, controllers.UpdateRecruitment)
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
		// public
		applicationRouter.POST("/", controllers.CreateApplication)
		applicationRouter.GET("/:aid", controllers.GetApplicationById)
		applicationRouter.PUT("/:aid", controllers.UpdateApplicationById)
		applicationRouter.DELETE("/:aid", controllers.DeleteApplicationById)
		applicationRouter.GET("/:aid/slots/:type", controllers.GetInterviewsSlots)
		applicationRouter.GET("/:aid/resume", controllers.GetResumeById)
		applicationRouter.PUT("/:aid/interview/:type", controllers.SetApplicationInterviewTimeById)

		// member
		applicationRouter.PUT("/:aid/abandoned", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.AbandonApplicationById)
		applicationRouter.GET("/recruitment/:rid", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.GetApplicationByRecruitmentId)
		applicationRouter.PUT("/:aid/slots/:type", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.SelectInterviewSlots)
		applicationRouter.PUT("/:aid/step", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.SetApplicationStepById)

		//applicationRouter.PUT("/interview/:type", controllers.SetApplicationInterviewTime)
	}

	// interviewRouter := r.Group("/interviews")
	// {
	// 	interviewRouter.GET("/recruitment/:rid/interviews/:name", controllers.SetRecruitmentInterviews)
	// }
	commentRouter := r.Group("/comments")
	{
		commentRouter.POST("/", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.CreateComment)
		commentRouter.DELETE("/:cid", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.DeleteComment)
	}

	smsRouter := r.Group("/sms")
	{
		smsRouter.POST("/", middlewares.CheckMemberRoleOrAdminMiddleWare, controllers.SendSMS)
	}
	return r
}
