namespace go base

struct BaseResp{
    1: i32 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
}

struct Video {
    1: i64 id // 视频唯一标识
    2: User author // 视频作者信息
    3: string play_url // 视频播放地址
    4: string cover_url // 视频封面地址
    5: i64 favorite_count // 视频的点赞总数
    6: i64 comment_count // 视频的评论总数
    7: bool is_favorite // true-已点赞，false-未点赞
    8: string title // 视频标题
}

struct User{
    1: i64 id // 用户id
    2: string name // 用户名称
    3: i64 follow_count // 关注总数
    4: i64 follower_count // 粉丝总数
    5: bool is_follow // true-已关注，false-未关注
    6: string avatar //用户头像
    7: string background_image //用户个人页顶部大图
    8: string signature //个人简介
    9: i64 total_favorited //获赞数量
    10: i64 work_count //作品数量
    11: i64 favorite_count //点赞数量
}

struct SocialInfo{
    1: i64 follow_count // 关注总数
    2: i64 follower_count // 粉丝总数
    3: bool is_follow // true-已关注，false-未关注
}

struct UserInteractInfo{
    1: i64 total_favorited,     //获赞数量
    2: i64 work_count,         //作品数量
    3: i64 favorite_count,     //点赞数量
}

struct VideoInteractInfo{
    1: i64 favorite_count // 视频的点赞总数
    2: i64 comment_count // 视频的评论总数
    3: bool is_favorite // true-已点赞，false-未点赞
}

struct Comment {
    1: i64 id // 视频评论id
    2: User user // 评论用户信息
    3: string content // 评论内容
    4: string create_date // 评论发布日期，格式 mm-dd
}

struct FriendUser{
    1: i64 id // 用户id
    2: string name // 用户名称
    3: i64 follow_count // 关注总数
    4: i64 follower_count // 粉丝总数
    5: bool is_follow // true-已关注，false-未关注
    6: string avatar //用户头像
    7: string background_image //用户个人页顶部大图
    8: string signature //个人简介
    9: i64 total_favorited //获赞数量
    10: i64 work_count //作品数量
    11: i64 favorite_count //点赞数量
    12: string message // 和该好友的最新聊天消息
    13: i64 msgType // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

struct Message {
    1: i64 id // Message id
    2: i64 to_user_id // The id of the recipient of the message
    3: i64 from_user_id // The id of the sender of the message
    4: string content // Message content
    5: i64 create_time // Message creation time
}

struct LatestMsg{
    1: string message // 和该好友的最新聊天消息
    2: i64 msgType // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}