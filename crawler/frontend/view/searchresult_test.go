package view

import (
	"learngo/crawler/engine"
	"learngo/crawler/frontend/model"
	model2 "learngo/crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	//tpl := template.Must(
	//	template.ParseFiles("template.html"))
	
	view := CreateSearchResultView("template.html")
	
	out, err := os.Create("template.test.html")

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1135756888",
		Type: "zhenai",
		Id:   "1135756888",
		Payload: model2.Profile{
			Name:       "马子",
			Gender:     "男性",
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
	
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	
	//err = tpl.Execute(out, page)
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
