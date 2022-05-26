namespace go comment

struct BaseResp {
    1:i32 status_code
    2:string status_message
}

struct Comment {
    1:i64 comment_id
    2:i64 user_id
    3:string content
    4:string create_date
}

struct CreateCommentRequest {
    1:i64 user_id
    2:i64 video_id
    3:string content
}

struct CreateCommentResponse {
    1:BaseResp base_resp
    2:Comment comment
}

struct DeleteCommentRequest {
    1:i64 user_id
    2:i64 video_id
    3:i64 comment_id
}

struct DeleteCommentResponse {
    1:BaseResp base_resp
}

struct QueryCommentsRequest {
    1:i64 video_id
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