package main

import (
    "fmt"
)

type A interface{
    Play()
}
func Test(o A){
    o.Play()
}

type Pet struct {
    age int
}
func (p *Pet) Play() {
    fmt.Println("showc age:", p.age)
}

func main() {
    //pf := fmt.Printf

    Test(&Pet{20}) //ok
    //pet.Play()

}
