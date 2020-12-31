package users

import (
	"bookstore-users-api/datasources/mysql/users_db"
	"bookstore-users-api/utils/date_utils"
	"bookstore-users-api/utils/errors"
	"fmt"
)

var (
	usersDB = make(map[int64]*User)
)
func something() {
	user := User{}

	if err := user.Get(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.FirstName)


}

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user nod fond",user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated  = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registerd",user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user already exists",user.Id))
	}
	user.DateCreated = date_utils.GetNowString()

	usersDB[user.Id] = user
	return nil
}