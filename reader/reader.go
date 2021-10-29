package reader

import (
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
////ReadPDF ...读excel文件
//func (fr FileReader) ReadPDF(path string) (string, error) {
//	var content string
//	//f, err := pdf.Open("D:\\gowork\\src\\known01\\data\\生产环境数据检查.pdf")
//	//if err != nil {
//	//	fmt.Println(err,f)
//	//	return "",err
//	//}
//	return content, nil
//}
//
////ReadPDF ...读excel文件
//func (fr FileReader) ReadDoc(path string) (string, error) {
//	var content string
//	f, err := document.Open("D:\\gowork\\src\\known01\\data\\生产环境数据检查.docx")
//	if err != nil {
//		fmt.Println(err, f)
//		return "", err
//	}
//	fmt.Println(f.MergeFields())
//	fmt.Println(content)
//	return content, nil
//}

//ReadTxt ... reader txt file
func (fr FileReader) ReadTxt(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, nil
}
