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
			t.Log("message_id:", msgID)
			pickupInfo.MessageID = msgID

		}
	}
	t.Log("==================", pickupInfo)
}
