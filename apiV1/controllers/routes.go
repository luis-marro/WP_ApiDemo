package controllers

import (
	"github.com/gin-gonic/gin"
)

var r *gin.Engine
var apiv1 *gin.RouterGroup

// InitServer function to init the Gin Default server
func InitServer() *gin.Engine {
	r = gin.Default()
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

//getAllPartsRoute Route to load all the parts currently in the server
