package constants

const (
	MySQLDefaultDSN   = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	UserTableName     = "user"
	VideoTableName    = "video"
	FavoriteTableName = "favorite"
	CommentTableName  = "comment"
	FollowTableName   = "follow"
	MessageTableName  = "message"
)
