namespace go user
include "../base.thrift"

struct douyin_user_register_request {
    1: string username// 注册用户名，最长32个字符
    2: string password// 密码，最长32个字符
}

struct douyin_user_register_response {
    1: base.BaseResp base_resp
    2: i64 user_id // 用户id
    3: string token // 用户鉴权token
}

struct douyin_user_login_request {
    1: string username // 登录用户名
    2: string password // 登录密码
}

struct douyin_user_login_response {
    1: base.BaseResp base_resp
    2: i64 user_id // 用户id
    3: string token // 用户鉴权token
}

struct douyin_batch_get_user_request{
    1: i64 local_user_id// 本地用户id，不清楚时设置为0
    2: list<i64> target_user_id_list// 要查询的用户id列表
}
struct douyin_batch_get_user_response {
    1: base.BaseResp base_resp
    2: list<base.User> user_list // 用户信息
}

struct douyin_get_relation_follow_list_request {
    1: i64 local_user_id // 本地用户id
    2: i64 target_user_id // 对方用户id
}

struct douyin_get_relation_follow_list_response {
    1: base.BaseResp base_resp
    2: list<base.User> user_list // 用户信息列表
}

struct douyin_get_relation_follower_list_request {
    1: i64 local_user_id // 本地用户id
    2: i64 target_user_id // 对方用户id
}

struct douyin_get_relation_follower_list_response {
    1: base.BaseResp base_resp
    2: list<base.User> user_list // 用户列表
}

struct douyin_get_relation_friend_list_request {
    1: i64 local_user_id // 本地用户id
    2: i64 target_user_id // 对方用户id
}

struct douyin_get_relation_friend_list_response {
    1: base.BaseResp base_resp
    2: list<base.FriendUser> user_list,     // 用户列表
}

service UserService {
    douyin_user_register_response Register(1: douyin_user_register_request req),
    douyin_user_login_response Login(1: douyin_user_login_request req),
    douyin_batch_get_user_response BatchGetUserInfo(1: douyin_batch_get_user_request req),
    douyin_get_relation_follow_list_response GetRelationFollowList(1: douyin_get_relation_follow_list_request req),
    douyin_get_relation_follower_list_response GetRelationFollowerList(1: douyin_get_relation_follower_list_request req),
    douyin_get_relation_friend_list_response GetRelationFriendList(1: douyin_get_relation_friend_list_request req),
}