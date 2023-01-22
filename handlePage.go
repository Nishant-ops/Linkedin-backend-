package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

func handlePage(c *gin.Context) {
	c.Writer.Header().Set("content-type", "application/json")
	var message Message
	requestBody, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "wrong body to marshall"})
	}
	 json.Unmarshal(requestBody, &message)

	c.JSON(http.StatusOK,message)

}
