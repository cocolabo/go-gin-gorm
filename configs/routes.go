package configs

import (
	"github.com/cocolabo/go-gin-gorm/helpers"
	"github.com/cocolabo/go-gin-gorm/models"
	"github.com/cocolabo/go-gin-gorm/repositories"
	"github.com/cocolabo/go-gin-gorm/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(contactRepository *repositories.ContactRepository) *gin.Engine {
	route := gin.Default()

	routeGroup := route.Group("/contacts")
	{
		routeGroup.POST("", func(context *gin.Context) {
			var contact models.Contact

			err := context.ShouldBindJSON(&contact)

			if err != nil {
				response := helpers.GenerateValidationResponse(err)

				context.JSON(http.StatusBadRequest, response)

				return
			}

			code := http.StatusOK

			response := services.CreateContact(&contact, *contactRepository)

			if !response.Success {
				code = http.StatusBadRequest
			}

			context.JSON(code, response)
		})

		routeGroup.GET("", func(context *gin.Context) {
			code := http.StatusOK

			response := services.FindAllContacts(*contactRepository)

			if !response.Success {
				code = http.StatusBadRequest
			}

			context.JSON(code, response)
		})

		routeGroup.GET("/:id", func(context *gin.Context) {
			code := http.StatusOK

			response := services.FindOneContactById(context.Param("id"), *contactRepository)

			if !response.Success {
				code = http.StatusBadRequest
			}

			context.JSON(code, response)
		})
	}

	return route
}
