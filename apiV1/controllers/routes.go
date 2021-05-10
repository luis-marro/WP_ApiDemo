package controllers

import (
	"github.com/gin-gonic/gin"
)

var r *gin.Engine
var apiv1 *gin.RouterGroup

// InitServer function to init the Gin Default server
func InitServer() *gin.Engine {
	r = gin.Default()
	r.Use(CORSMiddleware())
	// group the urls
	apiv1 = r.Group("api/v1")
	getAllPartsRoute()
	getSpecificItem()
	searchPartByName()
	sellPart()
	updatePart()
	createPart()

	// Routes for the cars
	getAllCars()
	getModelForMake()
	getCarEngine()
	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
