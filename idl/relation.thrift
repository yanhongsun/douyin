namespace go relation

struct BaseResponse {
    1:required  i32 status_code
    2:optional  string status_msg
}

struct RelationActionRequest{
    1:required  i64  user_id
    2:required  string    token
    3:required  i64  to_user_id
    4:required  i32 action_type
}
struct RelationActionResponse{
    1:BaseResponse base_resp
}

struct User{
    1:required  i64     id
    2:required  string  name
    3:optional  i64     follow_count
    4:optional  i64     follower_count
    5:required  bool    is_follow
}

struct GetFollowListRequest{
    1:required  i64 user_id
    2:required  string token
}

struct GetFollowListResponse{
    1:BaseResponse base_resp
    2:list<User> user_list
}

struct GetFollowerListRequest{
    1:required  i64 user_id
    2:required  string token
}

struct GetFollowerListResponse{
    1:BaseResponse base_resp
    2:list<User> user_list
}

struct IsFollowRequest{
    1:required  i64 user_id
    2:required  string token
    3:required  i64  to_user_id
}

struct IsFollowResponse{
    1:BaseResponse base_resp
    2:required  bool is_follow
}

service RelationService{
    RelationActionResponse   RelationAction(1:RelationActionRequest    req)
    GetFollowListResponse  GetFollowList(1:GetFollowListRequest  req)
    GetFollowerListResponse  GetFollowerList(1:GetFollowerListRequest  req)
    IsFollowResponse    IsFollow(1:IsFollowRequest req)
}


