package main

import (
	// "crud/controller"
	"crud/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB = nil
var err error

func main() {
	dsn := "host=localhost user=postgres password=0000 dbname=Test port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Fail to connect to postgres db")
	}
	r := gin.Default()
	db.AutoMigrate(&model.Test{})

	r.GET("/test", GetTests)
	r.GET("/test/:id", GetTest)
	r.POST("/test", PostTest)
	r.PUT("/test/:id", PutTest)
	r.PATCH("/test/:id", PatchTest)
	r.DELETE("/test/:id", DeleteTest)

	r.Run(":8080")
}

func GetTests(c *gin.Context) {
	var tests []model.Test
	limitStr := c.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.DefaultQuery("offset", "0")
	offset, _ := strconv.Atoi(offsetStr)

	db.Limit(limit).Offset(offset).Find(&tests)
	c.JSON(http.StatusOK, gin.H{
		"message": "Get the tests successfully",
		"data":    tests,
	})
}
func GetTest(c *gin.Context) {
	id := c.Param("id")
	var test model.Test
	db.First(&test, id)
	if test.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "This test Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get the test successfully",
		"data":    test,
	})
}
func GetTestById(c *gin.Context) model.Test {
	id := c.Param("id")
	var test model.Test
	db.First(&test, id)
	if test.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "This test not found",
		})
	}
	return test
}
func PostTest(c *gin.Context) {
	var test model.Test
	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to save new test",
		})
		return
	}
	test.Mobile = "01157689977"
	db.Create(&test)
	c.JSON(http.StatusOK, gin.H{
		"message": "Post the test successfully",
		"data":    test,
	})
}
func PutTest(c *gin.Context) {
	id := c.Param("id")
	var oldTest model.Test
	db.First(&oldTest, id)
	if oldTest.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "this test Not Found",
		})
		return
	}

	var newTest model.Test
	if err := c.ShouldBindJSON(&newTest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to Upadte the test",
		})
		return
	}
	oldTest.Email = newTest.Email
	oldTest.Name = newTest.Name
	oldTest.Mobile = newTest.Mobile
	db.Save(&oldTest)
	c.JSON(http.StatusOK, gin.H{
		"message": "Update the test successfully",
		"data":    oldTest,
	})
}
func PatchTest(c *gin.Context) {
	id := c.Param("id")
	var oldTest model.Test
	db.First(&oldTest, id)
	if oldTest.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "This test Not Found",
		})
		return
	}

	var newTest model.Test
	if err := c.ShouldBindJSON(&newTest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to Upadte the test",
		})
		return
	}
	oldTest.Email = newTest.Email
	oldTest.Name = newTest.Name
	oldTest.Mobile = newTest.Mobile
	db.Save(&oldTest)
	c.JSON(http.StatusOK, gin.H{
		"error":   "",
		"message": "Update the test successfully",
		"data":    oldTest,
	})
}
func DeleteTest(c *gin.Context) {
	id := c.Param("id")
	var test model.Test
	db.First(&test, id)
	if test.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "This test Not Found",
		})
		return
	}
	db.Delete(&test)
	// db.Unscoped().Delete(&test) // delete it hard
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete test by id successfully",
		"data":    test,
	})

}
