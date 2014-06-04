package main

//import (
//	"bytes"
//	"bufio"
//	"fmt"
//)
//
//func main() {
//	var rb = bytes.NewBuffer([]byte("123456789012345678901234567890,\nabcd,\nAbcd,\nABCD,\nABcd"))
//	var newReadSize = bufio.NewReaderSize(rb, 16)
//	var (
//		line []byte
//		isPrefix bool //如果一行数据大小超出缓冲区大小，.ReadLine 返回缓冲区大小的数据长度，剩下的数据下次再返回。所以isPrefix 会被设置为true。
//		err error
//	)
//	var str string
//	for  {
//		line, isPrefix, err = newReadSize.ReadLine()
//
//		if isPrefix {
//			str += string(line)
//			// fmt.Println(err, isPrefix, str)
//		} else {
//			str += string(line)
//			fmt.Println(err, isPrefix, str)
//			str = ""
//		}
//
//		if err != nil {
//			break
//		}
//	}
//	//<nil> false 1234567890,
//	//<nil> false abcd,
//	//<nil> false Abcd,
//	//<nil> false ABCD,
//	//<nil> false ABcd
//	//EOF false
//
//}

