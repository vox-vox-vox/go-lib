package main
import "io/ioutil"
import "os"
import (
    "fmt"
    //"path/filepath"
)


func createTempFile(dir string) error{
    //dir := filepath.Split(path)
    os.MkdirAll(dir, os.ModePerm)
    file, err := ioutil.TempFile(dir, "tmp_")
    if err!=nil{
        return err
    }else{
        fmt.Printf("create file:%#v\n", file.Name())
    }
    return err
    //defer os.Remove(file.Name())
}


func main(){
    os.Mkdir("tmp", 0700)
    path := "tmp/a/b/c/d"
    err := createTempFile(path)
    if err!=nil{
        panic(err)
    }

}
