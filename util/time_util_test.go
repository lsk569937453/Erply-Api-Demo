package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentTime(t *testing.T) {
	assert := assert.New(t)

	timex := GetCurrentTime()
	assert.Equal(len(timex), len(TimeFormat))

}
