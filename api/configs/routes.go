package configs

import (
	helpers2 "github.com/cocolabo/go-gin-gorm/api/helpers"
	models2 "github.com/cocolabo/go-gin-gorm/api/models"
	repositories2 "github.com/cocolabo/go-gin-gorm/api/repositories"
	services2 "github.com/cocolabo/go-gin-gorm/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(contactRepository *repositories2.ContactRepository) *gin.Engine {
	route := gin.Default()

	routeGroup := route.Group("/contacts")
	{
		routeGroup.POST("", func(context *gin.Context) {
			var contact models2.Contact

			err := context.ShouldBindJSON(&contact)

			if err != nil {
				response := helpers2.GenerateValidationResponse(err)

				context.JSON(http.StatusBadRequest, response)

				return
			}

			code := http.StatusOK

			response := services2.CreateContact(&contact, *contactRepository)

			if !response.Success {
				code = http.StatusBadRequest
			}

			context.JSON(code, response)
		})

		routeGroup.GET("", func(context *gin.Context) {
			code := http.StatusOK

			response := services2.FindAllContacts(*contactRepository)

			if !response.Success {
				code = http.StatusBadRequest
			}

			context.JSON(code, response)
		})

		routeGroup.GET("/:id", func(context *gin.Context) {
			code := http.StatusOK

			response := services2.FindOneContactById(context.Param("id"), *contactRepository)

			if !response.Success {
				code = http.StatusBadRequest
			}

			context.JSON(code, response)
		})
	}

	return route
}
