package services

import (
	"rabit-api-backend/models"

	"github.com/gin-gonic/gin"
)

type InterfaceService struct {
	InterfaceRepository *models.InterfaceModel
}

func (s *InterfaceService) GetInterface(ctx *gin.Context) {

}
