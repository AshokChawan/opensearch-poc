package client

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type ClientInfo struct {
	client   *http.Client
	endpoint string
}

func NewClient(url string) *ClientInfo {
	client := &ClientInfo{
		client:   &http.Client{Timeout: 5 * time.Second},
		endpoint: url,
	}
	return client
}

func (c *ClientInfo) PostRequest(body *strings.Reader) (int, error) {
	fmt.Println("PostRequest API Called")
	req, err := http.NewRequest(http.MethodPost, c.endpoint, body)
	if err != nil {
		log.Error().
			Err(err).
			Msg("PostRequest had an error when reading the request body")
		return http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		log.Error().
			Err(err).
			Msg("PostRequest had an error while sendig request")
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()
	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().
			Err(err).
			Msg("PostRequest had an error when reading the request body")
		return http.StatusInternalServerError, err
	}
	fmt.Printf("Response  %s\n", string(respbody))
	return resp.StatusCode, nil
}
