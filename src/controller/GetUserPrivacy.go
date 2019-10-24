package controller

import (
	"github.com/gin-gonic/gin"
	"govfunds_shengyi/src/common"
	"log"
	"strconv"
)

func GetUserPrivacy(ctx *gin.Context) {
	resp := make(map[string]interface{})
	privacys:=[]*common.ZFAdminPrivacy{}
	groupid:=ctx.Query("groupid")
	if len(groupid) == 0||groupid=="0" {
		log.Println("groupid不合法")
		resp["code"] = "5000"
		resp["info"] = "groupid 为0,没有相关数据"
		resp["data"] = nil
		ctx.JSON(200, resp)
		return
	}
	groupId, _ := strconv.Atoi(groupid)
	privacys1 := common.GetUserPrivacyByGroupId(groupId)
	if len(privacys1)==0{
		log.Println("没有查到数据")
		resp["code"] = "4003"
		resp["info"] = "没有查到数据"
		resp["data"] = nil
		ctx.JSON(200, resp)
		return
	}
	for _,priv:=range privacys1{
		privacy:=common.ZFAdminPrivacy{
			Groupid: priv.Groupid,
			Siteid:  priv.Siteid,
			Model:   priv.Model,
			Classid: priv.Classid,
			Action:  priv.Action,
		}

		privacys= append(privacys, &privacy)
	}
	resp["code"] = "2000"
	resp["info"] = "查询成功"
	resp["data"] = privacys

	ctx.JSON(200, resp)
}
