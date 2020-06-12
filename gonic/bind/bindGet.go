package main
import (
    "github.com/gin-gonic/gin"
    "net/http"
    _"fmt"
)
type Params struct {
    Text     string `form:"text" binding:"required"`
}

func main() {
    router := gin.Default()

    router.GET("/get", func(c *gin.Context) {
        var params Params
        if err := c.ShouldBind(&params); err == nil {
            c.JSON(http.StatusOK, gin.H{"text": params.Text})
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
    })

    router.Run(":8088")
}

