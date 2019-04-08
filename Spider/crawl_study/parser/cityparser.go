package parser

/*
从城市列表中，我们知道了相关的URl,接下来我们可以获取城市的信息
*/

import (
	"../engines"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

/*
解析城市相关具体信息
这样我们每次爬取一个城市的Url,也会获取该城市的第一页的用户的Url和用户名
*/
func parse_city(contens []byte) engine.ParseResult {

	re := regexp.MustCompile(cityRe)

	all := re.FindAllSubmatch(contens, -1)

	result := engine.ParseResult{} //定义一个类

	for _, c := range all {
		result.Items = append(result.Items, "User:"+string(c[2])) //用户名字
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(c[1]),
			ParserFunc: engine.NilParser,
		})
	}

	return result

}
