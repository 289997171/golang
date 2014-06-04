package main


//import (
//	"bytes"
//	"io"
//	"io/ioutil"
//	"fmt"
//)
//
//func main() {
//	bytesBuffer := bytes.NewBuffer([]byte("1234567890"))
//	bytesBuffer1 := bytes.NewBuffer([]byte("qwertyuiop"))
//	bytesBuffer2 := bytes.NewBuffer([]byte("asdfghjkl"))
//	bytesBuffer3 := bytes.NewBuffer([]byte("zxcvbnm"))
//	ioReader := io.MultiReader(bytesBuffer, bytesBuffer1, bytesBuffer2, bytesBuffer3)
//	b, err := ioutil.ReadAll(ioReader)
//	fmt.Println(string(b), err)
//	//1234567890qwertyuiopasdfghjklzxcvbnm <nil>
//
//}
//
