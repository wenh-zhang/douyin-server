namespace go interaction
include "../base.thrift"

struct douyin_favorite_action_request {
    1: i64 user_id // 用户id
    2: i64 video_id// 视频id
    3: i8 action_type// 1-点赞，2-取消点赞
}

struct douyin_favorite_action_response {
    1: base.BaseResp base_resp
}

struct douyin_get_favorite_video_id_list_request {
    1: i64 user_id// 要查询的用户id
}

struct douyin_get_favorite_video_id_list_response {
    1: base.BaseResp base_resp
    2: list<i64> video_id_list // 用户点赞视频id列表
}

struct douyin_comment_action_request {
    1: i64 user_id //用户id
    2: i64 video_id// 视频id
    3: i8 action_type// 1-发布评论，2-删除评论
    4: string comment_text// 用户填写的评论内容，在action_type=1的时候使用
    5: i64 comment_id // 要删除的评论id，在action_type=2的时候使用
}

struct douyin_comment_action_response {
    1: base.BaseResp base_resp
    2: base.Comment comment // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct douyin_get_comment_list_request {
    1: i64 video_id // 视频id
    2: i64 user_id // 用户id
}

struct douyin_get_comment_list_response {
    1: base.BaseResp base_resp
    2: list<base.Comment> comment_list // 评论列表
}

struct douyin_batch_get_video_interact_info_request{
    1: i64 user_id //用户id
    2: list<i64> video_id_list // 视频id列表
}

struct douyin_batch_get_video_interact_info_response{
    1: base.BaseResp base_resp
    2: list<base.VideoInteractInfo> video_interact_info_list
}

struct douyin_batch_get_user_interact_info_request{
    1: list<i64> target_user_id_list // 目标查询用户id列表
}

struct douyin_batch_get_user_interact_info_response{
    1: base.BaseResp base_resp
    2: list<base.UserInteractInfo> user_interact_info_list
}

service InteractionServer {
    douyin_favorite_action_response Favorite(1: douyin_favorite_action_request req),
    douyin_get_favorite_video_id_list_response GetFavoriteVideoIdList(1: douyin_get_favorite_video_id_list_request req),
    douyin_comment_action_response Comment(1: douyin_comment_action_request req),
    douyin_get_comment_list_response GetCommentList(1: douyin_get_comment_list_request req),
    douyin_batch_get_video_interact_info_response BatchGetVideoInteractInfo (1: douyin_batch_get_video_interact_info_request req),
    douyin_batch_get_user_interact_info_response BatchGetUserInteractInfo (1: douyin_batch_get_user_interact_info_request req),
}