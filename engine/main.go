package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type engineResponse struct {
	Code int         `json:"code" `
	Msg  string      `json:"msg" `
	Data interface{} `json:"data" `
}

var ErrArgs = engineResponse{
	Code: 1,
	Msg:  "输入的参数不能为空，检查参数",
	Data: nil,
}

func main() {
	args := os.Args
	if len(args) != 2 {
		content, _ := json.Marshal(ErrArgs)
		fmt.Println(string(content))
		return
	}
	result, err := pickupFiles(args[1])
	if err != nil {
		ErrArgs.Msg = err.Error()
		content, _ := json.Marshal(ErrArgs)
		fmt.Println(string(content))
		return
	}
	ret := engineResponse{
		Code: 0,
		Msg:  "succ",
		Data: result,
	}
	content, _ := json.Marshal(ret)
	fmt.Println(string(content))
}
