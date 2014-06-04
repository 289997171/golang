package main

//import (
//	"fmt"
//	"io"
//	"bytes"
//)
//
//func main() {
//	bytesReader := bytes.NewReader([]byte("1234567890a1234567890b1234567890c1234567890d1234567890"))
//	var off int64 = 2 //偏移，读取开始位置
//	var n int64 = 20 // 读取数据数量
//	ioSectionReader := io.NewSectionReader(bytesReader, off, n)
//	p := make([]byte, 10)
//	var o int64 = 0
//	nn, err := ioSectionReader.ReadAt(p, o) //偏移0位
//	fmt.Println(nn, err, p, string(p[:nn]))
//	//10 <nil> [51 52 53 54 55 56 57 48 97 49] 34567890a1
//
//	o = 1
//	nn, err = ioSectionReader.ReadAt(p, o) //偏移1位
//	fmt.Println(nn, err, p, string(p[:nn]))
//	//10 <nil> [52 53 54 55 56 57 48 97 49 50] 4567890a12
//
//	o = 2
//	nn, err = ioSectionReader.ReadAt(p, o) //偏移2位
//	fmt.Println(nn, err, p, string(p[:nn]))
//	//10 <nil> [53 54 55 56 57 48 97 49 50 51] 567890a123
//
//}

