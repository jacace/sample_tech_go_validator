package platform

import "errors"

func IsOSSupported(osName string) (bool, error) {
	if osName == "" {
		return false, errors.New("osName can't be empty or null")
	}

	if osName == "macOS" {
		return true, nil
	}
	return false, nil
}
