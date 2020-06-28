package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Test struct {
	Content string
}

func main() {
	t := new(Test)
	t.Content = "http://www.baidu.com?id=123&test=1<>"
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(t)
	fmt.Println(bf.String())
}
