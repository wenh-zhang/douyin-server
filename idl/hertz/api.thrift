namespace go api
include "../base.thrift"

struct douyin_feed_request {
    1: i64 latest_time // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: string token // 可选参数，登录用户设置
}

struct douyin_feed_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<base.Video> video_list // 视频列表
    4: i64 next_time // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct douyin_user_register_request {
    1: string username(api.query="username", api.vd="len($)>0 && len($)<33") // 注册用户名，最长32个字符
    2: string password(api.query="password", api.vd="len($)>0 && len($)<33") // 密码，最长32个字符
}

struct douyin_user_register_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: i64 user_id // 用户id
    4: string token // 用户鉴权token
}

struct douyin_user_login_request {
    1: string username(api.query="username", api.vd="len($)>0 && len($)<33") // 登录用户名
    2: string password(api.query="password", api.vd="len($)>0 && len($)<33") // 登录密码
}

struct douyin_user_login_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: i64 user_id // 用户id
    4: string token // 用户鉴权token
}

struct douyin_user_request {
    1: i64 user_id(api.query="user_id") // 用户id
    2: string token(api.query="token") // 用户鉴权token
}

struct douyin_user_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: base.User user // 用户信息
}

struct douyin_publish_action_request {
    1: string token(api.form="token") // 用户鉴权token
    2: string title(api.form="title", api.vd="len($)>0 && len($)<121") // 视频标题
}

struct douyin_publish_action_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
}

struct douyin_publish_list_request {
    1: i64 user_id(api.query="user_id") // 用户id
    2: string token(api.query="token") // 用户鉴权token
}

struct douyin_publish_list_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<base.Video> video_list // 用户发布的视频列表
}

struct douyin_favorite_action_request {
    1: string token(api.query="token") // 用户鉴权token
    2: i64 video_id(api.query="video_id") // 视频id
    3: i8 action_type(api.query="action_type", api.vd="$==1 || $==2") // 1-点赞，2-取消点赞
}

struct douyin_favorite_action_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
}

struct douyin_favorite_list_request {
    1: i64 user_id(api.query="user_id") // 用户id
    2: string token(api.query="token") // 用户鉴权token
}

struct douyin_favorite_list_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<base.Video> video_list // 用户点赞视频列表
}

struct douyin_comment_action_request {
    1: string token(api.query="token") // 用户鉴权token
    2: i64 video_id(api.query="video_id") // 视频id
    3: i8 action_type(api.query="action_type", api.vd="$==1 || $==2") // 1-发布评论，2-删除评论
    4: string comment_text(api.query="comment_text") // 用户填写的评论内容，在action_type=1的时候使用
    5: i64 comment_id(api.query="comment_id") // 要删除的评论id，在action_type=2的时候使用
}

struct douyin_comment_action_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: base.Comment comment // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct douyin_comment_list_request {
    1: string token(api.query="token") // 用户鉴权token
    2: i64 video_id(api.query="video_id") // 视频id
}

struct douyin_comment_list_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<base.Comment> comment_list // 评论列表
}

struct douyin_relation_action_request {
    1: string token(api.query="token") // 用户鉴权token
    2: i64 to_user_id(api.query="to_user_id") // 对方用户id
    3: i8 action_type(api.query="action_type", api.vd="$==1 || $==2") // 1-关注，2-取消关注
}

struct douyin_relation_action_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
}

struct douyin_relation_follow_list_request {
    1: i64 user_id(api.query="user_id") // 用户id
    2: string token(api.query="token") // 用户鉴权token
}

struct douyin_relation_follow_list_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<base.User> user_list // 用户信息列表
}

struct douyin_relation_follower_list_request {
    1: i64 user_id(api.query="user_id") // 用户id
    2: string token(api.query="token") // 用户鉴权token
}

struct douyin_relation_follower_list_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<base.User> user_list // 用户列表
}

struct douyin_relation_friend_list_request {
    1: i64 user_id(api.query="user_id") // 用户id
    2: string token(api.query="token") // 用户鉴权token
}

struct douyin_relation_friend_list_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<base.FriendUser> user_list,     // 用户列表
}

struct douyin_message_chat_request {
    1: string token(api.query="token") // 用户鉴权token
    2: i64 to_user_id(api.query="to_user_id") // 对方用户id
    3: i64 pre_msg_time(api.query="pre_msg_time")//上次最新消息的时间（新增字段-apk更新中）
}

struct douyin_message_chat_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<base.Message> message_list // 消息列表
}

struct douyin_message_action_request {
    1: string token(api.query="token") // 用户鉴权token
    2: i64 to_user_id(api.query="to_user_id") // 对方用户id
    3: i8 action_type(api.query="action_type", api.vd="$==1") // 1-发送消息
    4: string content(api.query="content", api.vd="len($)>0 && len($)<256") // 消息内容
}

struct douyin_message_action_response {
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
}

service ApiService {
    douyin_user_register_response UserRegister(1: douyin_user_register_request req)(api.post="/douyin/user/register/");
    douyin_user_login_response UserLogin(1: douyin_user_login_request req)(api.post="/douyin/user/login/");
    douyin_user_response UserInfo(1: douyin_user_request req)(api.get="/douyin/user/");

    douyin_feed_response Feed (1: douyin_feed_request req)(api.get="/douyin/feed/");
    douyin_publish_action_response PublishAction (1: douyin_publish_action_request req)(api.post="/douyin/publish/action/");
    douyin_publish_list_response PublishList (1: douyin_publish_list_request req)(api.get="/douyin/publish/list/");

    douyin_favorite_action_response FavoriteAction(1: douyin_favorite_action_request req)(api.post="/douyin/favorite/action/");
    douyin_favorite_list_response FavoriteList(1: douyin_favorite_list_request req)(api.get="/douyin/favorite/list/");
    douyin_comment_action_response CommentAction(1: douyin_comment_action_request req)(api.post="/douyin/comment/action/");
    douyin_comment_list_response CommentList(1: douyin_comment_list_request req)(api.get="/douyin/comment/list/");

    douyin_relation_action_response RelationAction(1: douyin_relation_action_request req)(api.post="/douyin/relation/action/");
    douyin_relation_follow_list_response FollowList(1: douyin_relation_follow_list_request req)(api.get="/douyin/relation/follow/list/");
    douyin_relation_follower_list_response FollowerList(1: douyin_relation_follower_list_request req)(api.get="/douyin/relation/follower/list/");
    douyin_relation_friend_list_response FriendList(1: douyin_relation_friend_list_request req)(api.get="/douyin/relation/friend/list/");

    douyin_message_chat_response MessageChat(1: douyin_message_chat_request req)(api.get="/douyin/message/chat/");
    douyin_message_action_response MessageAction(1: douyin_message_action_request req)(api.post="/douyin/message/action/");
}