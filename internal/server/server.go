package server

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

// 创建服务器
func NewServer() *Server {
	s := new(Server)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(global.ServerConfig.RunMode)

	models.SetupTables()

	memberRouter := r.Group("/members")
	{
		memberRouter.GET("/me")
		memberRouter.GET("/group")
		memberRouter.PUT("/me")
		memberRouter.PUT("/admin")
	}

	candidateRouter := r.Group("/candidates")
	{
		candidateRouter.POST("/")
		candidateRouter.GET("/me")
		candidateRouter.PUT("/me")
		candidateRouter.PUT("/me/password")
	}

	applicationRouter := r.Group("/applications")
	{
		applicationRouter.POST("/")
		applicationRouter.GET("/:aid")
		applicationRouter.PUT("/:aid")
		applicationRouter.DELETE("/:aid")
		applicationRouter.PUT("/:aid/abandoned")
		applicationRouter.GET("/:aid/slots/:type")
		applicationRouter.PUT("/:aid/slots/:type")
		applicationRouter.GET("/:aid/resume")
		applicationRouter.GET("/recruitment/:rid")
	}
	//to do apis

	s.httpServer = &http.Server{
		Addr:         global.ServerConfig.Addr,
		Handler:      r,
		ReadTimeout:  time.Duration(global.ServerConfig.ReadTimeout),
		WriteTimeout: time.Duration(global.ServerConfig.WriteTimeout),
	}
	return s
}
func (s *Server) ListenAndServe() {
	s.httpServer.ListenAndServe()
}
