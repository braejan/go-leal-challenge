package businesses

import (
	handlers "github.com/braejan/go-leal-challenge/internal/web/handlers/businesses"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	business := r.Group("/business")
	{
		business.GET("/:id", handlers.Get())
		business.GET("/", handlers.GetAll())
	}
}
