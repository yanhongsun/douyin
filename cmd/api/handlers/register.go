package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerVar RequestParam
	if err := c.ShouldBind(&registerVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), -1, "")
		return
	}
	fmt.Println("registerVar:", registerVar)
	if len(registerVar.Username) == 0 || len(registerVar.Password) == 0 {
		SendResponse(c, errno.ParamErr, -1, "")
		return
	}

	userID, token, err := rpc.CreateUser(context.Background(), &user.DouyinUserRegisterRequest{
		Username: registerVar.Username,
		Password: registerVar.Password,
	})

	if err != nil {
		SendResponse(c, errno.ConvertErr(err), -1, "")
		return
	}

	SendResponse(c, errno.Success, userID, token)
}
