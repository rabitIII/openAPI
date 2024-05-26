package services

import (
	"rabit-api-backend/models"

	"github.com/gin-gonic/gin"
)

type InterfaceService struct {
	InterfaceRepository *models.InterfaceModel
}

func (s *InterfaceService) GetInterface(ctx *gin.Context) {
	// TODO 业务逻辑处理
	err := s.InterfaceRepository.ReadList()
	if err != nil {
		return
	}
}
