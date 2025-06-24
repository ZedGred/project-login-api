package converter

import (
	"project/internal/entity"
	"project/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		CreateAt: user.CreatedAt,
	}
}

func UserTokenToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID: user.ID,
		
	}
}
