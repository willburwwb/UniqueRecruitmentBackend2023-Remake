package utils

import "fmt"

func FormatMsg(msg string, msgArgs ...interface{}) string {
	return fmt.Sprintf(msg, msgArgs...)
}
