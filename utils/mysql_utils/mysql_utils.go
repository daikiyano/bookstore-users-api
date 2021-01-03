package mysql_utils

import (
	"bookstore-users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"


)

func ParseError(err error) *errors.RestErr {
	sqlErr,ok := err.(*mysql.MySQLError)
	println("OK",ok,sqlErr)
	if !ok {
		if strings.Contains(err.Error(),errorNoRows) {
			println("errorNoRows")
			println(errors.NewNotFoundError("no recode matching given id"))
			return errors.NewNotFoundError("no recode matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}