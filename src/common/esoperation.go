package common

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"reflect"
	"strconv"
)

func CreateArticle(article ZfInfolist, id int) {
	es := ES
	esariticle := ZfInfolist{
		Id:                article.Id,
		Siteid:            article.Siteid,
		Classid:           article.Classid,
		Parentid:          article.Parentid,
		Parentstr:         article.Parentstr,
		Mainid:            article.Mainid,
		Mainpid:           article.Mainpid,
		Mainpstr:          article.Mainpstr,
		Title:             article.Title,
		Colorval:          article.Colorval,
		Boldval:           article.Boldval,
		Flag:              article.Flag,
		Source:            article.Source,
		Author:            article.Author,
		Linkurl:           article.Linkurl,
		Keywords:          article.Keywords,
		Description:       article.Description,
		Content:           article.Content,
		ContentText:       "",
		Picurl:            article.Picurl,
		Picarr:            article.Picarr,
		Hits:              article.Hits,
		Orderid:           article.Orderid,
		Posttime:          article.Posttime,
		Checkinfo:         article.Checkinfo,
		Delstate:          article.Delstate,
		Deltime:           article.Deltime,
		Live_prov:         article.Live_prov,
		Live_city:         article.Live_city,
		Live_country:      article.Live_country,
		Zc_type:           article.Zc_type,
		Zc_zb:             article.Zc_zb,
		N_se:              article.N_se,
		J_zc:              article.J_zc,
		Zc_ze:             article.Zc_ze,
		Ss_qk:             article.Ss_qk,
		Gx_jsqy:           article.Gx_jsqy,
		Hy_ly:             article.Hy_ly,
		Hy_ly2:            article.Hy_ly2,
		Z_tz:              article.Z_tz,
		Zs_cq:             article.Zs_cq,
		Jb:                article.Jb,
		Bw:                article.Bw,
		Cy:                article.Cy,
		Cy2:               article.Cy2,
		Td_gm:             article.Td_gm,
		Jsjd_zq:           article.Jsjd_zq,
		Spsx:              article.Spsx,
		Ry_zc:             article.Ry_zc,
		Year:              article.Year,
		Zm:                article.Zm,
		Zm_order:          article.Zm_order,
		Ckkj:              article.Ckkj,
		Rzhdk:             article.Rzhdk,
		Qhdq:              article.Qhdq,
		Implement_country: article.Implement_country,
		Implement_city:    article.Implement_city,
		Implement:         article.Implement,
		Taxes_country:     article.Taxes_country,
		Taxes_city:        article.Taxes_city,
		Taxes:             article.Taxes,
		Start_time:        article.Start_time,
		End_time:          article.End_time,
		Zcfx:              article.Zcfx,
		Zced:              article.Zced,
		Sblc:              article.Sblc,
		Zgbm:              article.Zgbm,
		Sbsj:              article.Sbsj,
		Zlxq:              article.Zlxq,
		Jshz:              article.Jshz,
		Sfbz:              article.Sfbz,
		Jfsj:              article.Jfsj,
		Bggs:              article.Bggs,
		Shfw:              article.Shfw,
		Lxr:               article.Lxr,
		Zxdh:              article.Zxdh,
		Gjz:               article.Gjz,
		Qy_prov:           article.Qy_prov,
		Qy_city:           article.Qy_city,
		Qy_country:        article.Qy_country,
		Stime:             article.Stime,
		Etime:             article.Etime,
		Extended_cid2:     article.Extended_cid2,
		Ym:                article.Ym,
		Registertime:      article.Registertime,
		Extended_cid:      article.Extended_cid,
		N_sesymbol:        article.N_sesymbol,
		Zc_zesymbol:       article.Zc_zesymbol,
		Z_tzsymbol:        article.Z_tzsymbol,
		Zc_zbsymbol:       article.Zc_zbsymbol,
		Td_gm_zesymbol:    article.Td_gm_zesymbol,
		J_zcsymbol:        article.J_zcsymbol,
		Week_time:         article.Week_time,
		Month_time:        article.Month_time,
		Season_time:       article.Season_time,
	}
	idstr := strconv.Itoa(id)
	//数据存在则不会插入es
	if IsExists(idstr) {
		log.Println("数据已存在，数据id为：", id)
		return
	}
	put, err := es.Index().
		Index("govfunds").
		Id(idstr).
		BodyJson(esariticle).
		Do(context.Background())
	if err != nil {
		log.Println("创建索引错误")
		panic(err)
	}
	log.Printf("数据创建成功，数据id为%s,文章标题为：%s", put.Id, article.Title)
}

//按id查询文章是否存在
func IsExists(id string) bool {
	es := ES
	res, err := es.Get().
		Index("govfunds").
		Id(id).
		Do(context.Background())
	if err != nil {
		log.Println("es 查询错误", err)
		return false
	}
	return res.Found
}

//分页查询文章
func SearchArticlePaging(keyword, cy2 string, page, size, year, jb, qyProv, qyCity, qyCounty, bw, cy int) *elastic.SearchResult {
	es := ES
	query := elastic.NewBoolQuery()
	isNoKeyword := true
	//年份
	if year != 0 {
		yearDefault := elastic.NewTermQuery("year", -1)
		yearValue := elastic.NewTermQuery("year", year)
		shouldYear := elastic.NewBoolQuery().Should(yearDefault, yearValue)
		query = query.Must(shouldYear)
		isNoKeyword = false
	}
	//级别
	if jb != 0 {
		jbDefault := elastic.NewTermQuery("jb", -1)
		jbValue := elastic.NewTermQuery("jb", jb)
		shouldJB := elastic.NewBoolQuery().Should(jbDefault, jbValue)
		query = query.Must(shouldJB)
		isNoKeyword = false
	}

	//省份
	if qyProv != 0 {
		qyProvDefault := elastic.NewTermQuery("qy_prov", -1)
		qyProvValue := elastic.NewTermQuery("qy_prov", qyProv)
		shouldProv := elastic.NewBoolQuery().Should(qyProvDefault, qyProvValue)
		query = query.Must(shouldProv)
		isNoKeyword = false
	}

	//城市
	if qyCity != 0 {
		qyCityDefault := elastic.NewTermQuery("qy_city", -1)
		qyCityValue := elastic.NewTermQuery("qy_city", qyCity)
		shouldCity := elastic.NewBoolQuery().Should(qyCityDefault, qyCityValue)
		query = query.Must(shouldCity)
		isNoKeyword = false
	}
	//县
	if qyCounty != 0 {
		qyCountyDefault := elastic.NewTermQuery("live_county", -1)
		qyCountyValue := elastic.NewTermQuery("live_county", qyCounty)
		shouldCounty := elastic.NewBoolQuery().Should(qyCountyDefault, qyCountyValue)
		query = query.Must(shouldCounty)
		isNoKeyword = false
	}
	//部委
	if bw != 0 {
		bwDefault := elastic.NewTermQuery("bw", -1)
		bwValue := elastic.NewTermQuery("bw", bw)
		shouldBW := elastic.NewBoolQuery().Should(bwDefault, bwValue)
		query = query.Must(shouldBW)
		isNoKeyword = false
	}

	//产业1
	if cy != 0 {
		cyDefault := elastic.NewTermQuery("cy", -1)
		cyValue := elastic.NewTermQuery("cy", cy)
		shouldCY := elastic.NewBoolQuery().Should(cyDefault, cyValue)
		query = query.Must(shouldCY)
		isNoKeyword = false
	}

	//产业2
	if len(cy2) != 0 {
		cy2Default := elastic.NewTermQuery("cy2", "")
		cy2Value := elastic.NewTermQuery("cy2", cy2)
		shouldCY2 := elastic.NewBoolQuery().Should(cy2Default, cy2Value)
		query = query.Must(shouldCY2)
		isNoKeyword = false
	}
	if len(keyword) != 0 {
		multi := elastic.NewMultiMatchQuery(keyword, "title", "content").TieBreaker(0.3)
		query = query.Must(multi)
		isNoKeyword = false
	}
	all := elastic.NewMatchAllQuery()
	fsc := elastic.NewFetchSourceContext(true).Include("title", "classid", "id", "jb", "end_time", "year")
	if isNoKeyword {
		res, err := es.Search().
			Index("govfunds").
			Query(all).
			FetchSourceContext(fsc).
			Pretty(true).
			Sort("year", false).
			From((page - 1) * size).
			Size(size).
			Do(context.Background())
		if err != nil {
			log.Println("es 查询错误", err)
			panic(err)
		}
		return res
	}
	res, err := es.Search().
		Index("govfunds").
		Query(query).
		FetchSourceContext(fsc).
		Pretty(true).
		Sort("year", false).
		From((page - 1) * size).
		Size(size).
		Do(context.Background())
	if err != nil {
		log.Println("es 查询错误", err)
		panic(err)
	}
	return res
}

//打印搜索结果内容
func printArticleInfo(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ ZfInfolist
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(ZfInfolist)
		fmt.Printf("%d,%#v,%d\n", t.Id, t.Title, t.Posttime)
	}
}

//查询当前govfunds索引的最后一个id
func SearchArticleId(es *elastic.Client) int {
	var esId int
	qureyById := elastic.NewMatchAllQuery()
	res, err := es.Search().
		Index("govfunds").
		Query(qureyById).
		Sort("id", false).
		Size(1).Do(context.Background())
	if err != nil {
		log.Println("无法获取当前es最大id", err)
		return 0
	}
	var typ ZfInfolist
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(ZfInfolist)
		esId = t.Id
	}
	return esId
}

//查询es当前最新更新时间戳
func SearchNewPostTime() int {
	es := ES
	var postTime int
	qureyById := elastic.NewMatchAllQuery()
	res, err := es.Search().
		Index("govfunds").
		Query(qureyById).
		Sort("posttime", false).
		Size(1).Do(context.Background())
	if err != nil {
		log.Println("无法获取当前es最新posttime", err)
		return 0
	}
	var typ ZfInfolist
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(ZfInfolist)
		postTime = t.Posttime
	}
	return postTime
}

//删除文章
func DeleteArticle(es *elastic.Client, idstr string) {

	res, err := es.Delete().Index("govfunds").
		Id(idstr).
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

//更新文章
func UpdateArticle(article ZfInfolist, id string) {
	es := ES
	esariticle := ZfInfolist{
		Id:                article.Id,
		Siteid:            article.Siteid,
		Classid:           article.Classid,
		Parentid:          article.Parentid,
		Parentstr:         article.Parentstr,
		Mainid:            article.Mainid,
		Mainpid:           article.Mainpid,
		Mainpstr:          article.Mainpstr,
		Title:             article.Title,
		Colorval:          article.Colorval,
		Boldval:           article.Boldval,
		Flag:              article.Flag,
		Source:            article.Source,
		Author:            article.Author,
		Linkurl:           article.Linkurl,
		Keywords:          article.Keywords,
		Description:       article.Description,
		Content:           article.Content,
		ContentText:       "",
		Picurl:            article.Picurl,
		Picarr:            article.Picarr,
		Hits:              article.Hits,
		Orderid:           article.Orderid,
		Posttime:          article.Posttime,
		Checkinfo:         article.Checkinfo,
		Delstate:          article.Delstate,
		Deltime:           article.Deltime,
		Live_prov:         article.Live_prov,
		Live_city:         article.Live_city,
		Live_country:      article.Live_country,
		Zc_type:           article.Zc_type,
		Zc_zb:             article.Zc_zb,
		N_se:              article.N_se,
		J_zc:              article.J_zc,
		Zc_ze:             article.Zc_ze,
		Ss_qk:             article.Ss_qk,
		Gx_jsqy:           article.Gx_jsqy,
		Hy_ly:             article.Hy_ly,
		Hy_ly2:            article.Hy_ly2,
		Z_tz:              article.Z_tz,
		Zs_cq:             article.Zs_cq,
		Jb:                article.Jb,
		Bw:                article.Bw,
		Cy:                article.Cy,
		Cy2:               article.Cy2,
		Td_gm:             article.Td_gm,
		Jsjd_zq:           article.Jsjd_zq,
		Spsx:              article.Spsx,
		Ry_zc:             article.Ry_zc,
		Year:              article.Year,
		Zm:                article.Zm,
		Zm_order:          article.Zm_order,
		Ckkj:              article.Ckkj,
		Rzhdk:             article.Rzhdk,
		Qhdq:              article.Qhdq,
		Implement_country: article.Implement_country,
		Implement_city:    article.Implement_city,
		Implement:         article.Implement,
		Taxes_country:     article.Taxes_country,
		Taxes_city:        article.Taxes_city,
		Taxes:             article.Taxes,
		Start_time:        article.Start_time,
		End_time:          article.End_time,
		Zcfx:              article.Zcfx,
		Zced:              article.Zced,
		Sblc:              article.Sblc,
		Zgbm:              article.Zgbm,
		Sbsj:              article.Sbsj,
		Zlxq:              article.Zlxq,
		Jshz:              article.Jshz,
		Sfbz:              article.Sfbz,
		Jfsj:              article.Jfsj,
		Bggs:              article.Bggs,
		Shfw:              article.Shfw,
		Lxr:               article.Lxr,
		Zxdh:              article.Zxdh,
		Gjz:               article.Gjz,
		Qy_prov:           article.Qy_prov,
		Qy_city:           article.Qy_city,
		Qy_country:        article.Qy_country,
		Stime:             article.Stime,
		Etime:             article.Etime,
		Extended_cid2:     article.Extended_cid2,
		Ym:                article.Ym,
		Registertime:      article.Registertime,
		Extended_cid:      article.Extended_cid,
		N_sesymbol:        article.N_sesymbol,
		Zc_zesymbol:       article.Zc_zesymbol,
		Z_tzsymbol:        article.Z_tzsymbol,
		Zc_zbsymbol:       article.Zc_zbsymbol,
		Td_gm_zesymbol:    article.Td_gm_zesymbol,
		J_zcsymbol:        article.J_zcsymbol,
		Week_time:         article.Week_time,
		Month_time:        article.Month_time,
		Season_time:       article.Season_time,
	}

	_, err := es.Update().
		Index("govfunds").
		Id(id).
		Doc(esariticle).
		Do(context.Background())
	if err != nil {
		log.Println("创建索引错误")
		panic(err)
	}
}
