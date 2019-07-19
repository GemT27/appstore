package keyword

import (
	"errors"
	"github.com/beevik/etree"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func DecodeXML(xml []byte) (string, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(xml); err != nil {
		return "", err
	}
	array := doc.SelectElement("plist").SelectElement("dict").SelectElement("array")

	if array.FindElement("dict") == nil {
		return "", errors.New("not found key")
	}
	result := make(map[string]string)
	for _, root := range array.SelectElements("dict") {
		name := root.SelectElement("string")
		hot := root.SelectElement("integer")
		if name != nil && hot != nil {
			result[name.Text()] = hot.Text()
		}
	}
	jsonByte, err := json.Marshal(&result)
	if err != nil {
		return "", err
	}
	return string(jsonByte), nil
}

func DecodeXML2(xml []byte) (map[string]string, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(xml); err != nil {
		return nil, err
	}
	array := doc.SelectElement("plist").SelectElement("dict").SelectElement("array")

	if array.FindElement("dict") == nil {
		return nil, errors.New("not found key")
	}
	result := make(map[string]string)
	for _, root := range array.SelectElements("dict") {
		name := root.SelectElement("string")
		hot := root.SelectElement("integer")
		if name != nil && hot != nil {
			result[name.Text()] = hot.Text()
		}
	}
	return result, nil
}
