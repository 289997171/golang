package main

//import (
//	"fmt"
//	"bytes"
//)
//
//func main() {
//	s := []byte("汉字，中国人的文字。123456789012345678901234567890")
//
//	bytesBuffer := bytes.NewBuffer(s)
//	bytesBuffer.Grow(10) // 设置每次缓冲数据10字节
//	p := make([]byte, 20)
//	n, _ := bytesBuffer.Read(p)
//	fmt.Println(string(p[:n]))
//	//<nil> 汉字，中国人
//
//
//
//	bytesBuffer.Write([]byte("hello world"))
//	fmt.Println(bytesBuffer.Len(), bytesBuffer)
//}
