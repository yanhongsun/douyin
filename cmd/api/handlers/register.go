package handlers

import (
	"context"
	"github.com/douyin/cmd/api/rpc"
	"github.com/douyin/kitex_gen/user"
	"github.com/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerVar RequestParam

	if err := c.ShouldBind(&registerVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), -1, "")
		return
	}

	if len(registerVar.Username) == 0 || len(registerVar.Password) == 0 {
		SendResponse(c, errno.ParamErr, -1, "")
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