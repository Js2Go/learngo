package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const genderRe = `<td[^>]*><span class="grayL">性别：</span>([^<]+)</td>`

var cRe = regexp.MustCompile(cityRe)
var gRe = regexp.MustCompile(genderRe)

func ParseCity(contents []byte) engine.ParseResult {
	matches := cRe.FindAllSubmatch(contents, -1)
	genderMatch := gRe.FindSubmatch(contents)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		gender := string(genderMatch[1])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name, gender)
			},
		})
	}

	return result
}
