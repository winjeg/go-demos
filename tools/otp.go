package tools

import (
	"fmt"
	"time"

	"github.com/czxichen/otpauth"
)

func GenerateCode(key string) string {
	now := time.Now().Unix()
	code, _, err := otpauth.GenerateCode(key, now)
	if err == nil {
		return fmt.Sprintf("%06d", code)
	}
	return ""
}
