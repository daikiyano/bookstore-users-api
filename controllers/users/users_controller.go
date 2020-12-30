package users

import (
	"bookstore-users-api/domain/users"
	"bookstore-users-api/services"
	"bookstore-users-api/utils/errors"

	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	//{0    }
	//リクエストBodyの読み込み
	//json.UnmarshalでUser構造体に格納
	//指定されたバインディングエンジンを使用して、渡された構造体ポインターをバインドする。
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated,result)
	}



func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented,"implement me!")
}

//func SearchUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented,"implement me!")
//}
