package main


//import (
//	"bytes"
//	"io"
//	"fmt"
//)
//
//func main() {
////	例子(数据长度不足情况)：
//	bytesReader := bytes.NewReader([]byte("1234567890a1234567890b1234567890c1234567890d1234567890"))
//	var off int64 = 2 //偏移，读取开始位置
//	var n int64 = 20 // 读取数据数量
//	ioSectionReader := io.NewSectionReader(bytesReader, off, n)
//	p := make([]byte, 2)
//	nn, err := ioSectionReader.Read(p)
//	fmt.Println(nn, err, p, string(p[:nn]))
//	//2 <nil> [51 52] 34
//
//	//解释： off 偏移到第2位置，读取2个数据，指针在4位置
//	var offset int64  = 4   //再次偏移4个位置
//	var whence int = 0      //0 表示从开头位置开始偏移4个位置，不是从指针位置哦!!
//	ioSectionReader.Seek(offset, whence)
//	p = make([]byte, 8)
//	nn, err = ioSectionReader.Read(p)
//	fmt.Println(nn, err, p, string(p[:nn]))
//	//8 <nil> [55 56 57 48 97 49 50 51] 7890a123
//
//	//解释：从当时指针开始偏移
//	//      这个和其它不一样，要注意。“已经截取8位字节”，后再从指针位置开始偏移4个位置。这样就剩下4位字节
//	offset = 4 //再次偏移4个位置
//	whence = 1 // 1 表示从当前指针位置开始偏移4个位置
//	ioSectionReader.Seek(offset, whence)
//	p = make([]byte, 8)
//	nn, err = ioSectionReader.Read(p)
//	fmt.Println(nn, err, p, string(p[:nn]))
//	//4 <nil> [56 57 48 98 0 0 0 0] 890b
//
//	//解释：从当时指针开始偏移-12个位置
//	//      这个和上面不一样，要注意。“已经截取10位字节”这个字节是“当前指针位置”到“-10位置”的字节
//	//      读取20个字节，因为已经截取到10位字节，所以无法读取更多字符
//	//      （错误的）可能你想设置offset为0或更多，因为2是结束位置，不是开始位置。所以无法读取数据。返回是 EOF
//	offset = -10 //再次偏移-12个位置
//	whence = 2 // 2 表示从当前指针位置开始偏移-12个位置
//	ioSectionReader.Seek(offset, whence)
//	p = make([]byte, 20)
//	nn, err = ioSectionReader.Read(p)
//	fmt.Println(nn, err, p, string(p[:nn]))
//	//10 <nil> [50 51 52 53 54 55 56 57 48 98 0 0 0 0 0 0 0 0 0 0] 234567890b
//
//}

