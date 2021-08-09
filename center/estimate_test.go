package center

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//判断维度是否是否为真假

func Test_Estimate(t *testing.T) {
	est, err := CreateEstimate()
	assert.Nil(t, err)
	subject := "kai发票"
	subject = est.AmendSubject(subject)
	v := est.GetCategory(subject)
	t.Log(v)
}
