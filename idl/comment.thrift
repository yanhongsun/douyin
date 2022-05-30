namespace go comment

struct BaseResp {
    1:i32 status_code
    2:string status_message
}

struct User {
    1:required i64 id
    2:required string name
    3:optional i64 follow_count
    4:optional i64 follower_count
    5:required bool is_follow
}

struct Comment {
    1:i64 comment_id
    2:User user
    3:string content
    4:string create_date
}

struct CreateCommentRequest {
    1:i64 user_id
    2:i64 video_id
    3:string content
    4:string token
}

struct CreateCommentResponse {
    1:BaseResp base_resp
    2:Comment comment
}

struct DeleteCommentRequest {
    1:i64 user_id
    2:i64 video_id
    3:i64 comment_id
    4:string token
}

struct DeleteCommentResponse {
    1:BaseResp base_resp
}

struct QueryCommentsRequest {
    1:i64 video_id
    2:optional string token
}

struct QueryCommentsResponse {
    1:BaseResp base_resp
    2:list<Comment> comments
}

struct QueryCommentNumberRequest {
    1:i64 video_id
}

struct QueryCommentNumberResponse {
    1:BaseResp base_resp
    2:i64 commentNumber
}

service CommentService {
    CreateCommentResponse CreateComment(1:CreateCommentRequest req)
    DeleteCommentResponse DeleteComment(1:DeleteCommentRequest req)
    QueryCommentsResponse QueryComments(1:QueryCommentsRequest req)
    QueryCommentNumberResponse QueryCommentNumber(1:QueryCommentNumberRequest req)
}