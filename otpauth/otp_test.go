package otpauth

import (
	"fmt"
	"testing"
)

func TestOtp(t *testing.T) {
	fmt.Println(GenerateCode("R2ZROX7DT7K632PA"))
}
