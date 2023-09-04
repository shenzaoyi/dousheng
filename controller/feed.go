package controller

import (
	"github.com/gin-gonic/gin"
)

type Author struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Follow_count     int    `json:"follow_count"`
	Follower_count   int    `json:"follower_count"`
	Is_follow        bool   `json:"is_follow"`
	Avatar           string `json:"avatar"`
	Background_image string `json:"background_image"`
	Signature        string `json:"signature"`
	Total_favorited  int    `json:"total_favorited"`
	Work_count       int    `json:"work_count"`
	Favorite_count   int    `json:"favorite_count"`
}

type Vedio struct {
	Id             int    `json:"id"`
	Author         Author `json:"arthur"`
	Play_url       string `json:"play_url"`
	Cover_url      string `json:"cover_url"`
	Favorite_count int    `json:"favorite_count"`
	Comment_count  int    `json:"comment_count"`
	Is_favorite    bool   `json:"is_favorite"`
	Title          string `json:"title"`
}

// Response； 包含Vedio_list, Vedio_list包含Arthur
type Response struct {
	Status_code int     `json:"status_code"`
	Status_msg  string  `json:"status_msg"`
	Next_time   int     `json:"next_time"`
	Vedio_list  []Vedio `json:"vedio_list"`
}

var m Vedio = Vedio{
	Id: 1,
	Author: Author{
		Id:               1,
		Name:             "shenzaoyi",
		Follow_count:     1,
		Follower_count:   1,
		Is_follow:        true,
		Avatar:           "z",
		Background_image: "z",
		Signature:        "z",
		Total_favorited:  1,
		Work_count:       1,
		Favorite_count:   1,
	},
	Play_url:       "http://192.168.43.28:8080/static/mmexport1675696231427.mp4",
	Cover_url:      "http://i2.sanwen.net/doc/1609/841-160Z5112434-50.jpg",
	Favorite_count: 1,
	Comment_count:  1,
	Is_favorite:    true,
	Title:          "test",
}

var vlist [3]Vedio

func Feed(c *gin.Context) {
	vlist[0] = m
	response := Response{
		Status_code: 200,
		Status_msg:  "刷新成功",
		Next_time:   2023,
		Vedio_list:  vlist[0:1],
	}
	c.JSON(200, response)
}
