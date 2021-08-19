package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	swagRouter(r)
	fileRouter(r)
	gormRouter(r)
	r.Any("/echo/*anypath", EchoServer)
	r.GET("/sleep/:second", sleepFunc)
	r.GET("/cpu/:second", cpuFunc)
	r.GET("/json/map", jsonMapFunc)
}

func jsonMapFunc(c *gin.Context) {
    m := map[string][]byte{
        "status": []byte("running!中国"),
    }
	c.JSON(http.StatusOK, m)
}

func cpuFunc(c *gin.Context) {
	seconds, _ := strconv.Atoi(c.Param("second"))
	n := longRun(seconds)
	msg := fmt.Sprintf("sleep second: %v s, n=%d\n", seconds, n)
	c.JSON(http.StatusOK, msg)
}

func longRun(seconds int) int {
	now := time.Now()
	end_time := now.Add(time.Duration(seconds) * time.Second)
	n := 0
	for ; end_time.After(now); now = time.Now() {
		for i := 0; i < 1e8; i += 1 {
			n += 1
		}
	}
	return n
}

func sleepFunc(c *gin.Context) {
	seconds, _ := strconv.Atoi(c.Param("second"))
	time.Sleep(time.Duration(seconds) * time.Second)
    fmt.Printf("%vs passed!\n", seconds)
	c.JSON(http.StatusOK, "sleep second: "+c.Param("second"))
}
