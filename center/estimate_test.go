package center

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//判断维度是否是否为真假
var estTest *estimate

func Test_Estimate(t *testing.T) {
	var err error
	estTest, err = CreateEstimate()
	assert.Nil(t, err)
	estTest.AuditAllEmailItems()
}

func Test_AmendSubject(t *testing.T) {
	var err error
	estTest, err = CreateEstimate()
	assert.Nil(t, err)
	subject := "@#$%^开发  （piao）!~  ﹠開﹠发﹠缥﹠"
	amendChars := []rune{}
	chars := []rune(subject)
	for _, v := range chars {
		t.Log(v)
		if v < 'A' || v > 'z' && v <= 255 {
			continue
		}
		amendChars = append(amendChars, v)
	}
	newSubject := estTest.AmendSubject(string(amendChars))
	t.Log(newSubject)
}
