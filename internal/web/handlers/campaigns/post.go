package campaigns

import (
	"log"

	"github.com/braejan/go-leal-challenge/internal/domain/campaigns/model"
	"github.com/gin-gonic/gin"
)

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		usecases, err := getCampaignUsecases()
		if err != nil {
			c.JSON(500, err)
			return
		}
		var campaign model.Campaign
		err = c.ShouldBind(&campaign)
		if err != nil {
			c.JSON(400, err)
			return
		}
		err = usecases.CreateNewCampaign(campaign)
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
