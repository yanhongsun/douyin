package handlers

import (
	"context"
	"github.com/douyin/cmd/api/rpc"
	"github.com/douyin/kitex_gen/user"
	"github.com/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerVar UserParam
	// 参数绑定
	if err := c.ShouldBind(&registerVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// 用户名或密码不能为空
	if len(registerVar.Username) == 0 || len(registerVar.Password) == 0 {
		SendResponse(c, errno.ParamErr, nil)
	}
	// 开始创建用户
	userID, token, err := rpc.CreateUser(context.Background(), &user.DouyinUserRegisterRequest{
		Username: registerVar.Username,
		Password: registerVar.Password,
	})
	// 创建用户失败
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// 创建用户成功, 将信息返回
	SendResponse(c, errno.Success, map[string]interface{}{"user_id": userID, "token": token})
}
