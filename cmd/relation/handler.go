package main

import (
	"context"
	"douyin/cmd/relation/pack"
	"douyin/cmd/relation/service"
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"
	"fmt"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

/*
type RelationActionRequest struct {
	UserId     int64  `thrift:"user_id,1,required" json:"user_id"`
	Token      string `thrift:"token,2,required" json:"token"`
	ToUserId   int64  `thrift:"to_user_id,3,required" json:"to_user_id"`
	ActionType int32  `thrift:"action_type,4,required" json:"action_type"`
}
type BaseResponse struct {
	StatusCode int32   `thrift:"status_code,1,required" json:"status_code"`
	StatusMsg  *string `thrift:"status_msg,2" json:"status_msg,omitempty"`
}

type GetFollowListResponse struct {
	BaseResp *BaseResponse `thrift:"base_resp,1" json:"base_resp"`
	UserList []*User       `thrift:"user_list,2" json:"user_list"`
}
*/
// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.RelationActionResponse)

	if req.ActionType == 1 {
		err = service.NewFollowService(ctx).Follow(req)
		if err != nil {
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp, err
		}

	} else if req.ActionType == 2 {
		fmt.Println("取消关注服务 handler.go")
		err = service.NewUnFollowService(ctx).UnFollow(req)
		if err != nil {
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp, err
		}

	} else {
		// 参数错误
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *relation.GetFollowListRequest) (resp *relation.GetFollowListResponse, err error) {
	// TODO: Your code here...

	resp = new(relation.GetFollowListResponse)

	user_list, err := service.NewGetFollowListService(ctx).GetFollowList(req)
	fmt.Println("handler.go")
	fmt.Println(user_list)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.UserList = nil
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	fmt.Println(resp.BaseResp)
	/*
			Id            int64  `thrift:"id,1,required" json:"id"`
		    Name          string `thrift:"name,2,required" json:"name"`
		    FollowCount   *int64 `thrift:"follow_count,3" json:"follow_count,omitempty"`
		    FollowerCount *int64 `thrift:"follower_count,4" json:"follower_count,omitempty"`
		    IsFollow      bool   `thrift:"is_follow,5,required" json:"is_follow"
	*/
	/*
			type UserList struct {
			ID             int64
			Name           string
			Follow_count   int64
			Follower_count int64
			Is_follow      bool
		}
	*/
	fmt.Println(len(user_list))
	//relation.User     	[]db.UserList) as []*
	//resp.UserList = user_list
	for _, v := range user_list {
		var tmp relation.User
		tmp.Id = v.ID
		tmp.Name = v.Name
		tmp.FollowCount = &v.Follow_count
		tmp.FollowerCount = &v.Follower_count
		tmp.IsFollow = v.Is_follow
		resp.UserList = append(resp.UserList, &tmp)
	}
	fmt.Println(resp)
	return resp, nil
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *relation.GetFollowerListRequest) (resp *relation.GetFollowerListResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.GetFollowerListResponse)
	user_list, err := service.NewGetFollowerListService(ctx).GetFollowerList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.UserList = nil
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	for _, v := range user_list {
		var tmp relation.User
		tmp.Id = v.ID
		tmp.Name = v.Name
		tmp.FollowCount = &v.Follow_count
		tmp.FollowerCount = &v.Follower_count
		tmp.IsFollow = v.Is_follow
		resp.UserList = append(resp.UserList, &tmp)
	}
	return resp, nil
}

// IsFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) IsFollow(ctx context.Context, req *relation.IsFollowRequest) (resp *relation.IsFollowResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.IsFollowResponse)
	tag, err := service.NewIsFollowService(ctx).IsFollow(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.IsFollow = false
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.IsFollow = tag
	return resp, nil
}
