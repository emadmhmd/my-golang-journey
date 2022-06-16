package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context, url string) {
	// Readthe body of main Request
	body, err := c.GetRawData()
	if err != nil {
		fmt.Println("Fail to read the body of main request request")
	}
	// Make new Request
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		fmt.Println("Fail to make new post request")
	}
	// Get main request headers and set them to new request
	for key := range c.Request.Header {
		req.Header.Set(key, c.Request.Header.Get(key))
	}

	// Make new client to send request
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Fail to send new post request")
	}

	// To close  at the end
	defer response.Body.Close()

	// To hold the body of response in result
	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)

	// Check if the result have data
	if len(result) == 0 {
		c.JSON(response.StatusCode, gin.H{
			"error": "Fail to post the data ",
		})
		return
	}

	// Attach result
	c.JSON(response.StatusCode, result)
}

func Get(c *gin.Context, url string) {
	// Make new Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Fail to make new get request")
	}
	// Get main request headers and set them to new request
	for key := range c.Request.Header {
		req.Header.Set(key, c.Request.Header.Get(key))
	}

	// Make new client to send request
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Fail to send new get request")
	}

	// To close  at the end
	defer response.Body.Close()

	// To hold the body of response in result
	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)

	// Check if the result have data
	if len(result) == 0 {
		c.JSON(500, gin.H{
			"error": "Fail to get data ",
		})
		return
	}

	// Attach result
	c.JSON(response.StatusCode, result)
}
