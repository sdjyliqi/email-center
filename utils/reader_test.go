package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFileNames(t *testing.T) {
	dir := "D:\\gowork\\src\\email-center\\data\\金融诈骗"
	files, err := GetFileNames(dir, "eml")
	t.Log(files, err)
	assert.Nil(t, err)

}
