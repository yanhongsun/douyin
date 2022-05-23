package pack

import "github.com/douyin/kitex_gen/douyin_user"

func BuildUserRegisterResp(err error) *douyin_user.DouyinUserRegisterResponse {
	if err == nil {
		return &douyin_user.DouyinUserRegisterResponse{
			StatusCode: 0,		// TODO: errno
			StatusMsg: "ok",		// TODO: errno
			UserId:
		}
	}
	return nil
}

func BuildUserLoginResp(err error) *douyin_user.DouyinUserLoginResponse {
	return nil
}

func BuildUserResp(err error) *douyin_user.DouyinUserResponse {
	return nil
}