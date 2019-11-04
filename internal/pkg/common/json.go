package common

func MakeJSONArray(items [][]byte) []byte {
	var jsonBody = []byte{'['}
	for idx, item := range items {
		if idx > 0 {
			jsonBody = append(jsonBody, []byte(",")...)
		}
		jsonBody = append(jsonBody, item...)
	}
	return jsonBody
}
