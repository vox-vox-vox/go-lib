package main

import (
	"fmt"
    "net/url"
	"net/http"
    "io/ioutil"
	_ "log"
	_ "reflect"
	_ "bytes"
)

func main() {
    //resp, _ := http.PostForm("https://httpbin.org/cookies/set?ahui=mo",
    resp, _ := http.PostForm("https://httpbin.org/post",
	url.Values{"key": {"Value"}, "id[]": {"123"}})
    pn := fmt.Println
    body, _:= ioutil.ReadAll(resp.Body)
    pn(string(body))
}
