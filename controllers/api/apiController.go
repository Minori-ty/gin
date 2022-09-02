package api

import (
	"fmt"
	"gin/models"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (con UserController) Index(c *gin.Context) {

	id := c.DefaultQuery("id", "1")
	c.JSON(200, gin.H{
		"id":  id,
		"msg": "success",
	})
}

type article struct {
	Title string `json:"title"`
	Age   int    `json:"age"`
}

func (con UserController) Article(c *gin.Context) {
	a := article{
		Title: "标题",
		Age:   27,
	}
	c.JSON(200, a)
}

func (con UserController) Upload(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")

	if err == nil {
		// 获取文件后缀名
		extName := path.Ext(file.Filename)

		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}

		// 判断文件是否合法
		if _, ok := allowExtMap[extName]; !ok {
			c.String(404, "不合法")
		}

		// 获取当前时间
		day := models.GetDay()
		dir := "./temp/" + day

		// 创建文件目录
		err := os.MkdirAll(dir, 0666)
		if err != nil {
			fmt.Println("创建失败")
			return
		}
		unix := models.GetUnix()

		// 生成文件名称
		fileName := strconv.FormatInt(unix, 10) + extName
		dst := path.Join(dir, fileName)
		c.SaveUploadedFile(file, dst)
	}
	c.String(200, "上传成功")
}
