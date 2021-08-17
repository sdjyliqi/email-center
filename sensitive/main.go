package main

import (
	"fmt"
	"io/ioutil"
)

func writeWords() error {
	m := createMap()

	content := "package utils \n"
	content = content + "var TagPropertyDict = map[string]LegalTag{ \n"

	for k, v := range m {
		line := fmt.Sprintf("    \"%s\":      %s,\n", k, v)
		content = content + line
	}
	content = content + "}\n"
	fmt.Println(content)
	err := ioutil.WriteFile("D:\\go_workplace\\src\\email-center\\utils\\words.go", []byte(content), 0666) //写入文件(字节数组)
	fmt.Println(err)
	return err
}
func main() {
	fmt.Println("======================")
	writeWords()
}
