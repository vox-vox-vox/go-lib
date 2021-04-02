package main
 
import (
	"fmt"
)
 
    type Cat struct {
        age     int
        name    string
        friends []string
    }

    func main() {
        a := &Cat{}
        fmt.Printf("%#v\n",a.friends)
        fmt.Printf("%#v\n",a.friends==nil)
        fmt.Printf("%#v\n",[]string(nil))
    }
