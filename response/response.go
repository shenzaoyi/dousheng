package response

type Userinfo struct {
	StatusCode int             `json:"status_code"`
	StatusMsg  string          `json:"status_msg"`
	User       Userinformatian `json:"user"`
	//User       struct {
	//	Id              int    `json:"id"`
	//	Name            string `json:"name"`
	//	FollowCount     int    `json:"follow_count"`
	//	FollowerCount   int    `json:"follower_count"`
	//	IsFollow        bool   `json:"is_follow"`
	//	Avatar          string `json:"avatar"`
	//	BackgroundImage string `json:"background_image"`
	//	Signature       string `json:"signature"`
	//	TotalFavorited  string `json:"total_favorited"`
	//	WorkCount       int    `json:"work_count"`
	//	FavoriteCount   int    `json:"favorite_count"`
	//} `json:"user"`
}

type Userinformatian struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	FollowCount     int    `json:"follow_count"`
	FollowerCount   int    `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int    `json:"total_favorited"`
	WorkCount       int    `json:"work_count"`
	FavoriteCount   int    `json:"favorite_count"`
}

type LoginRes struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int    `json:"user_id"`
	Token      string `json:"token"`
}

type RegisterRes struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int    `json:"user_id"`
	Token      string `json:"token"`
}
