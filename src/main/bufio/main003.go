package main

//import (
//	"fmt"
//	"bufio"
//)
//
//func main() {
//	var (
//		b []byte
//		advance int
//		token []byte
//		err error
//	)
//	b = []byte("中国人说中国话\r\n123456789")
//	advance, token, err = bufio.ScanRunes(b, false)
//	fmt.Println(advance, token, err)
//	//3 [228 184 173] <nil>
//	advance, token, err = bufio.ScanRunes(b, false)
//	fmt.Println(advance, token, err)
//	//3 [228 184 173] <nil>
//	advance, token, err = bufio.ScanRunes(b, true)
//	fmt.Println(advance, token, err)
//	//3 [228 184 173] <nil>
//	advance, token, err = bufio.ScanRunes(b, true)
//	fmt.Println(advance, token, err)
//	//3 [228 184 173] <nil>
//
//	b = []byte{}
//	advance, token, err = bufio.ScanRunes(b, true)
//	fmt.Println(advance, token, err)
//	//0 [] <nil>
//	advance, token, err = bufio.ScanRunes(b, false)
//	fmt.Println(advance, token, err)
//	//panic: runtime error: index out of range
//	//
//	//goroutine 1 [running]:
//	//bufio.ScanRunes(0x5008dc, 0x0, 0x0, 0x0, 0x0, ...)
//	//  C:/Users/ADMINI~1/AppData/Local/Temp/2/bindist664137607/go/src/pkg/bufio/scan.go:234 +0x1e6
//	//main.main()
//	//  G:/学习/go/test.go:38 +0x5f3
//	//exit status 2
//
//}

