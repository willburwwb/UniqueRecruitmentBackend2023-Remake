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
// @externalDocs.url https://uniquestudio.feishu.cn/docx/Yh96d2DoyoCe6zxlR0ecSU5snDd?from=from_copylink

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
