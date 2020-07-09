package main

import (
	"fmt"
	"path/filepath"
)
func main() {
	fmt.Println("dir match",filepath.Match("/home/catch/**", "/home/catch/foo/bar")) ///不支持**
}
