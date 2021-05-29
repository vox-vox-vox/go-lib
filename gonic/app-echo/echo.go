package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
    "flag"

	"github.com/gin-gonic/gin"
)

// This function's name is a must. App Engine uses it to drive the requests properly.
func main() {
	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	// Define your handlers
	r.Any("/*anypath", func(c *gin.Context) {
		// header
		res := c.Request.Method + " " +
			c.Request.Host +
			c.Request.URL.Path + "?" + c.Request.URL.RawQuery + " " + c.Request.Proto + " " +
			c.ClientIP() + "\n"
		headers := sortHeaders(c)
		for _, kv := range *headers {
			res += kv[0] + ": " + kv[1] + "\n"
		}
		res += "\n"

		// handler
		sendCookie(c)
        sendHeaders(c)

		// body
		buf, _ := ioutil.ReadAll(c.Request.Body)
		res += string(buf)

		c.String(http.StatusOK, res)
	})

	// Handle all requests using net/http
    port := flag.String("p", "8090", "Public Server Port")
    flag.Parse()

	r.Run(":"+*port)
	//http.Handle("/", r)
}

func sortHeaders(c *gin.Context) *[][2]string {
	headers := [][2]string{}
	for k, vs := range c.Request.Header {
		for _, v := range vs {
			headers = append(headers, [2]string{k, v})
		}
	}
	n := len(headers)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			jj := j - 1
			h1, h2 := headers[j], headers[jj]
			if h1[0] < h2[0] {
				headers[jj], headers[j] = headers[j], headers[jj]
			}
		}
	}
	return &headers
}

func sendHeaders(c *gin.Context){
    c.Writer.Header().Set("Remote", "echo-server by ahuigo")
    c.Writer.Header().Set("Location", "http://baidu1.com")
}

func sendCookie(c *gin.Context) {
	// c.Request.URL.
	hostname := strings.Split(c.Request.Host, ":")[0]
	countStr, _ := c.Cookie("count")
	if countStr == "" {
		countStr = "1"
	} else {
		count, _ := strconv.Atoi(countStr)
		countStr = strconv.Itoa(count + 1)
	}

	c.SetCookie("count", countStr, 86400, "", hostname, false, false)
	// fmt.Printf("h:%#v\n", c.Header)
}
