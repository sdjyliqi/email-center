package center

import (
	"email-center/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

//判断维度是否是否为真假
var estTest *Estimate

func Test_Estimate(t *testing.T) {
	var err error
	estTest, err = CreateEstimate()
	assert.Nil(t, err)
	condition := "is_identify=0"
	items, err := model.BodyModel.GetItemsByCondition(condition, 0, 10)
	assert.Nil(t, err)
	t.Log(len(items))
	estTest.AuditAllEmailItems()

	result := estTest.ReplaceContentByAC("我的手机号⑴①零")
	t.Log(result)
}
