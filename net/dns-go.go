package main

import (
	"net"
	"fmt"
	"os"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    n:=int(1e6)
    wg.Add(n)
    for i := 0; i < n; i++ {
        go lookup(&wg)
    }
    wg.Wait()
    println("don")

}

func lookup(wg *sync.WaitGroup){
    ips, err := net.LookupIP("localhost")
    defer wg.Done()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
    return
	for _, ip := range ips {
		fmt.Printf("%s\n", ip.String())
	}
}
