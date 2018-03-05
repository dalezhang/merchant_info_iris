package utils

import (
	"encoding/json"
)

func MapStruct(st interface{}) (map[string]string, error) {
	var js map[string]string
	jsStr, err := json.Marshal(st)
	if err != nil {
		return js, err
	}

	if err1 := json.Unmarshal([]byte(jsStr), &js); err1 != nil {
		return js, err1
	}
	return js, nil
}
