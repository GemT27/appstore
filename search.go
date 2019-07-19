package search

import (
	app "github.com/TT527/app-store-search/App"
	keyword "github.com/TT527/app-store-search/KeyWord"
	itunes "github.com/TT527/app-store-search/iTunes"
)

func App2Json(term, country string, num int) (string, error) {
	result, err := app.Sapp2Json(term, country, num)
	if err != nil {
		return "", err
	}
	return result, nil
}

func Itunes2Json(term, country string) (string, error) {
	result, err := itunes.SiTunes2Json(term, country)
	if err != nil {
		return "", err
	}
	return result, nil
}
func Itunes(term, country string) ([]itunes.Result, error) {
	result, err := itunes.SiTunes(term, country)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func KeyWord2Json(term, country string) (string, error) {
	result, err := keyword.Skeyword2Json(term, country)
	if err != nil {
		return "", err
	}
	return result, nil
}
func KeyWord(term, country string) (map[string]string, error) {
	result, err := keyword.SkeyWord(term, country)
	if err != nil {
		return nil, err
	}
	return result, nil
}
