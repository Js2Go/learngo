package main

import (
	"fmt"
	"learngo/crawler/config"
	"learngo/crawler/engine"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
	config2 "learngo/crawler_distributed/config"
	itemsaver "learngo/crawler_distributed/persist/client"
	worker "learngo/crawler_distributed/worker/client"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(
		fmt.Sprintf(":%d", config2.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	
	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}
	
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:        config.StartUrl,
		Parser: engine.NewFuncParser(parser.ParseCityList, config2.ParseCityList),
	})
}
