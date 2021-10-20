package center

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//判断维度是否是否为真假
var estTest *Estimate

func Test_Estimate(t *testing.T) {
	var err error
	estTest, err = CreateEstimate()
	assert.Nil(t, err)
	estTest.AuditAllEmailItems()
}
