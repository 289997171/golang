package main

//import (
//	"bytes"
//	"io"
//	"fmt"
//)
//
//func main() {
//	bytesReader := bytes.NewReader([]byte("1234567890"))
//	var off int64 = 2 //偏移，读取开始位置
//	var n int64 = 4 // 读取数据数量
//	ioSectionReader := io.NewSectionReader(bytesReader, off, n)
//	p := make([]byte, 10)
//	nn, err := ioSectionReader.Read(p)
//	fmt.Println(nn, err)
//	//4 <nil>
//	fmt.Println(p, string(p[:nn]))
//	//[51 52 53 54 0 0 0 0 0 0] 3456
//}

