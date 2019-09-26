package zip

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZip(t *testing.T) {
	err := Unzip("../uploads/a.zip", "/abc", "/abc")
	assert.Nil(t, err)
}

func TestStatic(t *testing.T) {
	fmt.Println(isStaticFile(".dasdas.dasasda.d"))
}
