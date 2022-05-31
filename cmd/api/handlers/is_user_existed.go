package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

type IsUserExistedParam struct {
	UserID int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

func IsUserExisted(c *gin.Context) {
	var existVar IsUserExistedParam

	if err := c.BindQuery(&existVar); err != nil {
		SendIsUserExistedResponse(c, errno.ConvertErr(err), false)
		return
	}

	isExisted, err := rpc.IsUserExisted(context.Background(), &user.DouyinUserExistRequest{
		TargetId: existVar.UserID,
	})

	if err != nil {
		SendIsUserExistedResponse(c, errno.ConvertErr(err), false)
		return
	}

	SendIsUserExistedResponse(c, errno.Success, isExisted)
}
