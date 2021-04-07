package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	genderRe = regexp.MustCompile(
		`<td[^>]*><span class="grayL">性别：</span>([^<]+)</td>`)
	cityUrlRe = regexp.MustCompile(
		`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	matches2 := genderRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for idx, m := range matches {
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			Parser: NewProfileParser(string(m[2]), string(matches2[idx][1])),
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return result
}
