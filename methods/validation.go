package methods

import (
	"fmt"
	"net/url"
	"strings"
)

func checkURL(urlStr string) (string, error) {
	if urlStr == "" {
		return "", fmt.Errorf("URL is needed")
	}

	prefixCheck := strings.HasPrefix(urlStr, "http://") || strings.HasPrefix(urlStr, "https://")
	if !prefixCheck {
		return "", fmt.Errorf("URL missing protocol or contains invalid protocol")
	}

	if _, err := url.Parse(urlStr); err != nil {
		return "", fmt.Errorf("URL is invalid")
	}

	return urlStr, nil
}
