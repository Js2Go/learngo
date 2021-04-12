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
	hokouRe = regexp.MustCompile(
		`<td[^>]*><span class="grayL">居住地：</span>([^<]+)</td>`)
	ageRe = regexp.MustCompile(
		`<td[^>]*><span class="grayL">年龄：</span>([^<]+)</td>`)
	eduRe = regexp.MustCompile(
		`<td[^>]*><span class="grayL">学   历：</span>([^<]+)</td>`)
	incomeRe = regexp.MustCompile(
		`<td[^>]*><span class="grayL">月   薪：</span>([^<]+)</td>`)
	marRe = regexp.MustCompile(
		`<td[^>]*><span class="grayL">婚况：</span>([^<]+)</td>`)
	heightRe = regexp.MustCompile(
		`<td[^>]*><span class="grayL">身   高：</span>([^<]+)</td>`)
	avatarRe = regexp.MustCompile(
		`<div class="photo"><a[^>]*><img src="([^"]+)"[^>]*></a></div>`)
	cityUrlRe = regexp.MustCompile(
		`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	matches2 := genderRe.FindAllSubmatch(contents, -1)
	matches3 := hokouRe.FindAllSubmatch(contents, -1)
	matches4 := ageRe.FindAllSubmatch(contents, -1)
	//matches5 := eduRe.FindAllSubmatch(contents, -1)
	//matches6 := incomeRe.FindAllSubmatch(contents, -1)
	matches7 := marRe.FindAllSubmatch(contents, -1)
	//matches8 := heightRe.FindAllSubmatch(contents, -1)
	matches9 := avatarRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for idx, m := range matches {
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			Parser: NewProfileParser(
				string(m[2]), string(matches2[idx][1]), string(matches3[idx][1]), string(matches4[idx][1]), "", "", string(matches7[idx][1]), "", string(matches9[idx][1])),
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
