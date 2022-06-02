namespace go user

struct douyin_user_register_request {
    1:required string username
    2:required string password
}

struct douyin_user_register_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required i64 user_id
    4:required string token
}

struct douyin_user_login_request {
    1:required string username
    2:required string password
}

struct douyin_user_login_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required i64 user_id
    4:required string token
}

struct douyin_user_request {
    1:required i64 user_id
    2:required string token
}

struct douyin_user_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required User user
}

struct User {
    1:required i64 id
    2:required string name
    3:optional i64 follow_count
    4:optional i64 follower_count
    5:required bool is_follow
}

struct douyin_user_exist_request {
    1:required i64 target_id
    # 2:required string token
}

struct douyin_user_exist_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required bool is_existed
}

struct douyin_query_user_request {
    1:required i64 user_id
    2:required i64 target_id
    3:required string token
}

struct douyin_mquery_user_request {
    1:required i64 user_id
    2:required list<i64> target_ids
    3:required string token
}

struct douyin_mquery_user_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<User> users
}

service UserService {
    douyin_user_register_response CreateUser(1:douyin_user_register_request req)
    douyin_user_login_response CheckUser(1:douyin_user_login_request req)
    douyin_user_response QueryCurUser(1:douyin_user_request req)
    douyin_user_response QueryOtherUser(1:douyin_query_user_request req)
    douyin_user_exist_response IsUserExisted(1:douyin_user_exist_request req)
    douyin_mquery_user_response MultiQueryUser(1:douyin_mquery_user_request req)
}
