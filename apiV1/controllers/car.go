package controllers

import (
	"WP_ApiDemo/apiV1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// getAllCars godoc
// @Summary Gets a list of all the car makes that are present in the Database
// @Description Get only the makes for all the cars in the DB
// @ID get-all-the-cars
// @Accept json
// @Produce json
// @Success 200 {array} string
// @Failure 500 {object} map[string]string
// @Router /getAllCars [get]
func getAllCars() {
	apiv1.GET("/getAllCars", func(c *gin.Context) {
		listOfMakes, err := model.GetAllCars()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message:": "Error loading makes",
				"Error":    err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Makes": listOfMakes,
		})
	})
}

// getSpecificPart godoc
// @Summary Get Model for Make
// @Description Handler function for the route to lookup all car models of a specific make
// @Accept json
// @Produce json
// @Param carMake query string true "Car make to filter by"
// @Success 200 {array} string
// @Failure 500 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /getSpecificCar [get]
func getModelForMake() {
	apiv1.GET("/getSpecificCar", func(c *gin.Context) {
		params := c.Request.URL.Query()
		if val, ok := params["carMake"]; ok {
			listOfModels, err := model.GetModelsByMake(val[0])
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Message": "Error getting the make's models",
					"Error":   err,
				})
				return
			}
			c.JSON(http.StatusOK, listOfModels)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message:": "The car make was not included in the URL",
			})
			return
		}
	})
}

// getCarEngine godoc
// @Summary Get available engines for make and model
// @Description Handler function to lookup all the available engines for a specific make and model
// @Produce json
// @Param carMake query string true "Query params with the Car's Make"
// @Param carModel query string true "Query params with the Car's Model"
// @Success 200 {array} string
// @Failure 500 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /getCarEngine [get]
func getCarEngine() {
	apiv1.GET("/getCarEngine", func(c *gin.Context) {
		params := c.Request.URL.Query()
		if make, ok := params["carMake"]; ok {
			if carModel, modelOk := params["carModel"]; modelOk {
				engines, err := model.GetCarEngine(make[0], carModel[0])
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"Message": "Error getting information from the DB",
					})
					return
				}
				c.JSON(http.StatusOK, engines)
				return
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"Message": "Car model not provided in the request",
				})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Car make not provided in the request",
			})
			return
		}
	})
}
