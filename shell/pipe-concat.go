package main

import (
    "bytes"
    "io"
    "os"
    "os/exec"
)

func main() {
    cmd := exec.Command("cat")

    var cmd bytes.Buffer
    cmd.Stdout = &b2

    if err := cmd.Run(); err != nil {
		return err
	}
    cmd.Close()
    cmd.Run()
    io.Copy(os.Stdout, &b2)
}
