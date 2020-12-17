package main

// matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
// detail: https://golang.org/src/regexp/example_test.go
import (
	"fmt"
	"regexp"
    "flag"
)

func compile(){
	var validID = regexp.MustCompile(`http(s)?://([\w\-]+\.hdmap\.momenta\.works|localhost|m)(:\d+)?$`)
	fmt.Println(validID.MatchString("http://dev-pm-ui.hdmap.momenta.works"))  //true
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))


    r := regexp.MustCompile(`\((?P<fatal>fatal: [^)]+)\)`)
    res := r.FindStringSubmatch("(fatal: do not exists)     ")
    fmt.Println("res:",res[0])
    fmt.Println("res:",res[1])

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
