package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//thewallmessage@gmail.com
//

//thewallmessage  thewall123456789
func Test_SendToMail(t *testing.T) {
	from, to, subject, text := "sdjyliqi@163.com", "sdjyliqi@163.com", "[thewall]Verification Code", []byte("Text Body is, of course, supported!")
	err := SendToMail(from, to, subject, text)
	assert.Nil(t, err)
}
