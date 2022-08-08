package main

import (
	"fmt"
	"os"
)

//$ go build command-line-arguments.go
//$ ./command-line-arguments a b c d
//[./command-line-arguments a b c d]
//[a b c d]
//c
func main() {
	args := os.Args
	argsWithoutProg := args[1:]
	arg3 := os.Args[3]

	fmt.Println(args)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg3)
}
