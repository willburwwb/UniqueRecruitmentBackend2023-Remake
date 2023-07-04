package utils

import (
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
