package main

//import (
//	"bytes"
//	"bufio"
//	"fmt"
//)
//
//func main() {
//	var rb = bytes.NewBuffer([]byte("1234567890,abcd,Abcd,ABCD,ABcd"))
//	var newReadSize = bufio.NewReaderSize(rb, 1000000)
//	var (
//		delim byte = ','
//		line string
//		err error
//	)
//	for  {
//		line, err = newReadSize.ReadString(delim)
//		fmt.Println(err, line)
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
//
