package branches

import (
	handlers "github.com/braejan/go-leal-challenge/internal/web/handlers/branches"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	business := r.Group("/branch")
	{
		business.GET("/:id", handlers.Get())
		business.GET("/business/:id", handlers.GetByBusinessID())
	}
}
