package users

import (
	"bookstore-users-api/datasources/mysql/users_db"
	"bookstore-users-api/utils/date_utils"
	"bookstore-users-api/utils/errors"
	"fmt"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
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
	println("Resecdc",user)
	stmt,err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	//必ずcloseする
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()

	insertResult,err :=  stmt.Exec(user.FirstName,user.LastName,user.Email,user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s",err.Error()))
	}
	//result,err := users_db.Client.Exec(queryInsertUser,user.FirstName,user.LastName,user.Email,user.DateCreated)
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s",err.Error()))
	}
	user.Id = userId
	return nil
	//
	//current := usersDB[user.Id]
	//if current != nil {
	//	if current.Email == user.Email {
	//		return errors.NewBadRequestError(fmt.Sprintf("email %s already registerd",user.Email))
	//	}
	//	return errors.NewBadRequestError(fmt.Sprintf("user already exists",user.Id))
	//}
	//user.DateCreated = date_utils.GetNowString()
	//
	//usersDB[user.Id] = user
	//return nil
}