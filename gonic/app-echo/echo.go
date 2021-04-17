package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This function's name is a must. App Engine uses it to drive the requests properly.
func main() {
	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	// Define your handlers
	r.Any("/*anypath", func(c *gin.Context) {
		// header
		res := c.Request.Method + " " + c.Request.URL.Path + "?" + c.Request.URL.RawQuery + " " + c.Request.Proto + "\n"
		for k, vs := range c.Request.Header {
			for _, v := range vs {
				res += k + ": " + v + "\n"
			}
		}
		res += "\n"

		// body
		buf, _ := ioutil.ReadAll(c.Request.Body)
		res += string(buf)

		c.String(http.StatusOK, res)
	})

	// Handle all requests using net/http
	r.Run(":8088")
	//http.Handle("/", r)
}
