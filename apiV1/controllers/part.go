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
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error, missing search keywords",
			})
			return
		}
	})
}
