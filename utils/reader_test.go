package utils

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetFileNames(t *testing.T) {
	dir := "D:\\gowork\\src\\email-center\\data\\金融诈骗"
	files, err := GetFileNames(dir, "eml")
	t.Log(files, err)
	assert.Nil(t, err)

}

func Test_PickupEmailBody(t *testing.T) {
	content := "asasasas Content-Type: text/plain; charset=utf-8\nContent-Transfer-Encoding: base64\n\n77yRM+KSjuWNg++8keKSi++8r+KSjuaekO+8l+KSiu+8mO+8kQ0KJm5ic3A7DQpocmNl\nYg0K5piv4oCc5a+E5Lq64oCd55qEDQombmJzcDsNCuS9huiLj+i9vOiupOS4ug0KMDPm\nl7YzNuWIhjU056eS\n\n----boundary_587039_3b834ecb-b29f-458c-b123-de6d2ed53204\nContent-Type: text/html; charset=utf-8\nContent-Transfer-Encoding: base64\n\nPFA+77yRM+KSjuWNg++8keKSi++8r+KSjuaekO+8l+KSiu+8mO+8kTwvUD4NCjxQPiZu\nYnNwOzwvUD4NCjxQPmhyY2ViPC9QPg0KPFA+5piv4oCc5a+E5Lq64oCd55qEPC9QPg0K\nPFA+Jm5ic3A7PC9QPg0KPFA+5L2G6IuP6L286K6k5Li6PC9QPg0KPFA+MDPml7YzNuWI\nhjU056eSPC9QPg==\n----boundary_587039_3b834ecb-b29f-458c-b123-de6d2ed53204--"
	result := PickupEmailBody(content)
	t.Log(result)

}

func TestReadEmail(t *testing.T) {
	msgIDPrefix, subjectPrefix := "Message-ID: <", "Subject:"
	pickupInfo := Email{}
	name := "D:\\gowork\\src\\email-center\\data\\金融诈骗\\20210625050710__sheila@tedstones.com__2021062505071015783270889696479419.eml"
	content, err := ReadEmail(name)
	t.Log(err)
	idx := strings.Index(string(content), "\n")
	t.Log(idx)
	//var sepRegx = "^\n$"
	idx = strings.Index(string(content), "\r\n\r")

	header := string(content)[0:idx]
	t.Log(header)

	body := string(content)[idx:len(string(content))]
	pickupInfo.ContentBody = body
	info := strings.Split(header, "\r")
	for _, v := range info {
		line := strings.Trim(v, " ")
		line = strings.ReplaceAll(line, "\n", "")
		//t.Log(k,line)
		switch {
		case strings.HasPrefix(line, "From:"):
			from := line[7 : len(line)-1]
			t.Log("from:", from)
			pickupInfo.From = from
		case strings.HasPrefix(line, "To:"):
			t.Log("to:", line)

		case strings.HasPrefix(line, "Date:"):
			sendTime := line[6 : len(line)-1]
			t.Log("Date:", sendTime)
			pickupInfo.Date = sendTime
		case strings.HasPrefix(line, subjectPrefix):
			subject := line[len(subjectPrefix) : len(line)-1]
			t.Log("Subject:", subject)
			pickupInfo.Subject = subject
		case strings.HasPrefix(line, msgIDPrefix):
			msgID := line[len(msgIDPrefix) : len(line)-1]
			pickupInfo.MessageID = msgID
		}
	}

}
