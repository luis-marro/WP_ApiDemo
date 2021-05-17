package controllers

import (
	"WP_ApiDemo/apiV1/model"
	"crypto/sha256"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// createUser Handler function to create a new user in the Database
func createUser() {
	apiv1.POST("/createUser", func(c *gin.Context) {
		var reqUser model.User
		if c.BindJSON(&reqUser) == nil {
			log.Println("Parsed Object ", reqUser)
			pwAsByte := []byte(reqUser.UnsafePassword + model.PasswordSalt)
			hasher := sha256.New()
			hasher.Write(pwAsByte)
			reqUser.SafePassword = string(pwAsByte)
			reqUser.UnsafePassword = ""
			if err := model.CreateNewUser(reqUser); err == nil {
				log.Println("User Successfully added to the Database")
				c.JSON(http.StatusOK, gin.H{
					"Status": "Success",
				})
				return
			} else {
				log.Println("Error creating the user in the database ", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"Message": "Error",
					"Error":   err,
				})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error",
				"Error":   "The body JSON was not structured correctly",
			})
			return
		}
	})
}

// findUser Handler function to lookup an user in the database.
func findUser() {
	apiv1.POST("/findUser", func(c *gin.Context) {
		var reqUser model.User
		if c.BindJSON(&reqUser) == nil {
			pwAsByte := []byte(reqUser.UnsafePassword + model.PasswordSalt)
			hasher := sha256.New()
			hasher.Write(pwAsByte)
			reqUser.SafePassword = string(pwAsByte)
			if foundUser, err := model.FetchUser(reqUser.Email, reqUser.SafePassword); err == nil {
				c.JSON(http.StatusOK, foundUser)
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Message": "Error retrieving the user information",
					"Error":   err,
				})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error parsing the user from the Body",
				"Error":   "The body of the request was in incorrect format.",
			})
			return
		}
	})
}
