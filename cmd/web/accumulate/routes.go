package accumulate

import (
	handlers "github.com/braejan/go-leal-challenge/internal/web/handlers/campaigns"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	business := r.Group("/accumulate")
	{
		business.GET("/", handlers.GetAccumulatedTransactionInfo())
	}
}
