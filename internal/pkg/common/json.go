package common

func UnionJSONAsArray(first []byte, second []byte) []byte {
	first = append(first, []byte(",")...)
	first = append(first, second...)
	return first
}
