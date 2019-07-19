package keyword

import (
	"errors"
	"fmt"
	. "github.com/TT527/app-store-search/Country"
	"io/ioutil"
	"net/http"
	"net/url"
)

const kUrl = "https://search.itunes.apple.com/WebObjects/MZSearchHints.woa/wa/hints?clientApplication=Software&%s"

func Search(term, city string) ([]byte, error) {
	v := url.Values{}
	v.Add("term", term)
	res := v.Encode()
	req, err := http.NewRequest("GET", fmt.Sprintf(kUrl, res), nil)
	if err != nil {
		return nil, err
	}
	code, ok := Country[city]
	if !ok {
		return nil, errors.New("not found Country")
	}
	req.Header.Set("X-Apple-Store-Front", fmt.Sprintf("%s-19,29", code[:6]))
	var Client = &http.Client{}
	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return body, nil
}
