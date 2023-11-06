package campaigns

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		usecases, err := getCampaignUsecases()
		if err != nil {
			c.JSON(500, err)
			return
		}
		ID := c.Param("id")
		if ID == "" {
			c.JSON(400, "id is empty")
			return
		}
		uuidID, err := uuid.Parse(ID)
		if err != nil {
			c.JSON(400, fmt.Sprintf("invalid id: %s\n", err))
		}
		business, err := usecases.GetCampaignByID(uuidID)
		if err != nil {
			// Validate not found
			if err == gorm.ErrRecordNotFound {
				c.JSON(404, "not found")
				return
			}
			c.JSON(500, err)
			return
		}
		c.JSON(200, business)

	}
}

func GetAllByBusinessID() gin.HandlerFunc {
	return func(c *gin.Context) {
		usecases, err := getCampaignUsecases()
		if err != nil {
			c.JSON(500, err)
			return
		}
		ID := c.Param("id")
		if ID == "" {
			c.JSON(400, "id is empty")
			return
		}
		uuidID, err := uuid.Parse(ID)
		if err != nil {
			c.JSON(400, fmt.Sprintf("invalid id: %s\n", err))
		}
		businesses, err := usecases.GetCampaignsByBusinessID(uuidID)
		if err != nil {
			// Validate not found
			if err == gorm.ErrRecordNotFound {
				c.JSON(404, "not found")
				return
			}
			c.JSON(500, err)
			return
		}
		c.JSON(200, businesses)

	}
}

func GetAllByBranchID() gin.HandlerFunc {
	return func(c *gin.Context) {
		usecases, err := getCampaignUsecases()
		if err != nil {
			c.JSON(500, err)
			return
		}
		ID := c.Param("id")
		if ID == "" {
			c.JSON(400, "id is empty")
			return
		}
		uuidID, err := uuid.Parse(ID)
		if err != nil {
			c.JSON(400, fmt.Sprintf("invalid id: %s\n", err))
		}
		businesses, err := usecases.GetCampaignsByBranchID(uuidID)
		if err != nil {
			// Validate not found
			if err == gorm.ErrRecordNotFound {
				c.JSON(404, "not found")
				return
			}
			c.JSON(500, err)
			return
		}
		c.JSON(200, businesses)
	}
}

func GetAccumulatedTransactionInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		usecases, err := getCampaignUsecases()
		if err != nil {
			c.JSON(500, err)
			return
		}
		resume, err := usecases.AccumulatePointsAndCashBack()
		if err != nil {
			// Validate not found
			if err == gorm.ErrRecordNotFound {
				c.JSON(404, "not found")
				return
			}
			c.JSON(500, err)
			return
		}
		c.JSON(200, resume)
	}
}
