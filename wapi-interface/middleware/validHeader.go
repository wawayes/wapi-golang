package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wapi-interface/utils"
)

func ValidHeader(c *gin.Context) {
	header := c.GetHeader("mio")
	fmt.Println("header: ", header)
	if header != "accept" {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.NoAuth, "非法的访问"))
		c.Abort()
	}
	c.Next()
}
