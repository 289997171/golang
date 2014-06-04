package main

//import (
//	"bytes"
//	"bufio"
//	"fmt"
//)
//
//func main() {
//	var newBuffer_a *bytes.Buffer = bytes.NewBuffer([]byte("1234567890"))
//	var newBuffer_b *bytes.Buffer = bytes.NewBuffer([]byte("中国人说中国话"))
//
//	var newWriterSize *bufio.Writer = bufio.NewWriterSize(newBuffer_a, 10)
//	newWriterSize.ReadFrom(newBuffer_b) // 追加数据
//	newWriterSize.Flush()
//
//	fmt.Println(newBuffer_a.String())
//	//1234567890中国人说中国话
//
//}
//
