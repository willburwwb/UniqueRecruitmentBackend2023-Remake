package global

import "C"
import (
	"UniqueRecruitmentBackend/configs"
)
import "github.com/mozillazg/go-cos"

var cosClient *cos.Client

func GetCosClient() *cos.Client {
	return cosClient
}
func setCOS() {
	s := configs.CosConfigs.Url
	cosUrl, _ := cos.NewBaseURL(configs.CosConfig)
}
