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
//		b byte
//		err error
//	)
//	b, err = newReadSize.ReadByte()
//	fmt.Println(err, string(b))
//	//<nil> 1
//	b, err = newReadSize.ReadByte()
//	fmt.Println(err, string(b))
//	//<nil> 2
//
//	newReadSize.UnreadByte() //复恢缓冲区的数据，所以下面的读出的数据是2
//
//	b, err = newReadSize.ReadByte()
//	fmt.Println(err, string(b))
//	//<nil> 2
//
//}

