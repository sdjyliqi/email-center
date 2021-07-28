package utils

import (
	"encoding/base64"
	"github.com/golang/glog"
)

//EncodingBase64 ... base64加密
func EncodingBase64(content []byte) string {
	return base64.StdEncoding.EncodeToString(content)
}

//DecodingBase64 ... base64解密
func DecodingBase64(str string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		glog.Errorf("DecodeString failed,err:%+v", err)
		return "", err
	}
	return string(decoded), nil
}
