package controllers

import (
	"github.com/dashboard_poc.com/client"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type LogMessage struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	RequestID string    `json:"request_id"`
	UserID    string    `json:"user_id"`
}

var endpoint = "https://dashboard-demo.kraftonamericas.com/logs"

type UploadController struct{}

func (h UploadController) UploadData(c *gin.Context) {
	osClient := client.NewClient(endpoint)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Match Controller Create had an error when reading the request body")
		c.JSON(http.StatusInternalServerError, nil)
		c.Abort()
	}
	body := strings.NewReader(string(bytes))
	code, err := osClient.PostRequest(body)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Match Controller Create had an error when reading the request body")
		c.JSON(http.StatusInternalServerError, nil)
		c.Abort()
	}
	c.JSON(code, "Success")
}
