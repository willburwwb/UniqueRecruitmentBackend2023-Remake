package utils

import (
	"UniqueRecruitmentBackend/global"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func GetCOSObjectResp(filename string) (*cos.Response, error) {
	cosClient := global.GetCosClient()
	response, err := cosClient.Object.Get(context.Background(), filename, nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}
