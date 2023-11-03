package branches

import (
	"log"

	"github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	"github.com/gin-gonic/gin"
)

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		usecases, err := getBranchUsecases()
		if err != nil {
			c.JSON(500, err)
			return
		}
		var branch model.Branch
		err = c.ShouldBind(&branch)
		if err != nil {
			c.JSON(400, err)
			return
		}
		err = usecases.CreateBranch(branch)
		if err != nil {
			log.Println("Error Create: ", err)
			c.JSON(503, err)
			return
		}
		c.Status(201)
	}
}
