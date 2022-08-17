package socket

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"

	"tdp-cloud/internal/dborm/user"
	"tdp-cloud/internal/slaver"
	"tdp-cloud/internal/webssh"
)

func agent(c *gin.Context) {

	at := strings.Replace(c.Param("auth"), "0:", "", 1)

	u, err := user.Fetch(&user.FetchParam{
		AppToken: at,
	})

	if err != nil || u.Id == 0 {
		c.AbortWithError(400, errors.New("授权失败"))
		return
	}

	c.Set("UserId", u.Id)

	slaver.Upgrader(c)

}

func ssh(c *gin.Context) {

	webssh.Handle(c)

}