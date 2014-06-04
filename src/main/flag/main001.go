package main

import (
	"flag"
	"fmt"
	"os"
)

var name *string = flag.String("name", "vicky", "enter your name")
var age *int = flag.Int("age", 27, "enter your age")
var email *string = new(string)

func init() {
	flag.StringVar(email, "email", "eclipser@163.com", "enter your email")
	flag.Parse()
}



func main() {
	fmt.Println("os.Args", os.Args)
	fmt.Println("flag.Args()", flag.Args(), "flag.NArg()", flag.NArg())
	fmt.Println("name", *name)
	fmt.Println("age", *age)
	fmt.Println("email", *email)


//go run main001.go -name=jack -age=199 -email=xxx@xxx 1 2 3 4 5 6
//os.Args [/tmp/go-build738785662/command-line-arguments/_obj/exe/main001 -name=jack -age=199 -email=xxx@xxx 1 2 3 4 5 6]
//flag.Args() [1 2 3 4 5 6] flag.NArg() 6
//name jack
//age 199
//email xxx@xxx
}

