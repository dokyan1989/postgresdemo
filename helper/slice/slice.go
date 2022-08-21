package slice

// ContainsString check slice strings include element
func ContainsString(slice []string, element string) bool {
	if len(slice) == 0 {
		return false
	}
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}
