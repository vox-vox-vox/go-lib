package main

import (
	"fmt"
	"regexp"
    "flag"
)

func compile(){
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	fmt.Println(validID.MatchString("adam[23]"))  //true
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))

}

func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
    // go run regexp.go -h
    cmd := flag.String("cmd", "compile", "a string")
    flag.Parse()
    if *cmd == "compile" {
        compile()
    }
}
