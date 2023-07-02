package main

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/router"
	"fmt"
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
	if err := global.Setup(); err != nil {
		panic(fmt.Errorf("global set up failed %s", err))
	}
	gin.SetMode(global.ServerConfig.RunMode)

	models.SetupTables()

	r := router.NewRouter()
	s := &http.Server{
		Addr:         global.ServerConfig.Addr,
		Handler:      r,
		ReadTimeout:  global.ServerConfig.ReadTimeout * time.Second,
		WriteTimeout: global.ServerConfig.WriteTimeout * time.Second,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
