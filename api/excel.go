package api

import (
	"github.com/gin-gonic/gin"
	"herbalBody/service"
	"herbalBody/util"
	"log"
	"net/url"
	"strconv"
)

func GetExcel(c *gin.Context) {
	//先生成xlsx文件
	userId := c.PostForm("user_id")
	ID, err := strconv.Atoi(userId)
	if err != nil {
		util.RespInternalErr(c)
		log.Printf("strconv atoi err:%v\n", err)
		return
	}
	filename, f, err := service.GenExcel(ID)
	if err != nil {
		if err.Error() == "用户答案未提交完全" {
			log.Printf("gen excel err:%v", err.Error())
			util.NormErr(c, 447, "用户答案未提交完全")
			return
		}
		util.RespInternalErr(c)
		log.Printf("gen excel err:%v\n", err)
		return
	}
	// 提供下载
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	disposition := "attachment; filename*=UTF-8''" + url.PathEscape(filename[1:])
	c.Writer.Header().Set("Content-Disposition", disposition)
	_ = f.Write(c.Writer)

}
