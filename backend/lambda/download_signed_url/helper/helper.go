package helper

import (
	"fmt"
	"time"

	apierror "github.com/ujjwal405/FileSharing/download_signed_url/apiError"
)

// IsCodeExpired checks if the time encoded in the code has passed
func IsCodeExpired(code string) error {
	// Validate code length (23 characters expected)
	if len(code) != 23 {
		return apierror.InvalidCode()
	}

	// Split code into timestamp (14 chars) and nanoseconds (9 chars)
	timestampStr := code[:14] // YYYYMMDDHHMMSS
	nanoStr := code[14:]      // 9-digit nanoseconds

	// Parse the timestamp portion into a time.Time
	codeTime, err := time.Parse("20060102150405", timestampStr)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %v", err)
	}

	// Parse nanoseconds into an integer
	var nano int
	_, err = fmt.Sscanf(nanoStr, "%09d", &nano)
	if err != nil {
		return fmt.Errorf("failed to parse nanoseconds: %v", err)
	}

	// Add nanoseconds to get the exact time
	codeTime = codeTime.Add(time.Duration(nano) * time.Nanosecond)

	// Compare with current time
	now := time.Now()
	isExpired := now.After(codeTime)
	if isExpired {
		return apierror.InvalidCode()
	}
	return nil
}
