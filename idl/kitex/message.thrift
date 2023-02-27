namespace go message
include"../base.thrift"

struct douyin_get_message_chat_request {
    1: i64 local_user_id //本地用户id
    2: i64 target_user_id // 目标用户id
    3: i64 pre_msg_time //上次最新消息的时间（新增字段-apk更新中）
}

struct douyin_get_message_chat_response {
    1: base.BaseResp base_resp
    2: list<base.Message> message_list // 消息列表
}

struct douyin_message_action_request {
    1: i64 local_user_id //本地用户id
    2: i64 target_user_id // 目标用户id
    3: i8 action_type // 1-发送消息
    4: string content // 消息内容
}

struct douyin_message_action_response {
    1: base.BaseResp base_resp
}

struct douyin_batch_get_latest_message_request{
    1: i64 local_user_id //本地用户id
    2: list<i64> target_user_id_list // 目标用户id列表
}

struct douyin_batch_get_latest_message_response{
    1: base.BaseResp base_resp
    2: list<base.LatestMsg> latestMsgList
}

service MessageService{
    douyin_message_action_response SendMessage(1: douyin_message_action_request req)
    douyin_get_message_chat_response GetMessageHistory(1: douyin_get_message_chat_request req)
    douyin_batch_get_latest_message_response BatchGetLatestMessage(1: douyin_batch_get_latest_message_request req)
}