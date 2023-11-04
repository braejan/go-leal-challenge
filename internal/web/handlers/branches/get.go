package branches

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		usecases, err := getBranchUsecases()
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
		branch, err := usecases.GetBranchByID(uuidID)
		if err != nil {
			// Validate not found
			if err == gorm.ErrRecordNotFound {
				c.JSON(404, "not found")
				return
			}
			c.JSON(500, err)
			return
		}
		c.JSON(200, branch)

	}
}

func GetByBusinessID() gin.HandlerFunc {
	return func(c *gin.Context) {
		usecases, err := getBranchUsecases()
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
		branches, err := usecases.GetBranchesByBusinessID(uuidID)
		if err != nil {
			// Validate not found
			if err == gorm.ErrRecordNotFound {
				c.JSON(404, "not found")
				return
			}
			c.JSON(500, err)
			return
		}
		c.JSON(200, branches)

	}
}
