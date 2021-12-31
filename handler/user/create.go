package user

import (
	"fmt"
	"go-restful-api-server-simple-practice/handler"
	"go-restful-api-server-simple-practice/pkg/errno"

	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Username string `json:"username"`
}

// @Summary 创建用户
// @Description 创建一个新的用户
// @Accept x-www-form-urlencoded
// @Product json
// @Param username body CreateUserRequest true "创建用户"
// @Tags user
// @Router /v1/user/:username [post]
// @Success 200 {string} json "{"code":0, "data": {}, "message": ""}"
func Create(c *gin.Context) {
	var r CreateUserRequest
	var err error
	if err := c.Bind(&r); err != nil {
		log.Errorf(err, "Bind an error")
		handler.SendResponse(c, errno.BindError, nil)
		return
	}
	log.Debugf("username is : [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		err = errno.New(errno.UserNotFoundError, fmt.Errorf("username can not found in db")).Add("this is add messgae")
		log.Errorf(err, "Get an error")
	}
	if errno.IsErrUSerNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}
	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}
	handler.SendResponse(c, err, CreateUserResponse{Username: r.Username})
	// code, message := errno.DecodeErr(err)
	// c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
