package utils

import (
	"regexp"
	"strings"
	"testing"
)

func Test_ChkAdv(t *testing.T) {
	content := "Vx:361212981你好1212，他的微信:sdjyliqi 一五九三二七的八折"
	var qqFormat = "().折"
	formatRegx := regexp.MustCompile(qqFormat)
	values := formatRegx.FindAllStringSubmatch(strings.ToLower(content), -1)
	t.Log(values)

}
