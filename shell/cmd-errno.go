package main

import "os/exec"
import "fmt"

func main() {
    cmd := exec.Command("sh", "-c", "exit 2")

    if err := cmd.Run() ; err != nil {
        if exitError, ok := err.(*exec.ExitError); ok {
            fmt.Printf("Exit Status: %d", exitError.ExitCode())
        }else{
            fmt.Println("err:", err)
        }
    }

}
