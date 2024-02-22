package utils

import (
	"UniqueRecruitmentBackend/global"
	"context"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"time"
)

func GenerateCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := ""
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(r.Intn(10))
	}

	return code
}

func GenerateTmpCode(ctx context.Context, id string, expire time.Duration) (string, error) {
	code := GenerateCode()
	err := global.GetRedisCli().Set(ctx, id, code, expire).Err()
	if err != nil {
		zapx.WithContext(ctx).Error("generate tmp code failed", zap.Error(err), zap.String("id", id))
		return "", err
	}
	return code, nil
}

func GetTmpCodeByID(ctx context.Context, id string) (code string, err error) {
	value := global.GetRedisCli().GetDel(ctx, id)
	if err = value.Err(); err != nil {
		zapx.WithContext(ctx).Error("getdel by id failed", zap.Error(err), zap.String("id", id))
		return "", err
	}
	if err = value.Scan(&code); err != nil {
		return
	}
	return
}
