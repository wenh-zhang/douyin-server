namespace go sociality
include "../base.thrift"
struct douyin_relation_action_request {
    1: i64 local_user_id // 本地用户id
    2: i64 target_user_id // 对方用户id
    3: i8 action_type // 1-关注，2-取消关注
}

struct douyin_relation_action_response {
    1: base.BaseResp base_resp
}

struct douyin_get_relation_user_id_list_request {
    1: i64 user_id // 用户id
    2: i8 option // 0-关注，1-粉丝，2-朋友
}

struct douyin_get_relation_user_id_list_response {
    1: base.BaseResp base_resp
    2: list<i64> user_id_list // 用户id列表
}

struct douyin_batch_get_social_info_request{
    1: i64 local_user_id
    2: list<i64> target_user_id_list
}

struct douyin_batch_get_social_info_response{
    1: base.BaseResp base_resp
    2: list<base.SocialInfo> social_info_list
}

service SocialityService {
    douyin_relation_action_response Relation(1: douyin_relation_action_request req),
    douyin_get_relation_user_id_list_response GetRelationUserIdList(1: douyin_get_relation_user_id_list_request req),
    douyin_batch_get_social_info_response BatchGetSocialInfo(1: douyin_batch_get_social_info_request req),
}