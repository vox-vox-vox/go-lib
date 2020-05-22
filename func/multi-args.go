package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
    return y, x
}
func multiArgsHandler(args ...interface{}) () {
    println(args[0].(string), args[1].(string))
}

func main() {
	fmt.Println(swap("a","b")) 
	multiArgsHandler(swap("a","b")) 
	println(swap("a","b")) 
}

