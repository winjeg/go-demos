package otpauth

import (
	"fmt"
	"github.com/czxichen/otpauth"
	"time"
)

func GenerateCode(key string) string {
	now := time.Now().Unix()
	code, _, err := otpauth.GenerateCode(key, now)
	if err == nil {
		return fmt.Sprintf("%06d", code)
	}
	return ""
}
