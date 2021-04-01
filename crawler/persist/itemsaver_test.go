package persist

import (
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
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

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	
	
	const index = "dating_test"
	err = save(client, index, expected)
	
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)

	if err != nil {
		panic(err)
	}
	
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
