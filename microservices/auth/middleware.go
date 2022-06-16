package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Services() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.RequestURI
		segments := strings.Split(url, "/")
		segmentOne := strings.Split(segments[1], "?")
		//IsService(c, segments[1])

		services := Serviceslist()

		serviceUrl := services[segmentOne[0]]
		if serviceUrl == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Not Found",
			})
			c.Abort()
			return
		}
		fullUrl := serviceUrl + "/" + segmentOne[0]
		authRoutes := AuthRoutelist()
		var token Token
		for route := range authRoutes {
			if strings.Contains(url, route) {
				fullUrl = serviceUrl + "/" + segmentOne[0] + "/" + route
				authorized := c.GetHeader("Authorization")
				if authorized == "" {
					c.JSON(401, gin.H{
						"error": "Unauthorized",
					})
					c.Abort()
					return
				}
				db.Where("token = ?", authorized).First(&token)
				if token.ID == 0 {
					c.JSON(401, gin.H{
						"error": "Unvalied token",
					})
					c.Abort()
					return
				}
				c.Request.Header.Set("userId", *token.UserId)
			}
		}

		fmt.Println("Full url", fullUrl)
		method := strings.ToLower(c.Request.Method)
		switch method {
		case "get":
			Get(c, fullUrl)
		case "post":
			Post(c, fullUrl)
		}
		// c.JSON(http.StatusOK, gin.H{
		// 	"error": "Found",
		// })
	}
}

func Serviceslist() map[string]string {
	m := make(map[string]string)
	// m[""] = "http://localhost:8080"
	m["users"] = "http://localhost:6060"
	m["posts"] = "http://localhost:5050"
	return m
}

func AuthRoutelist() map[string]bool {
	m := make(map[string]bool)
	// m[""] = "http://localhost:8080"
	m["my-posts"] = true
	return m
}

func IsService(c *gin.Context, route string) {
	m := make(map[string]bool)
	m[""] = true
	m["users"] = true
	m["posts"] = true
	if ok := m[route]; !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})
		return
	}
}
