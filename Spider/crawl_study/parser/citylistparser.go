package parser

import (
    "regexp"
    "../engines"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

//解析城市信息
func ParseCityList(contents []byte) engine.ParseResult {

    re := regexp.MustCompile(cityListRe)
    all := re.FindAllSubmatch(contents, -1)

    result := engine.ParseResult{}
    for _, c := range all {
        result.Items = append(result.Items, string(c[2])) //城市名字
        result.Requests = append(result.Requests, engine.Request{
            Url:        string(c[1]),
            ParserFunc: engine.NilParser,//这里将函数指针赋值
        })
    }

    return result
}