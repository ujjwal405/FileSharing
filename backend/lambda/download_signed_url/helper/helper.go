package helper

import (
	"fmt"
	"time"

	apierror "github.com/ujjwal405/FileSharing/download_signed_url/apiError"
)

// IsCodeExpired checks if the time encoded in the code has passed
func IsCodeExpired(code string) (string, error) {
	if len(code) != 60 || code[36] != '-' {
		return "", apierror.InvalidCode()
	}

	uid := code[:36] // e.g., "550e8400-e29b-41d4-a716-446655440000"

	// Extract timestamp+nano portion (last 23 characters)
	timePart := code[37:]         // "20250304123456123456789"
	timestampStr := timePart[:14] // "20250304123456"
	nanoStr := timePart[14:]

	// Parse the timestamp portion into a time.Time
	codeTime, err := time.Parse("20060102150405", timestampStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse timestamp: %v", err)
	}

	// Parse nanoseconds into an integer
	var nano int
	_, err = fmt.Sscanf(nanoStr, "%09d", &nano)
	if err != nil {
		return "", fmt.Errorf("failed to parse nanoseconds: %v", err)
	}

	// Add nanoseconds to get the exact time
	codeTime = codeTime.Add(time.Duration(nano) * time.Nanosecond)

	// Compare with current time
	now := time.Now()
	isExpired := now.After(codeTime)
	if isExpired {
		return "", apierror.InvalidCode()
	}
	return uid, nil
}
