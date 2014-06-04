package main

//import (
//	"bytes"
//	"bufio"
//	"fmt"
//)
//
//func main() {
//	var rb = bytes.NewBuffer([]byte("中国AB12"))
//	var newReadSize = bufio.NewReaderSize(rb, 1000000)
//	var (
//		utf8 rune
//		size int
//		err error
//	)
//	for  {
//		utf8, size, err = newReadSize.ReadRune()
//		fmt.Println(err, size, string(utf8))
//		if err != nil {
//			break
//		}
//	}
//	//<nil> 3 中
//	//<nil> 3 国
//	//<nil> 1 A
//	//<nil> 1 B
//	//<nil> 1 1
//	//<nil> 1 2
//	//EOF 0
//
//}

