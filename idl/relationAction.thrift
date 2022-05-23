namespace go relation

struct BaseResponse {
    1:required  i32 status_code
    2:optional  string status_msg
}

struct relationActionRequest{
    1:required  i64  user_id
    2:required  string    token
    3:required  i64  to_user_id
    4:required  i32 action_type
}

struct User{
    1:required  i64     id
    2:required  string  name
    3:optional  i64     follow_count
    4:optional  i64     follower_count
    5:required  bool    is_follow
}

struct GetFollwListRequest{
    1:required  i64 user_id
    2:required  string token
}

struct GetFollwListResponse{
    1:BaseResponse base_resp
    2:list<User> user_list
}

struct GetFollwerListRequest{
    1:required  i64 user_id
    2:required  string token
}

struct GetFollwerListResponse{
    1:BaseResponse base_resp
    2:list<User> user_list
}


