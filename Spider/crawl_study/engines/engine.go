package engine

import (
"../fetcher"
"log"
)

func Run(seeds ...Request) {//...表示可以接受任意个Request类型参数（Request是自定义结构体）
	var requests []Request

	for _, r := range seeds{//range返回2个值,(index,value)
		requests = append(requests,r)
	}

	for len(requests) > 0{
		// 1. 获取第一个 Request，并从 []requests 移除，实现了一个队列功能
		r := requests[0]
		requests = requests[1:]//??
		log.Printf("Feching %s",r.Url)
		// 2. 使用爬取器进行对 Request.Url 进行爬取
		body, err := fetcher.Fetch(r.Url)
        
        // 如果爬取出错，记录日志
		if(err != nil){
			log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
            continue
		}
        
        // 3. 使用 Request 的解析函数对怕渠道的内容进行解析
		parseResult := r.ParserFunc(body)//从网络上获取数据，然后由不同的解析器解析数据
		// 4. 将解析体中的 []Requests 加到请求任务队列 requests 的尾部
		requests = append(requests,parseResult.Requests...)//把slice里面的内容展开一个一个加进里面
		// 5. 遍历解析出来的实体，直接打印
		for _,item := range parseResult.Items{
			log.Printf("Got item %v",item)
		}
	}
}