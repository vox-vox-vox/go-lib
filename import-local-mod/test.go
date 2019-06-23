package main
import (
     anyname "github.com/ahuigo/go-hello"
     ahui2 "ahui1"
)
func main() {
    println("-------------start anyname.Test----------")
    anyname.Test()
    println("-------------start anyname.test2----------")
    anyname.Test2()
    println("-------------start ahui.Test3----------")
    ahui2.Test3()
    // ahui2 is package name
    ahui2 := ahui2.Add(1, 5)
    println(ahui2)
}
