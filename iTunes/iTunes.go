package itunes

func SiTunes2Json(term, country string) (string, error) {
	body, err := Search(term, country)
	if err != nil {
		return "", err
	}
	result, err := DecodeHTML(body)
	if err != nil {
		return "", err
	}
	return result, nil
}

func SiTunes(term, country string) ([]Result, error) {
	body, err := Search(term, country)
	if err != nil {
		return nil, err
	}
	result, err := DecodeHTML2(body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
