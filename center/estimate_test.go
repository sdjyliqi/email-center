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
	subject := "kai发票"
	subject = estTest.AmendSubject(subject)
	v, _ := estTest.GetCategory(subject)
	t.Log(v)
	estTest.AuditAllEmailItems()
}

func Test_AmendSubject(t *testing.T) {
	subject := "开发（piao）  ﹠開﹠发﹠缥﹠"
	newSubject := estTest.AmendSubject(subject)
	t.Log(newSubject)

}
