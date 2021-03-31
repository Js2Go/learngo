package parser

import (
	"io/ioutil"
	"learngo/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "马子", "女性")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+
			"element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
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
		House:      "",
		Car:        "",
	}

	if profile != expected {
		t.Errorf("expected %v, but was %v", expected, profile)
	}
}
