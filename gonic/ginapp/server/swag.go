package server

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func swagRouter(r *gin.Engine) {
    s:=r.Group(`/swag`)
	s.GET("read", swagReadHandler)
}

func swagReadHandler(c *gin.Context) {
	filepath := c.DefaultQuery("path", "tmp/a.txt")
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	res := string(buf)
	c.String(http.StatusOK, res)

}

