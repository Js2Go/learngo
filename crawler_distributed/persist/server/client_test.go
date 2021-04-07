package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)
	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	
	item := engine.Item{
		Url:     "http://album.zhenai.com/u/1629062944",
		Type:    "zhenai",
		Id:      "1629062944",
		Payload: model.Profile{
			Name:       "马子",
			Gender:     "女性",
			Age:        "22岁",
			Height:     "160cm",
			Weight:     "74kg",
			Income:     "月收入:3-5千",
			Marriage:   "未婚",
			Education:  "高中及以下",
			Occupation: "服务业",
			Hokou:      "毕节 ",
			Xinzuo:     "天蝎座(10.23-11.21)",
			House:      "已购房",
			Car:        "已购车",
		},
	}
	// Call Save
	result := ""
	err = client.Call(config.ItemSaverRpc,
		item, &result)
	
	if err != nil || result != "OK" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
