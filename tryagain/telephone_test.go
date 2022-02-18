package tryagain

import "testing"

func TestPickTelephone(t *testing.T) {
	var tests = []struct {
		input  string
		output bool
	}{
		{"尊敬的王先生,您近期违规用卡逾期未还,我行于今日冻结您名下信用卡进黑名单,如需恢复正常请致电:031100797943【招商银行】", false},
		{"截止今晚24点我行将自动从您银行卡上扣除年费1200元。如有疑问,咨询电话:08776613854-08776615821 【农行通知】", true},
		{"尊敬的储蓄卡用户:您于本月18日在广晟百货用卡购买电器9886元,此款将从您帐上扣,如有问题请联 系0762-3926317 工商银行。", true},
		{"温馨提示:我行将于20: 00之前扣除您信用卡年费1280元,如有疑问详情请致电农行客服中心4008530010 【中国农业银行】", false},
		{"尊敬的用户:您的电子密码器即将失效,请尽快登录手机银行http//wap.95588op.com/升级维护01081234567232387给您带来不便敬请谅解《工商银行》", true},
		{"尊敬的用户:您的电子密码器即将失效,请尽快登录手机银行http//wap.95588op.com/升级维护01091234567给您带来不便敬请谅解《工商银行》", false},
		{"尊敬的工行用户:您的账户积分58513即将逾期清空,请登录手机网wap.noicco.com兑换988.5元现金,逾期失效【工商银行】", false},
	}
	for _, test := range tests {
		output1 := PickTelephone(test.input)
		t.Log(output1)
		if (output1 != nil) != test.output {
			t.Errorf("输入内容为\"%s\"，提取到的银行卡号为：\"%s\"", test.input, output1)
		}
	}
}
