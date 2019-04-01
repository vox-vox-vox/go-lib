package main
import (
     anyname "github.com/ahuigo/go-hello"
     ahui2 "ahui1"
)
func main() {
    anyname.Test()
    anyname.Test2()
    ahui2.Test3()
    // ahui2 is package name
    ahui2 := ahui2.Add(1, 5)
    println(ahui2)
}
