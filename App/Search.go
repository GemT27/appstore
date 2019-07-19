package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const aUrl = "https://itunes.apple.com/search?%s"

func Search(term, city string, num int) (string, error) {
	v := url.Values{}
	v.Add("media", "software")
	v.Add("term", term)
	v.Add("limit", strconv.Itoa(num))
	v.Add("country", city)
	query := v.Encode()
	resp, err := http.Get(fmt.Sprintf(aUrl, query))
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result := strings.Trim(string(body), "\r\n")
	return result, nil
}
