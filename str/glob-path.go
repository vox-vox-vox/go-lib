package main

import (
	"fmt"
	"path/filepath"
)
func main() {
    m,err:=filepath.Match("/home/catch/**", "/home/catch/foo/bar") ///不支持**
	fmt.Println("dir match", m, err) 
    m, err = filepath.Match("data[0-9]*", "data123.csv")
	fmt.Println("dir match", m, err) 
}
