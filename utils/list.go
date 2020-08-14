package utils

func StringInList(str string, list []string) bool {
	for _, d := range list {
		if d == str {
			return true
		}
	}
	return false
}
