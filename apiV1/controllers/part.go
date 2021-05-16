package controllers

import (
	"WP_ApiDemo/apiV1/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// getAllPartsRoute function to register the route that loads all the parts from the database.
func getAllPartsRoute() {
	apiv1.GET("/viewParts", func(c *gin.Context) {
		/*params := c.Request.URL.Query()
		start, startErr := strconv.Atoi(params["start"][0])
		limit, limitErr := strconv.Atoi(params["limit"][0])*/ // Parse Params Example
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

//getSpecificItem Function to load an item from an ID received in the URL
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

//createNewItem Function for POST request to create a new item
func createNewItem() {
	apiv1.POST("/createPart", func(c *gin.Context) {

	})
}

// searchPartByName Function to search a part in the system by the search parameters
// Returns a list with all the retrieved courses, and an error if applies
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

// sellPart Function to register the route used to diminish the inventory of a part by 1
func sellPart() {
	apiv1.DELETE("/sellPart", func(c *gin.Context) {
		params := c.Request.URL.Query()
		if val, ok := params["partId"]; ok {
			partId := val[0]
			err := model.DiminishInventory(partId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Message": "The part was not found",
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

// createPart Handler function to create a new part in the system.
func createPart() {
	apiv1.POST("/createPart", func(c *gin.Context) {
		var part *model.Part
		if err := c.BindJSON(&part); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message: ": "Error parsing body JSON",
			})
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
func updatePart() {
	apiv1.PATCH("/updatePart", func(c *gin.Context) {
		var part *model.Part
		if err := c.BindJSON(&part); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "The request contains an incorrect body",
			})
			return
		}
		err := model.UpdatePart(part.Id, part.Name, part.Description, part.Category, part.Price, part.Cars,
			part.Pictures, part.IsNew, part.Inventory)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message:": "Error updating the part",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message: ": "The part was updated successfully",
		})
		return
	})
}
