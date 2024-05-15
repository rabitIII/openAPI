package services

import "rabit-api-backend/internal/models"

type userService struct {
	// 依赖注入模型
	userRepository *models.UserModel
}

func (s *userService) RegisterUser(user *models.UserModel) error {
	userModel := &userService{}
}
