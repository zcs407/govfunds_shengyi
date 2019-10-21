package common

import (
	"context"
	"github.com/olivere/elastic"
	"log"
	"os"
)

var ES *elastic.Client
var err error
var esHost = "http://127.0.0.1:9200/"

func InitES() {
	errorLog := log.New(os.Stdout, "app", log.LstdFlags)
	ES, err = elastic.NewClient(elastic.SetErrorLog(errorLog), elastic.SetURL(esHost))
	if err != nil {
		log.Println("create es client error:", err)
		panic(err)
	}
	_, _, err := ES.Ping(esHost).Do(context.Background())
	if err != nil {
		log.Println("无法连接es server, err:", err)
		panic(err)
	}
}
