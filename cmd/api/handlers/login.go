package handlers

import (
	"context"
	"github.com/douyin/cmd/api/rpc"
	"github.com/douyin/kitex_gen/user"
	"github.com/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginVar UserParam
	// 参数绑定
	if err := c.ShouldBind(&loginVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// 用户名或密码不能为空
	if len(loginVar.Username) == 0 || len(loginVar.Password) == 0 {
		SendResponse(c, errno.ParamErr, nil)
	}
	// 远程过程调用 - 登录
	userID, token, err := rpc.CheckUser(context.Background(), &user.DouyinUserLoginRequest{
		Username: loginVar.Username,
		Password: loginVar.Password,
	})

	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(c, errno.Success, map[string]interface{}{"user_id": userID, "token": token})
}
