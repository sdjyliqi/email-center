package tryagain

import (
	"testing"
)

func Test_GetAddr(t *testing.T) {
	text := `我的籍贯是河北省灵寿县。我的单位地址是北京市海淀区汇通路12号中国光大银行1层，
    现居住地址为北京市丰台区太平桥西里1号楼907室， 
    老家地址为河北省石家庄市灵寿县塔上镇万里村，那边比较穷" +
    "我买的房子地址是北京市丰台区正商名苑1号楼1单位888
    我喜欢吃北京烤鸭，云南火腿，山西老陈醋
    `
	result := PickAddr(text)
	t.Log(result)
}
