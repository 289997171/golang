package main

//import (
//	"fmt"
//	"bytes"
//	"bufio"
//)
//
//func main() {
//	var rb = bytes.NewBuffer([]byte("1234567890,abcd,Abcd,ABCD,4ABcd"))
//	var newReadSize = bufio.NewReaderSize(rb, 1000000)
//	var (
//		delim byte = '4'
//		buf []byte
//		err error
//	)
//	for  {
//		buf, err = newReadSize.ReadBytes(delim)  // 相当与在缓冲区查找 ‘4’ 结尾
//		fmt.Println(err, string(buf))
//		if err != nil {
//			break
//		}
//	}
////	<nil> 1234
////	<nil> 567890,abcd,Abcd,ABCD,4
////	EOF ABcd
//
//}

