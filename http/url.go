package main

import (
	"fmt"
    "net/url"
	_ "log"
	_ "reflect"
	_ "bytes"
)

func main() {
    pn := fmt.Println
    m, _:= url.ParseQuery(`x=1&y=2&y=3;z`)
	pn(m)

    u, _:= url.Parse("http://bing.com/search?q1=dotnet")
	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	fmt.Println(u.String())
}
