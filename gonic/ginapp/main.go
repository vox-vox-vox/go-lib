package main

import (
    "flag"
	"github.com/gin-gonic/gin"
    "ginapp/server"
    "github.com/DeanThompson/ginpprof"
)

// This function's name is a must. App Engine uses it to drive the requests properly.
func main() {
	// gin+port
	r := gin.New()
    port := flag.String("p", "4500", "Public Server Port")
    flag.Parse()

	// Define your handlers
    server.Register(r)

    // automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	ginpprof.Wrap(r)
	// ginpprof also plays well with *gin.RouterGroup
	// group := r.Group("/debug/pprof")
	// ginpprof.WrapGroup(group)


	r.Run(":"+*port)
	//http.Handle("/", r)
}

