package api

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"wapi-interface/utils"
)

func (Api) GetSentence(c *gin.Context) {
	n := rand.Intn(20)
	c.JSON(http.StatusOK, utils.ResponseOK(Sentences[n]))
}
