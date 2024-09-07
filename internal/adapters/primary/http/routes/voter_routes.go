package routes

import (
	"github.com/axel-andrade/opina-ai-api/internal/infra"
	"github.com/gin-gonic/gin"
)

func configureVoterRoutes(r *gin.RouterGroup, d *infra.Dependencies) {
	voters := r.Group("voters")
	{
		voters.POST("/", d.CreateVoterController.Handle)
	}
}
