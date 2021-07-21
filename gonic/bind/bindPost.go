package main
import (
    "github.com/gin-gonic/gin"
    "net/http"
    _ "fmt"
)

type UserReq struct {
	Password string `gorm:"not null" json:"password" form:"password"`
}

/**
nc -l 8088
curl  'm:8088' -F 'password=ahui1'
curl  -X POST -H 'Content-Type: application/x-www-form-urlencoded' 'm:8088' -d 'password=ahui1'

 */
func main() {
    router := gin.Default()

    router.POST("/", func(c *gin.Context) {
        req := UserReq{}
        if err := c.ShouldBind(&req); err != nil {
            c.AbortWithError(http.StatusBadRequest, err)
            return
        }
        c.JSON(http.StatusOK, req)
    })

    router.Run(":8088")
}

