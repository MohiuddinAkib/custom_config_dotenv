package config

func contains(s [4]string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
