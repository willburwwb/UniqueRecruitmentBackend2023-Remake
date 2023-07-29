package main

import (
	"UniqueRecruitmentBackend/configs"
	"UniqueRecruitmentBackend/internal/models"
	"UniqueRecruitmentBackend/internal/router"
	"UniqueRecruitmentBackend/internal/tracer"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @titile Swagger api
// @version 1.0
// @description  This is backend of recruitment system for Unique Studio.
// @BasePath /api/v1/

func main() {
	gin.SetMode(configs.Config.Server.RunMode)

	shutdown, err := tracer.SetupTracing(
		configs.Config.Apm.Name,
		configs.Config.Server.RunMode,
		configs.Config.Apm.ReportBackend,
	)
	if err != nil {
		zapx.Warn("setup tracing report backend failed", zap.Error(err))
	}

	models.SetupTables()

	r := router.NewRouter()
	s := &http.Server{
		Addr:         configs.Config.Server.Addr,
		Handler:      r,
		ReadTimeout:  configs.Config.Server.ReadTimeout * time.Second,
		WriteTimeout: configs.Config.Server.WriteTimeout * time.Second,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer shutdown(ctx)
	if err := s.Shutdown(ctx); err != nil {
		zapx.With(zap.Error(err)).Error("Server Shutdown error")
	}
	zapx.Info("Server exiting")
}
