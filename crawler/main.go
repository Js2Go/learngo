package main

import (
	"learngo/crawler/config"
	"learngo/crawler/engine"
	"learngo/crawler/persist"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
)

func main() {
	//parser.ParseProfile([]byte(
	//	`<div class="CONTAINER f-fl" style="width: 880px;"><div data-v-499ba28c="" class="m-userInfo"><div data-v-499ba28c="" class="top f-cl"><div data-v-499ba28c="" class="logo f-fl" style="background-image: url(&quot;https://photo.zastatic.com/images/photo/327514/1310052470/48995038437672485.jpg?scrop=1&amp;crop=1&amp;cpos=north&amp;w=200&amp;h=200&quot;);"></div> <div data-v-499ba28c="" class="right f-fl"><div data-v-499ba28c="" class="info"><div data-v-499ba28c="" class="name"><h1 data-v-499ba28c="" class="nickName">木木</h1> <div data-v-499ba28c="" class="flag-box"><span data-v-39acb34a="" data-v-499ba28c="" title="珍心会员" class="FLAG zhenxin" style="width: 20px; height: 20px; margin-left: 0px;"></span> <!----> <span data-v-39acb34a="" data-v-499ba28c="" title="实名认证" class="FLAG realname" style="width: 20px; height: 20px; margin-left: 0px; margin-right: 8px;"></span></div> <!----></div> <div data-v-499ba28c="" class="id">ID：1310052470</div> <div data-v-499ba28c="" class="des f-cl">株洲 | 26岁 | 大专 | 未婚 | 158cm | 8001-12000元<a data-v-499ba28c="" href="//www.zhenai.com/n/login?channelId=905819&amp;fromurl=https%3A%2F%2Falbum.zhenai.com%2Fu%2F1310052470" target="_self" class="online f-fr">查看最后登录时间</a></div></div> <div data-v-499ba28c="" class="m-photos"><div data-v-499ba28c="" class="ctrl prev"></div> <div data-v-499ba28c="" class="ctrl next"></div> <div data-v-499ba28c="" class="photoWrapper"><div data-v-499ba28c="" class="photoBox" style="width: 210px; left: 0px;"><div data-v-499ba28c="" href="https://photo.zastatic.com/images/photo/327514/1310052470/48995038437672485.jpg" class="photoItem z-cursor-big active"><img data-v-499ba28c="" src="https://photo.zastatic.com/images/photo/327514/1310052470/48995038437672485.jpg?scrop=1&amp;crop=1&amp;cpos=north&amp;w=110&amp;h=110" alt=""> <div data-v-499ba28c="" class="num">1/2</div></div><div data-v-499ba28c="" href="https://photo.zastatic.com/images/photo/327514/1310052470/49024632127659352.jpg" class="photoItem z-cursor-big"><img data-v-499ba28c="" src="https://photo.zastatic.com/images/photo/327514/1310052470/49024632127659352.jpg?scrop=1&amp;crop=1&amp;cpos=north&amp;w=110&amp;h=110" alt=""> <div data-v-499ba28c="" class="num">2/2</div></div></div></div></div> <!----></div></div> <div data-v-499ba28c="" class="bottom"><div data-v-499ba28c="" class="actions f-cl"><div data-v-499ba28c="" class="item sayHi f-fl">打招呼</div> <div data-v-499ba28c="" class="item sendMsg f-fl">发消息</div> <div data-v-499ba28c="" class="item hongliang f-fl">红娘牵线</div></div></div></div> <div class="CONTAINER" style="width: 100%; margin: 20px auto 0px;"><div data-v-8b1eac0c="" class="m-userInfoDetail"><div data-v-8b1eac0c="" class="m-title">内心独白</div> <div data-v-8b1eac0c="" class="m-content-box m-des"><span data-v-8b1eac0c="">外人眼中的高冷难接近，熟人面前逗比，拥有有趣的灵魂</span><!----></div> <div data-v-8b1eac0c="" class="m-title">个人资料</div> <div data-v-8b1eac0c="" class="m-content-box"><div data-v-8b1eac0c="" class="purple-btns"><div data-v-8b1eac0c="" class="m-btn purple">未婚</div><div data-v-8b1eac0c="" class="m-btn purple">26岁</div><div data-v-8b1eac0c="" class="m-btn purple">魔羯座(12.22-01.19)</div><div data-v-8b1eac0c="" class="m-btn purple">158cm</div><div data-v-8b1eac0c="" class="m-btn purple">45kg</div><div data-v-8b1eac0c="" class="m-btn purple">工作地:株洲芦淞区</div><div data-v-8b1eac0c="" class="m-btn purple">月收入:8千-1.2万</div><div data-v-8b1eac0c="" class="m-btn purple">大专</div></div> <div data-v-8b1eac0c="" class="pink-btns"><div data-v-8b1eac0c="" class="m-btn pink">汉族</div><div data-v-8b1eac0c="" class="m-btn pink">籍贯:湖南邵阳</div><div data-v-8b1eac0c="" class="m-btn pink">体型:苗条</div><div data-v-8b1eac0c="" class="m-btn pink">不吸烟</div><div data-v-8b1eac0c="" class="m-btn pink">没有小孩</div><div data-v-8b1eac0c="" class="m-btn pink">何时结婚:时机成熟就结婚</div></div></div> <div data-v-8b1eac0c="" class="m-title">兴趣爱好</div> <div data-v-8b1eac0c="" class="m-content-box m-interest f-cl"><div data-v-8b1eac0c="" class="item f-fl"><div data-v-8b1eac0c="" class="question f-fl">喜欢的一道菜：</div> <div data-v-8b1eac0c="" class="answer f-fl">未填写</div></div><div data-v-8b1eac0c="" class="item f-fl"><div data-v-8b1eac0c="" class="question f-fl">欣赏的一个名人：</div> <div data-v-8b1eac0c="" class="answer f-fl">未填写</div></div><div data-v-8b1eac0c="" class="item f-fl"><div data-v-8b1eac0c="" class="question f-fl">喜欢的一首歌：</div> <div data-v-8b1eac0c="" class="answer f-fl">未填写</div></div><div data-v-8b1eac0c="" class="item f-fl"><div data-v-8b1eac0c="" class="question f-fl">喜欢的一本书：</div> <div data-v-8b1eac0c="" class="answer f-fl">未填写</div></div><div data-v-8b1eac0c="" class="item f-fl"><div data-v-8b1eac0c="" class="question f-fl">喜欢做的事：</div> <div data-v-8b1eac0c="" class="answer f-fl">未填写</div></div></div> <div data-v-8b1eac0c="" class="m-title">择偶条件</div> <div data-v-8b1eac0c="" class="m-content-box"><div data-v-8b1eac0c="" class="gray-btns"><div data-v-8b1eac0c="" class="m-btn">26-32岁</div><div data-v-8b1eac0c="" class="m-btn">172cm以上</div><div data-v-8b1eac0c="" class="m-btn">工作地:湖南邵阳大祥区</div><div data-v-8b1eac0c="" class="m-btn">月薪:1.2万以上</div><div data-v-8b1eac0c="" class="m-btn">未婚</div><div data-v-8b1eac0c="" class="m-btn">没有小孩</div></div></div> <div data-v-8b1eac0c="" class="m-title">她的动态</div> <div data-v-8b1eac0c="" class="m-content-box m-news"><p data-v-8b1eac0c="">该用户还发布了<span data-v-8b1eac0c="">6条</span>动态分享她的生活<br data-v-8b1eac0c="">扫描下载珍爱APP查看她的动态</p> <div data-v-8b1eac0c="" class="app"></div></div></div></div></div>`))
	//e := engine.ConcurrentEngine{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	WorkerCount: 100,
	//}
	//e.Run(engine.Request{
	//	Url: config.StartUrl,
	//	ParserFunc: parser.ParseCityList,
	//})

	//e := engine.SimpleEngine{}
	//e.Run(engine.Request{
	//	Url: config.StartUrl,
	//	ParserFunc: parser.ParseCityList,
	//})
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:        config.StartUrl,
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}
