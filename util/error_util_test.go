package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecoverSuccess(t *testing.T) {
	// assert := assert.New(t)
	defer ReturnError()
	panic("test")
}
func TestCheckTypeSuccess(t *testing.T) {
	assert := assert.New(t)
	result := CheckType("a")
	assert.Equal(result, "a")
	inResult := CheckType(10)
	assert.Equal(inResult, "int type")
	errResult := CheckType(fmt.Errorf("common error"))
	assert.Equal(errResult, "common error")
	otherResult := CheckType(nil)
	assert.Equal(otherResult, "not found type")
}
