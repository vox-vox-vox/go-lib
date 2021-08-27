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
	validID = regexp.MustCompile(`^([a-z\d]+-)?osm[a-z\d\-]+\.hdmap\.momenta\.works$`)
	fmt.Println(validID.MatchString("osm3.hdmap.momenta.works"))  //true
	fmt.Println(validID.MatchString("dev-osm3.hdmap.momenta.works"))  //true
	fmt.Println(validID.MatchString("staging-osm3.hdmap.momenta.works"))  //true
	fmt.Println(validID.MatchString("staging.osm3.hdmap.momenta.works"))  //true


    r := regexp.MustCompile(`\((?P<fatal>fatal: [^)]+)\)`)
    res := r.FindStringSubmatch("(fatal: do not exists)     ")
    fmt.Println("res[0]:",res[0])
    fmt.Println("res[1]:",res[1])

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
