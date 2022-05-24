package pack

import (
	"github.com/douyin/kitex_gen/user"
	"github.com/douyin/pkg/errno"
)

func BuildRegisterResp(err error) *user.DouyinUserRegisterResponse {
	if err == nil {
		return &user.DouyinUserRegisterResponse{
			StatusCode: errno.Success.ErrCode,
			StatusMsg: &errno.Success.ErrMsg,
			UserId:
			Token: "",
		}
	}
}

func BuildLoginResp(err error) *user.DouyinUserLoginResponse {

}

func BuildUserResponse(err error) *user.DouyinUserResponse {

}