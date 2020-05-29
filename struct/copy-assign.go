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
        wilson := Cat{7, "Wilson", []string{"Tom", "Tabata", "Willie"}}
        nikita := wilson        //浅copy Cat
        nikita.age = 8
        nikita.friends[0]= "newuser"

        fmt.Println(wilson)
        fmt.Println(nikita)
    }
