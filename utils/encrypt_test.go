package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Test_Base64 ... 测试base64 加解密的算法
func Test_Base64(t *testing.T) {
	src, encodeValue := "abcd1234", "YWJjZDEyMzQ="
	value := EncodingBase64([]byte(src))
	assert.Equal(t, value, encodeValue)
	decodeValue, err := DecodingBase64(value)
	assert.Nil(t, err)
	assert.Equal(t, src, decodeValue)
}
