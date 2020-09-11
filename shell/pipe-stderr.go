package main

import (
    "bytes"
    "fmt"
    "log"
    "os/exec"
)

const ShellToUse = "bash"

func Shellout() (error, string, string) {
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command("sh", "-c", "echo 'sleep 3' && sleep 3 && echo 'sleep end';cmd-not-existed||echo ok")
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    fmt.Println("start run:")
    err := cmd.Run() // cmd.Run() 内含cmd.Wait()
    return err, stdout.String(), stderr.String()
}

func main() {
    err, out, errout := Shellout()
    if err != nil {
        log.Printf("error: %v\n", err)
    }
    fmt.Println("--- stdout ---")
    fmt.Println(out)
    fmt.Println("--- stderr ---")
    fmt.Println(errout)
}
