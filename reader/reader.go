package reader

import (
	"bytes"
	"fmt"
	"github.com/ledongthuc/pdf"
	"io/ioutil"
)

var FReader FileReader

type FileReader struct {
}

////ReadExcel ...读excel文件
//func (fr FileReader) ReadExcel(path string) (string, error) {
//	var content string
//	f, err := excelize.OpenFile("D:\\gowork\\src\\known01\\data\\test.xlsx")
//	if err != nil {
//		fmt.Println(err)
//		return "", err
//	}
//	sheetCount := f.SheetCount
//	for i := 0; i < sheetCount; i += 1 {
//		sheetNames := f.GetSheetName(i)
//		rows, err := f.GetRows(sheetNames)
//		fmt.Println(rows, err)
//		if err != nil {
//			return "", err
//		}
//		sheetContent, err := json.Marshal(rows)
//		if err != nil {
//			return "", err
//		}
//		content += string(sheetContent)
//	}
//	return content, nil
//}
//

func (fr FileReader) ReadTxtFromPdf(path string) (string, error) {
	pdf.DebugOn = true
	content, err := readPdf("E:\\gowork\\email-center\\doc\\test1030.pdf") // Read local pdf file
	if err != nil {
		return "", err
	}
	fmt.Println(content)
	return content, nil
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	fmt.Println(r.NumPage())
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}

//ReadTxt ... reader txt file
func (fr FileReader) ReadTxt(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, nil
}
