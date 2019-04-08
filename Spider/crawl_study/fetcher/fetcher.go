/*
对url进行fetcher，获取这个url对应的utf-8的文本。
*/
package fetcher

import (
	"bufio" //带缓存的IO包
	"fmt"
	"io/ioutil"
	"net/http"
	// "log"
	// "golang.org/x/text/encoding"//要用到golang/x库
	// 	"golang.org/x/text/transform"
	// 	"bufio" //带缓存的IO包
	// 	"golang.org/x/net/html/charset"
	// 	"golang.org/x/text/encoding/unicode"
)

func Fetch(url string) ([]byte, error) {
	/*
		resp,err := http.Get(url)

		if(err != nil){
			return nil,err
		}
	*/
	//如果访问用户页面，没有请求头，会返回403
	//add header
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	resp, _ := http.DefaultClient.Do(request)
	//add header end
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error:status code:%d", resp.StatusCode)
	}

	//如果页面不是utf8，需要转换为utf8格式
	bodyReader := bufio.NewReader(resp.Body)
	//e := determineEncoding(bodyReader)

	//utf8Reader := transform.NewReader(bodyReader,e.NewDecoder())

	// return ioutil.ReadAll(utf8Reader)
	return ioutil.ReadAll(bodyReader)
}

// //转换编码
// func determineEncoding(r *bufio.NewReader) encoding.Encoding{//r *bufio.NewReader定义bufio.NewReader的指针r
// 	bytes ,err := r.Peek(1024)

// 	if(err !=nil){
// 		log.Printf("Fecher error:%v",err) //go中有格式化%v
// 		return unicode.UTF8
// 	}

// 	e,_,_ := charset.determineEncoding(bytes,"")
// 	return e;
// }
