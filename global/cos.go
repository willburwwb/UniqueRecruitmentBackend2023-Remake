package global

import (
	"UniqueRecruitmentBackend/configs"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"mime/multipart"
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

func UpLoadAndSaveFileToCos(file *multipart.FileHeader, fileName string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	_, err = cosClient.Object.Put(context.Background(), fileName, src, nil)
	if err != nil {
		return err
	}
	return nil
}

func GetCOSObjectResp(filename string) (*cos.Response, error) {
	response, err := cosClient.Object.Get(context.Background(), filename, nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}
