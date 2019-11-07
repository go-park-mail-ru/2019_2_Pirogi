package common

import "encoding/json"

func MakeJSONArray(items [][]byte) []byte {
	var jsonBody = []byte{'['}
	for idx, item := range items {
		if idx > 0 {
			jsonBody = append(jsonBody, []byte(",")...)
		}
		jsonBody = append(jsonBody, item...)
	}
	jsonBody = append(jsonBody, []byte{']'}...)
	return jsonBody
}

func UnionToJSON(items ...interface{}) (response []byte) {
	for _, item := range items {
		body, err := json.Marshal(item)
		if err != nil {
			continue
		}
		response = append(response, body...)
		response = append(response, byte(','))
	}
	return
}
