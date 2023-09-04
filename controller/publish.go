package controller

import (
	"dousheng/global"
	"github.com/gin-gonic/gin"
	"log"
)

func Upload(c *gin.Context) {
	//title := c.PostForm("title")
	f, err := c.FormFile("data")
	if err != nil {
		log.Println(err)
		log.Println("读取文件失败")
		return
	}
	dst := "D:/CODES/go/dousheng/static" + "/" + f.Filename
	err = c.SaveUploadedFile(f, dst)
	if err != nil {
		log.Println(err)
		log.Println("保存文件出错")
	}
	r := global.Res{
		StatusCode: 200,
		StatusMsg:  "投稿上传成功",
	}
	c.JSON(200, r)

}
