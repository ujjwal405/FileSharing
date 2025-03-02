package helper

import "time"

func GenerateExpiryTime() int64 {
	return time.Now().Add(30 * 24 * time.Hour).Unix()
}
