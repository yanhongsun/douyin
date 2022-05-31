package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetFollowerList(c *gin.Context) {
	fmt.Println("sun yan hong")
	var queryVar struct {
		UserId int64  `json:"user_id" form:"user_id"`
		Token  string `json:"token" form:"token"`
	}
	//TODO错误处理
	if err := c.BindQuery(&queryVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	fmt.Println(queryVar.UserId)
	fmt.Println(queryVar.Token)
	//TODO错误处理
	if queryVar.UserId < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	/*


			type GetFollowerListRequest struct {
			UserId int64  `thrift:"user_id,1,required" json:"user_id"`
			Token  string `thrift:"token,2,required" json:"token"`
		}
	*/
	req := relation.GetFollowerListRequest{
		UserId: queryVar.UserId,
		Token:  queryVar.Token,
	}
	user_list, err := rpc.GetFollowerList(context.Background(), &req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// TODO  这里序列化有问题
	SendResponse(c, errno.Success, user_list)
}
