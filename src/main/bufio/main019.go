package main

//import (
//	"bufio"
//	"bytes"
//	"fmt"
//)
//
//func main() {
//	var (
//		b []byte
//	)
//	b = []byte("1234567890")
//	var newReader *bufio.Reader = bufio.NewReaderSize(bytes.NewBuffer(b), 10)
//	var newBuffer_bb *bytes.Buffer = bytes.NewBuffer([]byte("中国人说中国话 1234567890 abcdef "))
//	newReader.WriteTo(newBuffer_bb) // 追加数据
//	fmt.Println(newBuffer_bb.String())
//	//中国人说中国话 1234567890 abcdef 1234567890
//
//}
//
