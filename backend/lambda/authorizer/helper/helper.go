package helper

import "time"

func IsExpired(expiresAt int64) bool {
	return expiresAt < time.Now().Unix()
}
