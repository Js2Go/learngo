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
	contents []byte, url, name, gender, hokou, age, edu, income, mar, height, avatar string) engine.ParseResult {
	//matches := re.FindAllSubmatch(contents, -1)
	//nameMatch := nRe.FindSubmatch(contents)
	//ocpMatch := oRe.FindSubmatch(contents)

	//fmt.Println(name, gender)
	profile := model.Profile{
		Name:      name,
		Gender:    gender,
		Age:       age,
		Height:    height,
		Income:    income,
		Marriage:  mar,
		Education: edu,
		Hokou:  hokou,
		Avatar: avatar,
		//House:  "",
		//Car:    "",
		//Weight:     string(matches[4][1]),
		//Occupation: string(matches[7][1]),
		//Xinzuo:     string(matches[2][1]),
	}

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Id:      extractString([]byte(url), idUrlRe),
				Url:     url,
				Type:    "zhenai",
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

// unknown parser name error 是因为结构体的名字必须和调用rpc的名字保持一致
type ParseProfile struct {
	UserName string
	Gender   string
	Hokou    string
	Age      string
	Edu      string
	Income   string
	Mar      string
	Height   string
	Avatar   string
}

func (p *ParseProfile) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(
		contents, url, p.UserName, p.Gender, p.Hokou,
		p.Age, p.Edu, p.Income, p.Mar, p.Height, p.Avatar)
}

func (p *ParseProfile) Serialize() (name string, args interface{}) {
	return config.ParseProfile, ParseProfile{
		UserName: p.UserName,
		Gender:   p.Gender,
		Hokou:    p.Hokou,
		Age:      p.Age,
		Edu:      p.Edu,
		Income:   p.Income,
		Mar:      p.Mar,
		Height:   p.Height,
		Avatar:   p.Avatar,
	}
}

func NewProfileParser(
	name, gender, hokou, age, edu, income, mar, height, avatar string) *ParseProfile {
	return &ParseProfile{
		UserName: name,
		Gender:   gender,
		Hokou:    hokou,
		Age:      age,
		Edu:      edu,
		Income:   income,
		Mar:      mar,
		Height:   height,
		Avatar:   avatar,
	}
}
