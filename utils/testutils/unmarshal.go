package testutils

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"hemtjan.st/zcl"
	"testing"
)

func Unmarshal(t *testing.T, v zcl.Val, str string) {
	data, _ := hex.DecodeString(str)
	out, err := v.UnmarshalZcl(data)
	assert.Nil(t, err)
	assert.Empty(t, out)
}
