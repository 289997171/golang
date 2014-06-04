package main

//import (
//	"bytes"
//	"io"
//	"fmt"
//)
//
//func main () {
//	bytesBuffer := bytes.NewBuffer([]byte("1234567890"))
//	bytesBuffer1 := bytes.NewBuffer(nil)
//	ioReader := io.TeeReader(bytesBuffer, bytesBuffer1)
//	p := make([]byte, 6)
//	n, err := ioReader.Read(p) // 这里读取出多少，bytesBuffer1 将会写入多少。
//	fmt.Println(n, string(p), err)
//	//6 123456 <nil>
//	fmt.Println(bytesBuffer1.String())
//	//123456
//
//}
//
