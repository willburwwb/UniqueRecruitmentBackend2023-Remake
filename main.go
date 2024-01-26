package main

import (
	"UniqueRecruitmentBackend/internal/cmd"
)

// @title           UniqueStudio Recruitment API
// @version         0.1
// @description     This is API doc of UniqueStudio Recruitment.

// @contact.email  wwbstar07@gmail.com

// @host      https://dev.back.recruitment2023.hustunique.com/

// @externalDocs.description  飞书 doc
// @externalDocs.url todo

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
