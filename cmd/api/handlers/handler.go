package handlers

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/20 17:36
 **/

type UserInfoReq struct {
	UserId int64  `json:"user_id" form:"user_id" binding:"required"`
	Token  string `json:"token" form:"token" binding:"required"`
}

type UserInfoResp struct {
	StatusCode int32     `json:"status_code"`
	StatusMsg  string    `json:"status_msg"`
	User       ProtoUser `json:"user"`
}

type UserRegisterRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserRegisterResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

type UserLoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserLoginResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

type PublishActionReq struct {
	UserId int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
	Data   []byte `json:"data" form:"data"`
	Title  string `json:"title" form:"title"`
}

type PublishActionResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	VideoId    int64  `json:"video_id"`
}

type VideoFeedResp struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	VideoList  []Video `json:"video_list"`
	NextTime   int64   `json:"next_time"`
}

type Video struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	Author        ProtoUser `json:"author"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	IsFavorite    bool      `json:"is_favorite"`
}

type ProtoUser struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type FollowListResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	UserList   []ProtoUser `json:"user_list"`
}

type FollowerListResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	UserList   []ProtoUser `json:"user_list"`
}

type RelationActionRequest struct {
	// // UserId 		int64 	`json:"user_id" form:"user_id"`
	Token      string `json:"token" form:"token"`
	ToUserId   int64  `json:"to_user_id" form:"to_user_id"`
	ActionType int32  `json:"action_type" form:"action_type"`
}

type RelationActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type FavoriteActionRequest struct {
	UserId     int64  `json:"user_id" form:"user_id"`
	Token      string `json:"token" form:"token"`
	VideoId    int64  `json:"video_id" form:"video_id"`
	ActionType int32  `json:"action_type" form:"action_type"`
}

type FavoriteActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type PublishListResp struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	VideoList  []Video `json:"video_list"`
}

type FavoriteListResponse struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	VideoList  []Video `json:"video_list"`
}

type CommentRequest struct {
	UserId      int64  `json:"user_id" form:"user_id"`
	Token       string `json:"token" form:"token"`
	VideoId     int64  `json:"video_id" form:"video_id" binding:"required"`
	ActionType  int32  `json:"action_type" form:"action_type" binding:"required"`
	CommentText string `json:"comment_text" form:"comment_text"`
	CommentId   int64  `json:"comment_id" form:"comment_id"`
}

type CommentResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type CommentListReq struct {
	UserId  int64
	Token   string `form:"token"`
	VideoId int64  `form:"video_id" binding:"required"`
}

type CommentListResp struct {
	StatusCode  int32     `json:"status_code"`
	StatusMsg   string    `json:"status_msg"`
	CommentList []Comment `json:"comment_list"`
}

type Comment struct {
	Id          int64     `json:"id"`
	CommentUser ProtoUser `json:"user"`
	Content     string    `json:"content"`
	CreateDate  string    `json:"create_date"`
}
