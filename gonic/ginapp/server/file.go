package server

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func fileRouter(r *gin.Engine) {
	r.GET("/f", fileReadHandler)
}

// EchoHandler _
func fileReadHandler(c *gin.Context) {
	filepath := c.DefaultQuery("path", "tmp/a.txt")
	buf, err := ioutil.ReadFile(filepath)
	c.Writer.Header().Set("Content-type", "text/html; charset=utf-8")
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	res := string(buf)
	c.String(http.StatusOK, res)

}
