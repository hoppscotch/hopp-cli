package methods

import (
	"net/http"
	"time"
)

// The HTTP client can be modified via params and cli flags later
func getHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}
