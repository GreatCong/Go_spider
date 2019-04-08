/*
定义2个类型
 */
package engine

//解析返回的结果
type ParseResult struct{
	Requests []Request
	Items []interface{}
}

type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult //处理这个URL的函数 //这里是表示定义函数指针的意思？
}

func NilParser([] byte) ParseResult {
	return ParseResult{}
}
