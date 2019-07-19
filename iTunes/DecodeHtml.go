package itunes

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func DecodeHTML(html []byte) (string, error) {
	body := bytes.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return "", err
	}
	var results []Result
	doc.Find(".paginated-content ").Each(func(i int, s *goquery.Selection) {
		s.Find("ul").Each(func(i int, selection *goquery.Selection) {
			href, _ := selection.Find(".name a").Attr("href")
			genre := selection.Find(".genre").Text()
			name := selection.Find(".name").Text()
			results = append(results, Result{Href: href, ID: i + 1, Genre: genre, Name: name})
		})
		s.Find(".mobile-app-icon-clip img").Each(func(i int, selection *goquery.Selection) {
			logo, _ := selection.Attr("src-swap")
			results[i].Logo = logo
		})
	})

	jsonByte, err := json.Marshal(&results)
	if err != nil {
		return "", err
	}
	return string(jsonByte), nil
}

func DecodeHTML2(html []byte) ([]Result, error) {
	body := bytes.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}
	var results []Result
	doc.Find(".paginated-content ").Each(func(i int, s *goquery.Selection) {
		s.Find("ul").Each(func(i int, selection *goquery.Selection) {
			href, _ := selection.Find(".name a").Attr("href")
			genre := selection.Find(".genre").Text()
			name := selection.Find(".name").Text()
			results = append(results, Result{Href: href, ID: i + 1, Genre: genre, Name: name})
		})
		s.Find(".mobile-app-icon-clip img").Each(func(i int, selection *goquery.Selection) {
			logo, _ := selection.Attr("src-swap")
			results[i].Logo = logo
		})
	})

	return results, nil
}

type Result struct {
	ID    int    `json:"id"`
	Logo  string `json:"logo"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Href  string `json:"href"`
}
