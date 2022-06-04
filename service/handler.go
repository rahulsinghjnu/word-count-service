package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	wc "github.com/rahulsinghjnu/word-count-service/word-count"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	WC map[string]int `json:"wc"`
}

func GetWordCount(c *gin.Context) { // gin.Context parameter.
	var requestBody Request
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}
	wordCount, _ := wc.New().CountWords(requestBody.Message)
	c.IndentedJSON(http.StatusOK, wordCount)
}
