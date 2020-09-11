package main
import (
    "github.com/creack/pty"
    "io"
    "os"
    "os/exec"
)

func main() {
    c := exec.Command("sh", "-c","sleep 5 && echo ok")
    f, err := pty.Start(c)
    if err != nil {
        panic(err)
    }
    println("start　不阻塞....")

    go func() {
        f.Write([]byte("foo\n"))
        f.Write([]byte("bar\n"))
        f.Write([]byte("baz\n"))
        f.Write([]byte{4}) // EOT
    }()
    io.Copy(os.Stdout, f)
}
