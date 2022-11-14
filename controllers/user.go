package controllers

import (
	"ads/database"
	"ads/helpers"
	"ads/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetInforUser(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": helpers.MiddlewareErr,
		})
		return
	}
	uid := user.(models.User).Id
	fmt.Printf("i' i` iiii", uid)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  user,
	})
}

func ListUser(c *gin.Context) {
	user, exists := c.Get("user")
	fmt.Printf("Error %s", user)
	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "helpers.MiddlewareErr",
		})
		return
	}
	var accountUser []models.User
	database.DB.Find(&accountUser)
	c.JSON(http.StatusOK, &accountUser)
}

func DetailUser(c *gin.Context) {
	user, exists := c.Get("user")
	fmt.Printf("Error %s", user)
	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": helpers.MiddlewareErr,
		})
		return
	}
	var accountUser []models.User
	database.DB.First(&accountUser)
	c.JSON(http.StatusOK, &accountUser)
}
func DeleteUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": helpers.MiddlewareErr,
		})
		return
	}
	uid := user.(models.User).RoleId

	fmt.Printf("", uid)
	var delete_product models.User
	database.DB.Where("id = ?", c.Param("id")).Delete(&delete_product)
	c.JSON(http.StatusOK, &delete_product)
}
