package main
import (
    "bytes"
    "fmt"
    "log"
)

func main() {
    var (
        buf    bytes.Buffer
        logger = log.New(&buf, "prefix: ", log.Lshortfile)
    )

    logger.Print("log with file:lineno!")

    fmt.Print(&buf)
}
