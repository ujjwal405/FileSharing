package helper

import (
	"fmt"
	"time"
)

func GenerateUniqueCode(keyID string) string {
	now := time.Now()
	futureTime := now.Add(30 * time.Minute)
	timestamp := futureTime.Format("20060102150405") // YYYYMMDDHHMMSS
	nano := futureTime.Nanosecond()
	code := fmt.Sprintf("%s-%s%09d", keyID, timestamp, nano)
	return code
}
