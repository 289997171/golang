package main

//import (
//	"fmt"
//	"bytes"
//	"bufio"
//)
//
//func main() {
//	var rb = bytes.NewBuffer([]byte("1234567890,abcd,Abcd,ABCD,ABcd"))
//	var newReadSize = bufio.NewReaderSize(rb, 1000000)
//	var (
//		b rune
//		size int
//		err error
//	)
//	b, size, err = newReadSize.ReadRune()
//	fmt.Println(err, size, string(b))
//	//<nil> 1 1
//	b, size, err = newReadSize.ReadRune()
//	fmt.Println(err, size, string(b))
//	//<nil> 1 2
//
//	newReadSize.UnreadRune() //复恢缓冲区的数据，所以下面的读出的数据是2
//
//	b, size, err = newReadSize.ReadRune()
//	fmt.Println(err, size, string(b))
//	//<nil> 1 2
//
//}


