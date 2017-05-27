package util

import (
	"github.com/maprost/playground/userserver/client"
	"github.com/maprost/playground/userserver/core/data"
)

func ConvertUserToUserDto(user data.UserAccount) client.UserDto {
	return client.UserDto{
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}

func ConvertUsersToUserDtos(users []data.UserAccount) []client.UserDto {
	userDtos := make([]client.UserDto, len(users))
	for i, user := range users {
		userDtos[i] = ConvertUserToUserDto(user)
	}

	return userDtos
}
