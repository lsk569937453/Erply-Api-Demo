package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenTokenSuccess(t *testing.T) {
	assert := assert.New(t)

	token, err := GenToken("lsk")
	assert.Equal(err, nil)
	users, err := ParseToken(token)
	assert.Equal(err, nil)
	name := users.Username
	assert.Equal(name, "lsk")
}
func TestParseTokenFail(t *testing.T) {
	assert := assert.New(t)
	_, err := ParseToken("testxxxx")
	assert.NotEqual(err, nil)
}
func TestParseTokenSuccess(t *testing.T) {
	assert := assert.New(t)
	users, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJleHAiOjE2MTkyNTEwNTIsImlzcyI6ImVycGx5LWFwaSJ9.IpvcmHPTi4ubbVs00MKLPbe034oO9P43ko4dOpCHKFY")
	assert.Equal(err, nil)
	name := users.Username
	assert.Equal(name, "test")
}
