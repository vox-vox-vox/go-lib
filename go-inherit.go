/**
Animal 通过interface 继承其实是多余的, 而是通过结构体继承
type Pet interface {
    Name() string
    Speak() string
    Play()
}
type Dog interface {
    Pet
    Breed() string
}


*/
package main

import (
    "fmt"
)

// Det interface struct func define
type pet struct {
    speaker func() string
    name    string
}
func (p *pet) Name() string {
    return p.name
}
func (p *pet) Speak() string {
    //missing spaker = Factory.speak
    return p.speaker()
}
func (p *pet) Play() {
    fmt.Println(p.Speak())
}
func (p *pet) speak() string {
    return fmt.Sprintf("my name is %v ", p.name)
}
func NewPet(name string) *pet {
    p := &pet{
        name: name,
    }
    p.speaker = p.speak
    return p
}

// Dog define
type dog struct {
    pet
    breed string
}
func (d *dog) speak() string {
   return fmt.Sprintf("speak(): %v\n", d.pet.speak())
}

func NewDog(name, breed string) *dog {
    d := &dog{
        pet:   pet{name: name},
        breed: breed,
    }
    d.speaker = d.speak
    return d
}



func main() {
    d := NewDog("spot", "pointer")
    fmt.Println(d.Name())
    fmt.Println(d.Speak())
    d.Play()
}
