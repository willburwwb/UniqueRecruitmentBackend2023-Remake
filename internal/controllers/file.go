package controllers

import (
	"UniqueRecruitmentBackend/global"
	"context"
	"mime/multipart"
)

func upLoadAndSaveFileToCos(file *multipart.FileHeader, fileName string) error {
	cosClient := global.GetCosClient()
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
