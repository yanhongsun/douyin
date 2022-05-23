namespace go video

struct BaseResp {
    1:required  i32 status_code
    2:optional  string status_msg
}

struct User{
    1:required  i64     id
    2:required  string  name
    3:optional  i64     follow_count
    4:optional  i64     follower_count
    5:required  bool    is_follow
}
struct Video{
    1:required  i64     id
    2:required  User    author
    3:required  string  play_url
    4:required  string  cover_url
    5:required  i64     favorite_count
    6:required  i64     comment_count
    7:required  bool    is_favorite
    8:required  string  title
}

struct PublishVideoRequest{
    1:required  string  token
    2:required  byte    data
    3:required  string  title
}
struct PublishVideoResponse{
    1:BaseResp base_resp
}
struct GetPublishListRequest{
    1:required  i64     user_id
    2:required  string  token
}
struct GetPublishListResponse{
    1:BaseResp      base_resp
    2:list<Video>   video_list
}
struct GetFeedRequest{
    1:required  i64     user_id
    2:required  string  token
}
struct GetFeedResponse{
    1:BaseResp      base_resp
    2:list<Video>   video_list
}

service PublishService{
    PublishActionResponse   PublishAction(1:PublishActionRequest    req)
    GetPublishListResponse  GetPublishList(1:GetPublishListRequest  req)
    GetFeedResponse  GetPublishList(1:GetFeedRequest  req)
}