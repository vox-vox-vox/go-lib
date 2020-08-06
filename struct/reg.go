package main
import (
    "fmt"
    "regexp"
)

func main(){
    origin:= "http://a.ahuigo:999"
    r := regexp.MustCompile("^http(s)?://(\\w+\\.ahuigo\\.com|localhost|m)(:[0-9]+)?$")
	v := r.MatchString(origin)
    fmt.Println(v)
}
