package main
import "io/ioutil"
import "os"
import "fmt"

func main(){
    os.Mkdir("tmp", 0700)
    file, err := ioutil.TempFile("tmp", "prefix")
    if err!=nil{
        fmt.Println(err)
    }else{
        fmt.Println("crate file:", file)
    }
}
