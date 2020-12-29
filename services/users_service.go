package services

import (
	"bookstore-users-api/domain/users"
)

func CreateUser(user users.User) (*users.User,error){
	return &user,nil
}