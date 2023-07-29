package main

import (
	"UniqueRecruitmentBackend/internal/cmd"
)

// @titile Swagger api
// @version 1.0
// @description  This is backend of recruitment system for Unique Studio.
// @BasePath /api/v1/

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
