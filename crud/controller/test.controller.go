package controller

import (
	"crud/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get test successfully",
	})
}
func PostTest(c *gin.Context) {
	var test model.Test
	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// db.Create(test)
}
func PutTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Put test successfully",
	})
}
func PatchTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Patch test successfully",
	})
}
func DeleteTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete test successfully",
	})
}
