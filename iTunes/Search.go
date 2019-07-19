package itunes

import (
	"errors"
	"fmt"
	. "github.com/TT527/app-store-search/Country"
	"io/ioutil"
	"net/http"
	"net/url"
)

const iUrl = "https://search.itunes.apple.com/WebObjects/MZSearch.woa/wa/search?startIndex=0&entity=software&media=software&page=1&restrict=false&prevIndex=300&%s"

func Search(term, city string) ([]byte, error) {
	v := url.Values{}
	v.Add("term", term)
	res := v.Encode()
	req, err := http.NewRequest("GET",
		fmt.Sprintf(iUrl, res), nil)
	if err != nil {
		return nil, err
	}
	code, ok := Country[city]
	if !ok {
		return nil, errors.New("not found Country")
	}
	req.Header.Set("Host", "search.itunes.apple.com")
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Proxy-Connection", "keep-alive")
	req.Header.Set("User-Agent", "iTunes/12.9.5 (Macintosh; OS X 10.14.5) AppleWebKit/607.2.6.1.1  (dt:1)")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("X-Apple-Store-Front", code)
	req.Header.Set("cache-control", "max-age=0")
	var Client = &http.Client{}
	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
