package main


//import (
//	"bytes"
//	"bufio"
//	"fmt"
//)
//
//func main() {
//	rb := bytes.NewBuffer([]byte("1234567890"))
//	r := bufio.NewReader(rb)
//	fmt.Println(r.Buffered()) //这里的时候还不会显示出缓冲区字节大小
//	//0
//	var (
//		buf [4]byte
//		n int
//		err error
//	)
//	n, err = r.Read(buf[:]) // 只有调用.Read 数据才会开始加载入缓冲区。
//	fmt.Println(n, err, r.Buffered(), string(buf[:])) //这里.Buffered 就可以判断出缓冲区里面还有多少数据。
//	//4 <nil> 6 1234
//
//	n, err = r.Read(buf[:])
//	fmt.Println(n, err, r.Buffered(), string(buf[:]))
//	//4 <nil> 2 5678
//}

