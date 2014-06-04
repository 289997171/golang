package main

//import (
//	"bytes"
//	"bufio"
//	"fmt"
//)
//
//// 功能和.ReadBytes或.ReadString一样
//func main() {
//	var rb = bytes.NewBuffer([]byte("1234567890,abcd,Abcd,ABCD,ABcd"))
//	var newReadSize = bufio.NewReaderSize(rb, 1000000)
//	var (
//		delim byte = ','
//		line []byte
//		err error
//	)
//	for  {
//		line, err = newReadSize.ReadSlice(delim)
//		fmt.Println(err, string(line))
//		if err != nil {
//			break
//		}
//	}
//	//<nil> 1234567890,
//	//<nil> abcd,
//	//<nil> Abcd,
//	//<nil> ABCD,
//	//EOF ABcd
//
//}

