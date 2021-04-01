package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s", contents)
	results := ParseCityList(contents, "")

	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"City 阿坝",
		"City 阿克苏",
		"City 阿拉善盟",
	}
	if len(results.Requests) != resultSize {
		t.Errorf("result should have %d" +
			"requests; but had %d",
			resultSize,
			len(results.Requests))
	}

	for i, url := range expectedUrls {
		if results.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but" +
				"was %s",
				i, results.Requests[i].Url, url)
		}
	}

	for i, city := range expectedCities {
		if results.Items[i].Payload.(string) != city {
			t.Errorf("expected url #%d: %s; but" +
				"was %s",
				i, city, results.Items[i].Payload.(string))
		}
	}

	// verify result
}
