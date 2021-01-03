package users

import (
	"bookstore-users-api/domain/users"
	"bookstore-users-api/services"
	"bookstore-users-api/utils/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)

	//{0    }
	//リクエストBodyの読み込み
	//json.UnmarshalでUser構造体に格納
	//指定されたバインディングエンジンを使用して、渡された構造体ポインターをバインドする。
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("heyyyyyyyyyyyyy")
		fmt.Println(err)
		//無効なJSONの場合、エラー返し
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
	// int型に変換し、userIdを取得
	userId,userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id should be anumber")
		c.JSON(err.Status,err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusCreated,user)
}

func UpdateUser(c *gin.Context) {
	userId,userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id should be anumber")
		c.JSON(err.Status,err)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		//無効なJSONの場合、エラー返し
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}

	user.Id = userId

	result,err := services.UpdateUser(user)
	if err != nil {
		c.JSON(err.Status,err)
		return
	}
	c.JSON(http.StatusOK,result)
}

//func SearchUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented,"implement me!")
//}
