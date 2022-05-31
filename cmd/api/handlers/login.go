package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginVar RequestParam

	if err := c.BindQuery(&loginVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), -1, "")
		return
	}
	if len(loginVar.Username) == 0 || len(loginVar.Password) == 0 {
		SendResponse(c, errno.ParamErr, -1, "")
		return
	}

	userID, token, err := rpc.CheckUser(context.Background(), &user.DouyinUserLoginRequest{
		Username: loginVar.Username,
		Password: loginVar.Password,
	})

	if err != nil {
		SendResponse(c, errno.ConvertErr(err), -1, "")
		return
	}

	SendResponse(c, errno.Success, userID, token)
}
