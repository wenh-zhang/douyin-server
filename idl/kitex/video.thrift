namespace go video
include"../base.thrift"


struct douyin_feed_request {
      1:i64 latest_time // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
      2: i64 user_id // 用户id
}

struct douyin_feed_response {
    1: base.BaseResp base_resp
    2:list<base.Video> video_list  // 视频列表
    3:i64 next_time  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct douyin_publish_action_request {
    1: i64 user_id // 用户id
    2: string play_url //视频url
    3: string cover_url //封面url
    4:string title  // 视频标题
}

struct douyin_publish_action_response {
    1: base.BaseResp base_resp
}

struct douyin_get_publish_list_request {
    1: i64 local_user_id  // 本地用户id
    2: i64 target_user_id //目标用户id
}

struct douyin_get_publish_list_response {
    1: base.BaseResp base_resp
    2: list<base.Video> video_list  // 用户发布的视频列表
}

struct douyin_get_favorite_list_request {
    1: i64 local_user_id  // 本地用户id
    2: i64 target_user_id //目标用户id
}

struct douyin_get_favorite_list_response {
    1: base.BaseResp base_resp
    2: list<base.Video> video_list  // 用户点赞视频列表
}

struct douyin_get_publish_video_id_list_request{
    1: i64 user_id;
}

struct douyin_get_publish_video_id_list_response{
    1: base.BaseResp base_resp
    2: list<i64> video_id_list
}

service VideoService {
    douyin_feed_response Feed(1: douyin_feed_request req),
    douyin_publish_action_response Publish(1: douyin_publish_action_request req),
    douyin_get_publish_list_response GetPublishedVideoList(1: douyin_get_publish_list_request req),
    douyin_get_favorite_list_response GetFavoriteVideoList(1: douyin_get_favorite_list_request req),
    douyin_get_publish_video_id_list_response GetPublishVideoIdList(1: douyin_get_publish_video_id_list_request req),
}