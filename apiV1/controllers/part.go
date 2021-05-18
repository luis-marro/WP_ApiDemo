package controllers

import (
	"WP_ApiDemo/apiV1/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// getAllPartsRoute godoc
// @Summary View All Parts
// @Description Handler function to register the route that loads all the parts from the database.
// @Produce json
// @Success 200 {array} model.Part
// @Failure 500 {object} map[string]string
// @Router /viewParts [get]
func getAllPartsRoute() {
	apiv1.GET("/viewParts", func(c *gin.Context) {
		parts, err := model.GetAllParts()
		if err != nil {
			log.Println("Error retrieving all the events from the database: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error loading the parts from the database",
			})
			return
		}
		c.JSON(http.StatusOK, parts)
	})
}

// getSpecificItem godoc
// @Summary Get a specific Part
// @Description Function to load an item from an ID received in the URL
// @Produce json
// @Param id path string true "the ID of the part to be reviewed"
// @Success 200 {object} model.Part
// @Failure 500 {object} map[string]string
// @Router /viewPart/{id} [get]
func getSpecificItem() {
	apiv1.GET("/viewParts/:id", func(c *gin.Context) {
		id := c.Param("id")
		foundPart, err := model.GetPartByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error finding the specific part",
			})
			return
		}
		log.Println(foundPart)
		// search the id in the local array
		c.JSON(http.StatusOK, foundPart)
	})
}

// searchPartByName godoc
// @Summary Search a part by name
// @Description Search all the parts in the DB that contain the Keywords passed in the query string
// @Produces json
// @Param searchQuery query string true "Keywords to lookup a part, separated by - "
// @Success 200 {array} model.Part
// @Failure 500 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /searchParts [get]
func searchPartByName() {
	apiv1.GET("/searchParts", func(c *gin.Context) {
		params := c.Request.URL.Query()
		if val, ok := params["searchQuery"]; ok {
			keywords := val[0]
			log.Println("Keywords: ", keywords)
			foundParts, err := model.GetPartByName(keywords)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Message": "Error searching the database for the requested words.",
				})
				return
			}
			c.JSON(http.StatusOK, foundParts)
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error, missing search keywords",
			})
			return
		}
	})
}

// sellPart godoc
// @Summary Diminish inventory for a part
// @Description Substract 1 to the inventory of an specific part
// @Produces json
// @Param partId query string true "Id of the part to which the inventory must be diminshed"
// @Success 200 {string} string
// @Failure 500 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /sellPart [delete]
func sellPart() {
	apiv1.DELETE("/sellPart", func(c *gin.Context) {
		params := c.Request.URL.Query()
		if val, ok := params["partId"]; ok {
			partId := val[0]
			err := model.DiminishInventory(partId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Message": "The part was not found",
					"Error":   err,
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"Message": "Part Sold!",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error, no valid Item ID",
			})
			return
		}
	})
}

// createPart godoc
// @Summary Create a new Part
// @Description Handler function to create a new part in the system
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /createPart [post]
func createPart() {
	apiv1.POST("/createPart", func(c *gin.Context) {
		var part *model.Part
		if err := c.BindJSON(&part); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message: ": "Error parsing body JSON",
				"Error":     err,
			})
			log.Println("Error parsing body for Create Part API: ", err)
			return
		}
		log.Println("Received Part: ", part)
		ref, err := model.CreateNewPart(part.Name, part.Description, part.Category, part.Price,
			part.Pictures, part.IsNew, part.Inventory, part.CarMake, part.CarModel)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message: ": "Error creating the part",
				"Error":     err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message: ":   "Part created Successfully",
			"createdPart": ref,
		})
		return
	})
}

// updatePart Handler function for the route that updates a part in the database
// @Summary Update part fields
// @Description Handler function to update the fields of a part in the system
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /updatePart [patch]
func updatePart() {
	apiv1.PATCH("/updatePart", func(c *gin.Context) {
		var part *model.Part
		if err := c.BindJSON(&part); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "The request contains an incorrect body",
				"Error":   err,
			})
			return
		}
		err := model.UpdatePart(part.Id, part.Name, part.Description, part.Category, part.Price, part.Cars,
			part.Pictures, part.IsNew, part.Inventory)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message:": "Error updating the part",
				"Error":    err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message: ": "The part was updated successfully",
		})
		return
	})
}
