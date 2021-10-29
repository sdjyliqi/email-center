package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReadExcel(t *testing.T) {
	var reader FileReader
	content, err := reader.ReadExcel("")
	assert.Nil(t, err)
	t.Log(content)
}

func Test_ReadPDF(t *testing.T) {
	var reader FileReader
	content, err := reader.ReadDoc("")
	assert.Nil(t, err)
	t.Log(content)
}

func Test_ReadPPT(t *testing.T) {

}
