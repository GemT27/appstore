package app

func Sapp2Json(term, country string, num int) (string, error) {
	result, err := Search(term, country, num)
	if err != nil {
		return "", err
	}
	return result, nil
}
