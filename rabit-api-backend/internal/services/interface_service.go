package services

import (
	"github.com/gin-gonic/gin"
	"rabit-api-backend/internal/models"
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
