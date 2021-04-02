package main
import (
    "github.com/gin-gonic/gin"
    "net/http"
    _ "fmt"
)

type UserReq struct {
	Password string `gorm:"not null" json:"password" form:"password"`
}

func main() {
    router := gin.Default()

    router.POST("/", func(c *gin.Context) {
        req := UserReq{}
        if err := c.ShouldBind(req); err != nil {
            c.AbortWithError(http.StatusBadRequest, err)
            return
        }
        c.JSON(http.StatusOK, "succ")
    })

    router.Run(":8088")
}

