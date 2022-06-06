package service

import (
	"fmt"
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

func GetWordCount(c *gin.Context) {
	var requestBody Request
	var wcs wc.WordCountInterface
	// Bind the request Body
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	// Create Word Count Service
	wcs = wc.New()
	// Create the filers
	punctFilter := wc.NewPunctFilter()
	// Add the filter into Word Count Service
	wcs = wcs.(wc.WordCountService).Add(punctFilter)

	// Get Words Along with frequencies
	wordCount, err := wcs.CountWords(requestBody.Message)

	// Error Handling
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Set the response status and word count as response body
	c.IndentedJSON(http.StatusOK, wordCount)
}
