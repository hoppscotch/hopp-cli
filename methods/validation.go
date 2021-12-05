package methods

import (
	"errors"
	"net/url"
	"strings"
)

func checkURL(urlStr string) (string, error) {
	if urlStr == "" {
		return "", errors.New("URL is needed")
	}

	prefixCheck := strings.HasPrefix(urlStr, "http://") || strings.HasPrefix(urlStr, "https://")
	if !prefixCheck {
		return "", errors.New("URL missing protocol or contains invalid protocol")
	}

	if _, err := url.Parse(urlStr); err != nil {
		return "", errors.New("URL is invalid")
	}

	return urlStr, nil
}
