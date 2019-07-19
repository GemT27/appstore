package keyword

func Skeyword2Json(term, country string) (string, error) {
	xml, err := Search(term, country)
	if err != nil {
		return "", err
	}
	result, err := DecodeXML(xml)
	if err != nil {
		return "", err
	}
	return result, nil
}

func SkeyWord(term, country string) (map[string]string, error) {
	xml, err := Search(term, country)
	if err != nil {
		return nil, err
	}
	result, err := DecodeXML2(xml)
	if err != nil {
		return nil, err
	}
	return result, nil
}
