package json

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

func UnionToJSON(names []string, items ...[][]byte) (response []byte) {
	addSymbol('{', &response)
	for i, item := range items {
		addKey(names[i], &response)
		response = append(response, MakeJSONArray(item)...)
		if i != len(items)-1 {
			addSymbol(',', &response)
		}
	}
	addSymbol('}', &response)
	return
}

func UnionToJSONBytes(names []string, items [][]byte) (response []byte) {
	addSymbol('{', &response)
	for i, item := range items {
		addKey(names[i], &response)
		response = append(response, item...)
		if i != len(items)-1 {
			addSymbol(',', &response)
		}
	}
	addSymbol('}', &response)
	return
}

func addSymbol(sym rune, response *[]byte) {
	*response = append(*response, byte(sym))
}

func addKey(line string, response *[]byte) {
	addSymbol('"', response)
	*response = append(*response, []byte(line)...)
	addSymbol('"', response)
	addSymbol(':', response)
}
