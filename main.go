package main

import (
	"UniqueRecruitmentBackend/configs"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// @titile Swagger api
// @version 1.0
// @description  This is backend of recruitment system for Unique Studio.
// @BasePath /api/v1/

func main() {
	gin.SetMode(configs.Config.Server.RunMode)

	models.SetupTables()

	r := router.NewRouter()
	s := &http.Server{
		Addr:         configs.Config.Server.Addr,
		Handler:      r,
		ReadTimeout:  configs.Config.Server.ReadTimeout * time.Second,
		WriteTimeout: configs.Config.Server.WriteTimeout * time.Second,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
