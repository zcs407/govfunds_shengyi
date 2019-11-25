package controller

import (
	"govfunds_shengyi/src/common"
	"log"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Firstsave(ctx *gin.Context) {
	common.FirstSaveMysqlData2es()
	ctx.JSON(200, "同步完成")
}
func Search(ctx *gin.Context) {
	resp := make(map[string]interface{})
	type resArticle struct {
		Id       int    `json:"id"`
		Classid  int    `json:"classid"`
		Title    string `json:"title"`
		End_time int    `json:"end_time"`
		Jb       int    `json:"jb"`
		Year     int    `json:"year"`
	}
	var json struct {
		Value      string `json:"value"`
		Year       string `json:"year"`
		Jb         string `json:"jb"`
		LiveProv   string `json:"live_prov"`
		LiveCity   string `json:"live_city"`
		LiveCounty string `json:"live_county"`
		Bw         string `json:"bw"`
		Cy         string `json:"cy"`
		Cy2        string `json:"cy_2"`
		Page       string `json:"page"`
		Size       string `json:"size"`
		IsContent  bool   `json:"iscontent"`
	}
	var resArticlesInfo []resArticle
	_ = ctx.BindJSON(&json)
	keyword := json.Value
	year, _ := strconv.Atoi(json.Year)
	jb, _ := strconv.Atoi(json.Jb)
	prov, _ := strconv.Atoi(json.LiveProv)
	city, _ := strconv.Atoi(json.LiveCity)
	county, _ := strconv.Atoi(json.LiveCounty)
	bw, _ := strconv.Atoi(json.Bw)
	cy, _ := strconv.Atoi(json.Cy)
	cy2 := json.Cy2
	page := json.Page
	pageInt, _ := strconv.Atoi(page)
	size := json.Size
	sizeInt, _ := strconv.Atoi(size)
	if len(page) == 0 || len(size) == 0 {
		log.Println("没有获取到查询条件")
		resp["code"] = "4004"
		resp["info"] = "value page 或size为空,请设置!"
		resp["resultTotal"] = "0"
		resp["resultInfo"] = resArticlesInfo
		ctx.JSON(200, resp)
		return
	}
	log.Println(json)
	res := common.SearchArticlePaging(keyword, cy2, pageInt, sizeInt, year, jb, prov, city, county, bw, cy, json.IsContent)
	var typ common.ZfInfolist
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(common.ZfInfolist)
		article := resArticle{
			Id:       t.Id,
			Classid:  t.Classid,
			Title:    t.Title,
			End_time: t.End_time,
			Jb:       t.Jb,
			Year:     t.Year,
		}
		resArticlesInfo = append(resArticlesInfo, article)
	}
	log.Println("搜索结果总数:", res.TotalHits())
	resp["code"] = "2000"
	resp["info"] = "查询成功!"
	resp["resultTotal"] = res.TotalHits()
	resp["resultInfo"] = resArticlesInfo
	ctx.JSON(200, resp)
}
func SearchForColumnListInfo(ctx *gin.Context) {
	resp := make(map[string]interface{})
	type resArticle struct {
		Id        int    `json:"id"`
		Classid   int    `json:"classid"`
		Checkinfo string `json:"checkinfo"`
		Classname string `json:"classname"`
		Title     string `json:"title"`
		Posttime  int    `json:"posttime"`
		Author    string `json:"author"`
		Hits      int    `json:"hits"`
	}
	var json struct {
		Year       string `json:"year"`
		Jb         string `json:"jb"`
		Value      string `json:"value"`
		LiveProv   string `json:"live_prov"`
		LiveCity   string `json:"live_city"`
		LiveCounty string `json:"live_county"`
		Bw         string `json:"bw"`
		Cy         string `json:"cy"`
		Cy2        string `json:"cy_2"`
		Page       string `json:"page"`
		Size       string `json:"size"`
		Classid    string `json:"classid"`
		Checkinfo  string `json:"checkinfo"`
		IsContent  bool   `json:"iscontent"`
	}
	var resArticlesInfo []resArticle
	_ = ctx.BindJSON(&json)
	keyword := json.Value
	year, _ := strconv.Atoi(json.Year)
	jb, _ := strconv.Atoi(json.Jb)
	prov, _ := strconv.Atoi(json.LiveProv)
	city, _ := strconv.Atoi(json.LiveCity)
	county, _ := strconv.Atoi(json.LiveCounty)
	bw, _ := strconv.Atoi(json.Bw)
	cy, _ := strconv.Atoi(json.Cy)
	cy2 := json.Cy2
	page := json.Page
	pageInt, _ := strconv.Atoi(page)
	size := json.Size
	sizeInt, _ := strconv.Atoi(size)
	classid, _ := strconv.Atoi(json.Classid)
	checkinfo := json.Checkinfo
	if len(page) == 0 || len(size) == 0 {
		log.Println("没有获取到查询条件")
		resp["code"] = "4004"
		resp["info"] = "value page 或size为空,请设置!"
		resp["resultTotal"] = "0"
		resp["resultInfo"] = resArticlesInfo
		ctx.JSON(200, resp)
		return
	}
	log.Println(json)
	res := common.SearchColumnInfo(keyword, cy2, checkinfo, pageInt, sizeInt, year, jb, prov, city, county, bw, cy, classid, json.IsContent)
	var typ common.ZfInfolist
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(common.ZfInfolist)
		classname := common.GetClassNameByClassId(t.Classid)
		article := resArticle{
			Id:        t.Id,
			Classid:   t.Classid,
			Checkinfo: t.Checkinfo,
			Classname: classname,
			Title:     t.Title,
			Posttime:  t.Posttime,
			Author:    t.Author,
			Hits:      t.Hits,
		}
		resArticlesInfo = append(resArticlesInfo, article)
	}
	log.Println("搜索结果总数:", res.TotalHits())
	resp["code"] = "2000"
	resp["info"] = "查询成功!"
	resp["resultTotal"] = res.TotalHits()
	resp["resultInfo"] = resArticlesInfo
	ctx.JSON(200, resp)
}
