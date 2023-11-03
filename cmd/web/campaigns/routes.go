package campaigns

import (
	handlers "github.com/braejan/go-leal-challenge/internal/web/handlers/campaigns"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	business := r.Group("/campaign")
	{
		business.GET("/:id", handlers.Get())
		business.GET("/business/:id", handlers.GetAllByBusinessID())
		business.GET("/branch/:id", handlers.GetAllByBranchID())
		business.POST("/", handlers.Create())
		// business.PUT("/:id", handlers.Update())
		// business.DELETE("/:id", handlers.Delete())
	}
}
