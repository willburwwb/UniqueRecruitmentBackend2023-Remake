package main

import (
	"UniqueRecruitmentBackend/global"
	"UniqueRecruitmentBackend/internal/server"
)

// @titile Swagger api
// @version 1.0
// @description  This is backend of recruitment system for Unique Studio.
// @BasePath /api/v1/

func main() {
	if err := global.Setup(); err != nil {

	}
	s := server.NewServer()
	s.ListenAndServe()
}
