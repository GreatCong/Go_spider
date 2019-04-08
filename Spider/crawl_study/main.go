/*
简单的爬虫例子
 */
package main

import (
"fmt"
"./engines"
"./parser"
)

/*
整个框架的实现：
1.Engine 驱动整个程序，需要种子，一般是启动页面，不止需要url，还需要对应的parser。
2.当种子送到engine上后，先添加到任务队列，然后从队列中取出任务执行。
  执行的时候，对url进行fetcher，获取这个url对应的utf-8的文本。
3.然后将这个文本，通过parser解析器，就能够获取request列表和item列表。
 */

func main() {
	url := "http://www.zhenai.com/zhenghun"
    engine.Run(engine.Request{
        Url:url,
        ParserFunc:parser.ParseCityList,
    })//结构体可以像类json的定义

    fmt.Println("Spider Over!")
}