package common

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

//首次执行将数据库中的数据id保存至文件
func FirstSaveMysqlData2es() {
	db := DBSQL
	zf_infolist := []ZfInfolist{}
	//db.Raw(`SELECT id FROM zf_infolist WHERE (classid = ? OR parentstr LIKE ?)  AND   delstate=? AND checkinfo=? ORDER BY  ? ASC`, 7, "%,7,%", "", true, "id").Find(&zf_infolist)
	db.Raw(`SELECT id FROM zf_infolist`).Find(&zf_infolist)
	for _, article := range zf_infolist {
		CreateMysql2es(article.Id)
	}
}

//创建按id查询数据并保存至es的函数
func CreateMysql2es(id int) {
	db := DBSQL
	zf_infolist := ZfInfolist{}
	db.Where("id = ?", id).First(&zf_infolist)
	CreateArticle(zf_infolist, id)
}

//更新es中的数据
func UpdateMysql2es(id string) {
	db := DBSQL
	zf_infolist := ZfInfolist{}
	db.Where("id = ?", id).First(&zf_infolist)
	UpdateArticle(zf_infolist, id)
}

/*通过SearchNewPostTime函数获取到当前es最新数据，
作为条件查询mysql中满足条件的数据，插入到es中
*/
func TouchMysql2es() {
	postTime := SearchNewPostTime()
	db := DBSQL
	zf_infolist := []ZfInfolist{}
	//db.Raw(`SELECT id FROM zf_infolist WHERE posttime > ? AND (classid = ? OR parentstr LIKE ?)  AND   delstate=? AND checkinfo=?`, postTime, 7, "%,7,%", "", true).Find(&zf_infolist)
	db.Raw(`SELECT id FROM zf_infolist WHERE posttime > ?  AND   delstate=?`, postTime, "").Find(&zf_infolist)
	if len(zf_infolist) == 0 {
		log.Println("当前数据最新无需更新！")
		return
	}
	for _, article := range zf_infolist {
		aid := strconv.Itoa(article.Id)
		if IsExists(aid) {
			UpdateMysql2es(aid)
		} else {
			CreateMysql2es(article.Id)
		}
	}
}
func GetClassNameByClassId(classid int) string {
	db := DBSQL
	infoClass := ZFInfoClass{}
	db.Raw(`SELECT classname FROM zf_infoclass WHERE id = ?`, classid).Find(&infoClass)
	return infoClass.Classname
}
func GetClassList() []ZFInfoClass {
	db := DBSQL
	classList := []ZFInfoClass{}
	db.Raw(`SELECT * FROM zf_infoclass`).Find(&classList)
	return classList
}
