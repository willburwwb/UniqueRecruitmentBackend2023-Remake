package global

import (
	"UniqueRecruitmentBackend/configs"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"net/http"
	"net/url"
)

var cosClient *cos.Client

func GetCosClient() *cos.Client {
	return cosClient
}

func setupCOS() {
	u, _ := url.Parse(configs.Config.COS.CosUrl)
	b := &cos.BaseURL{BucketURL: u}
	cosClient = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  configs.Config.COS.CosSecretID,
			SecretKey: configs.Config.COS.CosSecretKey,
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory rerror.
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
	})
}
