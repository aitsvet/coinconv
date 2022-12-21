package json_extractor

import (
	"encoding/json"
	"fmt"
	"strings"
)

type StringField struct{}

func (e *StringField) Extract(from string, path []string) (string, error) {
	var jsonMap map[string]interface{}
	dec := json.NewDecoder(strings.NewReader(from))
	dec.UseNumber()
	err := dec.Decode(&jsonMap)
	if err != nil {
		return "", err
	}
	var ok bool
	pathString := ""
	jsonPart := interface{}(jsonMap)
	for _, key := range path {
		jsonMap, ok = jsonPart.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("could not convert to map at %s", pathString)
		}
		pathString += "." + key
		jsonPart, ok = jsonMap[key]
		if !ok {
			return "", fmt.Errorf("could not extract key at %s", pathString)
		}
	}
	result, ok := jsonPart.(string)
	if !ok {
		stringer, ok := jsonPart.(fmt.Stringer)
		if !ok {
			return "", fmt.Errorf("could not convert to string %s", pathString)
		}
		result = stringer.String()
	}
	return result, nil
}
