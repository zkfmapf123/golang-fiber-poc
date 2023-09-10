package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringToInt(t *testing.T) {

	v1 := StringToInt("1234")
	v2 := StringToInt("4000")

	assert.Equal(t, v1, 1234)
	assert.Equal(t, v2, 4000)
}

func Test_StringToBool(t *testing.T) {

	v1 := StringToBool("true")
	v2 := StringToBool("false")

	assert.Equal(t, v1, true)
	assert.Equal(t, v2, false)
}
