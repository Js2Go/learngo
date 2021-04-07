package main

import (
	"fmt"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
	"learngo/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1629062944",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: parser.ProfileParser{
				UserName: "马子",
				Gender:   "女性",
			},
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
