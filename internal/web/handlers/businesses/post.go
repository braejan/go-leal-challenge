package businesses

import (
	"log"

	"github.com/braejan/go-leal-challenge/internal/domain/businesses/model"
	"github.com/gin-gonic/gin"
)

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		service, err := getBusinessUsecases()
		if err != nil {
			c.JSON(500, err)
			return
		}
		var business model.Business
		err = c.ShouldBind(&business)
		if err != nil {
			c.JSON(400, err)
			return
		}
		err = service.CreateBusiness(business)
		if err != nil {
			log.Println("Error Create: ", err)
			if err.Error() == "user already exist" {
				c.JSON(409, err)
				return
			}
			c.JSON(503, err)
			return
		}
		c.Status(201)
	}
}
