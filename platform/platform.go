package platform

func IsOSSupported(osName string) bool {
	if osName == "macOS" {
		return true
	}
	return false
}
