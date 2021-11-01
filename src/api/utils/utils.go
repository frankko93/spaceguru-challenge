package utils

func Contains(list []string, val string) bool {
	for _, value := range list {
		if value == val {
			return true
		}
	}
	return false
}
