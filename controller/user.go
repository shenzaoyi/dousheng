package controller

import (
	"dousheng/db"
	response2 "dousheng/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
)

func Register(c *gin.Context) {
	//接收用户名和密码
	name := c.Query("username")
	password := c.Query("password")
	//检索用户合法性：用户名

	//对用户做存储
	u := db.User{
		Name:             name,
		Follow_count:     0,
		Follower_count:   0,
		Is_follow:        false,
		Avatar:           "",
		Background_image: "",
		Signature:        "",
		Total_favorited:  0,
		Work_count:       0,
		Favorite_count:   0,
	}
	err := db.AddUser(&u)
	if err != "" {
		log.Println(err)
	}
	//计算token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   u.ID,
		"name": u.Name,
	})
	secret := []byte("this is a secret")
	token.SignedString(secret)
	//返回响应
	r := response2.RegisterRes{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		UserId:     2,
		Token:      token.Raw,
	}
	log.Println(name + password)
	c.JSON(200, r)
}

func Login(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	token := c.Query("token")

	response := response2.LoginRes{
		StatusCode: 0,
		StatusMsg:  "完事",
		UserId:     2,
		Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.\neyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.\ncThIIoDvwdueQB468K5xDc5633seEFoqwxjF_xSJyQQ",
	}
	c.JSON(200, response)
}

func User(c *gin.Context) {
	//userId := c.Query("user_id")
	//token := c.Query("token")
	//log.Println("-----------------")
	//log.Println(userId)
	//log.Println("-----------------")
	//d := db.Init()
	//user := db.User{}
	//d.First(&user, userId)
	response := response2.Userinfo{
		StatusCode: 0,
		StatusMsg:  "获取用户信息成功",
		User: response2.Userinformatian{
			Id:            2,
			Name:          "shenzaoyi",
			FollowCount:   250,
			FollowerCount: 250,
			IsFollow:      false,
			//Avatar:          "http://i2.sanwen.net/doc/1609/841-160Z5112434-50.jpg",
			//BackgroundImage: "http://i2.sanwen.net/doc/1609/841-160Z5112434-50.jpg",
			Avatar:          "",
			BackgroundImage: "",
			Signature:       "",
			TotalFavorited:  1,
			WorkCount:       0,
			FavoriteCount:   0,
		},
	}
	c.JSON(200, response)
}
