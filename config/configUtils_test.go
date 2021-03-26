package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValueSuccess(t *testing.T) {
	str, err := GetValue("mysql", "ipAddrees")

	assert.Equal(t, nil, err)

	assert.Equal(t, "127.0.0.1", str)
}
func TestGetValueFail(t *testing.T) {
	_, err := GetValue("mysql", "ipAddrees2")

	assert.NotEqual(t, "127.0.0.1:6379", err)
}
