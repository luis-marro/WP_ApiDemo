package apiV1

import (
	"WP_ApiDemo/apiV1/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var PartsList = []model.Part{
	model.Part{
		Id:          1,
		Name:        "Turbo Mitsubishi",
		Description: "Turbo Mitsubishi compatible con Volvo",
		Price:       1299.99,
		ImgPath:     "",
	},
	model.Part{
		Id:          2,
		Name:        "Pastillas de Freno Ceramicas",
		Description: "Pastillas de Freno ceramicas resistentes a altas temperaturas",
		Price:       450.00,
		ImgPath:     "",
	},
	model.Part{
		Id:          3,
		Name:        "Radiador",
		Description: "Radiador compatilble con Motores Mazda SkyActiv 2.0 y 2.5",
		Price:       875.90,
		ImgPath:     "",
	},
}

func InitServer() *gin.Engine {
	r := gin.Default()
	// group the urls
	apiv1 := r.Group("api/v1")
	getRoute(apiv1)
	getSpecificItem(apiv1)

	return r
}

func getRoute(api *gin.RouterGroup) {
	api.GET("/viewParts", func(c *gin.Context) {
		c.JSON(http.StatusOK, PartsList)
	})
}

func getSpecificItem(api *gin.RouterGroup) {
	api.GET("/viewParts/:id", func(c *gin.Context) {
		id := c.Param("id")
		var specificPart model.Part = model.Part{}
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error, item id must be an integer",
			})
			return
		}
		// search the id in the local array
		for _, s := range PartsList {
			if s.Id == parsedId {
				specificPart = s
			}
		}
		// response
		if (model.Part{}) == specificPart {
			// return not found
			c.JSON(http.StatusNotFound, gin.H{
				"Status":  "404",
				"Message": "Item not found",
			})
		} else {
			// return the object
			c.JSON(http.StatusOK, specificPart)
		}
	})
}
