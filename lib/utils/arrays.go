package utils

// Check if an array contains a specific string	
func ArrayContains(haystack []string, needle string) bool {
	for _, str := range haystack {
		if str == needle {
			return true
		}
	}

	return false
}