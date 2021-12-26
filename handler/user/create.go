package user

import (
	"fmt"
	"go-restful-api-server-simple-practice/pkg/errno"
	"net/http"

	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var err error
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": errno.BindError.Code, "message": errno.BindError.Message})
		return
	}
	log.Debugf("username is : [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		err = errno.New(errno.UserNotFoundError, fmt.Errorf("username can not found in db")).Add("this is add messgae")
		log.Errorf(err, "Get an error")
	}
	if errno.IsErrUSerNotFound(err) {
		log.Debug("err type is ErrUSerNotFound")
	}
	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
