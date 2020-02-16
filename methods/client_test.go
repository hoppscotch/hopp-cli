package methods

import (
	"testing"
	"time"
)

func Test_getHTTPClient(t *testing.T) {
	t.Run("default http client", func(t *testing.T) {
		client := getHTTPClient()
		if client == nil {
			t.Error("got nil http client")
		}
		dur := 10 * time.Second
		if client.Timeout != dur {
			t.Errorf("timeout doesn't match what's expected: %v", client.Timeout)
		}
	})
}
