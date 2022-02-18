package reader

import (
	"fmt"
	"testing"
)

func Test_ReadPDF(t *testing.T) {
	res, err := readPdf("E:\\gowork\\email-center\\doc\\test2222.pdf")
	fmt.Println(res, err)
}

func Test_ReadPPT(t *testing.T) {

}
