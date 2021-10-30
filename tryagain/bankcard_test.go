package tryagain

import "testing"

func TestBankCardPick(t *testing.T) {
	var tests = []struct {
		input  string
		output bool
	}{
		{",6217856100006167719,", true},
		{",6223631800000665572,", true},
		{",6230662433002269891,", true},
		{",4693803227213619,", true},
		{",6259063290566796,", true},
		{",6214831046190432,", true},
		{",6217856100072455410,", true},
		{",6225561646149135,", true},
		{",6226622006275867,", true},
		{",6217993000077079286,", true},
		{",6212260200196866244,", true},
		{",370502102541981,", true},
	}
	for _, test := range tests {
		bankcard := PickBankCard(test.input)
		//t.Logf("%v", bankcard)
		if (bankcard != nil) != test.output {
			t.Errorf("输入内容为\"%s\"，提取到的银行卡号为：\"%s\"", test.input, bankcard)
		}
	}
}
