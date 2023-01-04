package utils

import "encoding/json"

const DefaultJsonString = "{}"

func Dump(data interface{}) (res string) {
	res = DefaultJsonString
	if data == nil {
		return
	}
	switch data.(type) {
	case []byte:
		return string(data.([]byte))
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return
	}
	return string(bytes)
}
