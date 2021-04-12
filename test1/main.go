package main

import (
	"fmt"
	"regexp"
)

var (
	hokouRe = regexp.MustCompile(
		`<td[^>]*><span class="grayL">居住地：</span>([^<]+)</td>`)
	ageRe = regexp.MustCompile(
		`<td[^>]*><span class="grayL">年龄：</span>([^<]+)</td>`)
	eduRe = regexp.MustCompile(
		`<td[^>]*><span class="grayL">学&nbsp;&nbsp;&nbsp;历：</span>([^<]+)</td>`)
	incomeRe = regexp.MustCompile(
		`<td[^>]*><span class="grayL">月&nbsp;&nbsp;&nbsp;薪：</span>([^<]+)</td>`)
	marRe = regexp.MustCompile(
		`<td[^>]*><span class="grayL">婚况：</span>([^<]+)</td>`)
	heightRe = regexp.MustCompile(
		`<td[^>]*><span class="grayL">身&nbsp;&nbsp;&nbsp;高：</span>([^<]+)</td>`)
	avatarRe = regexp.MustCompile(
		`<div class="photo"><a[^>]*><img src="([^"]+)"[^>]*></a></div>`)
)

func main() {
	html := `<div class="list-item"><div class="photo"><a href="http://album.zhenai.com/u/1241762648" target="_blank"><img src="https://photo.zastatic.com/images/photo/310441/1241762648/8029484725579649.png?scrop=1&amp;crop=1&amp;w=140&amp;h=140&amp;cpos=north" alt="寻找另一个的你"></a></div> <div class="content"><table><tbody><tr><th><a href="http://album.zhenai.com/u/1241762648" target="_blank">寻找另一个的你</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>男士</td> <td><span class="grayL">居住地：</span>北京</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>38</td> <!----> <td><span class="grayL">月&nbsp;&nbsp;&nbsp;薪：</span>8001-12000元</td></tr> <tr><td width="180"><span class="grayL">婚况：</span>未婚</td> <td width="180"><span class="grayL">身&nbsp;&nbsp;&nbsp;高：</span>180</td></tr></tbody></table> <div class="introduce">你觉得几岁差距可以接受？但是我看上去比同龄人年轻十岁。之前遇到的是年龄相差很大的女朋友，最后还是没有结果。真的想结婚了，一起走下去。请联系我吧。</div></div> <div class="item-btn">打招呼</div></div>`
	matchString := incomeRe.FindStringSubmatch(html)
	fmt.Println(matchString[1])
}
