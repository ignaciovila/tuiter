package service

import (
	"github.com/ignaciovila/tuiter/src/domain"
)

var userList []*domain.User

func AddUser(user *domain.User) {
	userList = append(userList, user)
}
