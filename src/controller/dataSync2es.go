package controller

import (
	"github.com/gin-gonic/gin"
	"govfunds_shengyi/src/common"
)

func DataSync2es(ctx *gin.Context) {
	common.TouchMysql2es()
	resp := make(map[string]interface{})
	resp["code"] = "2000"
	resp["info"] = "数据同步完成!"
	ctx.JSON(200, resp)
}
