package utils

func InStr(elem string, array []string) bool {
	for _, a := range array {
		if elem == a {
			return true
		}
	}
	return false
}
