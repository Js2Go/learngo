package parser

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"learngo/crawler_distributed/config"
	"regexp"
)

//const ageRe = `<div[^>]*class="m-btn purple">([\d]+)岁</div>`
//const nameRe = `<h1[^>]*class="nickName">([^<]+)</h1>`
//const ocpRe = `<div[^>]*class="des f-cl"[^>]*>([^|]+)`

//var re = regexp.MustCompile(config.BaseProfileRe)

//var nRe = regexp.MustCompile(nameRe)
//var oRe = regexp.MustCompile(ocpRe)

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func parseProfile(
	contents []byte, url, name, gender string) engine.ParseResult {
	//matches := re.FindAllSubmatch(contents, -1)
	//nameMatch := nRe.FindSubmatch(contents)
	//ocpMatch := oRe.FindSubmatch(contents)


	//fmt.Println(name, gender)
	profile := model.Profile{
		Name:       name,
		Gender:     gender,
		//Age:        string(matches[1][1]),
		//Height:     string(matches[3][1]),
		//Weight:     string(matches[4][1]),
		//Income:     string(matches[6][1]),
		//Marriage:   string(matches[0][1]),
		//Education:  string(matches[8][1]),
		//Occupation: string(matches[7][1]),
		//Hokou:      string(ocpMatch[1]),
		//Xinzuo:     string(matches[2][1]),
		House:      "",
		Car:        "",
	}

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Id: extractString([]byte(url), idUrlRe),
				Url: url,
				Type: "zhenai",
				Payload: profile,
			},
		},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ParseProfile struct {
	UserName string
	Gender string
}

func (p *ParseProfile) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.UserName, p.Gender)
}

func (p *ParseProfile) Serialize() (name string, args interface{}) {
	return config.ParseProfile, ParseProfile{
		UserName: p.UserName,
		Gender: p.Gender,
	}
}

func NewProfileParser(name string, gender string) *ParseProfile {
	return &ParseProfile{
		UserName: name,
		Gender: gender,
	}
}
