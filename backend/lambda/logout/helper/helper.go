package helper

import "time"

func GenerateTime() int64 {
	return time.Now().Unix()
}
