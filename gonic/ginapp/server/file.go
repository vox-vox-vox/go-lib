package server

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func fileRouter(r *gin.Engine) {
	r.GET("/file/read", fileReadHandler)
}

// EchoHandler _
func fileReadHandler(c *gin.Context) {
	filepath := c.DefaultQuery("path", "a.txt")
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	res := string(buf)
	c.String(http.StatusOK, res)

}
