package methods

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

// BasicRequestWithBody sends put|patch|post|delete requests
func BasicRequestWithBody(c *cli.Context, method string) (string, error) {
	url, err := checkURL(c.Args().Get(0))
	if err != nil {
		return "", err
	}

	// Check if we're being passed a request body from stdin.
	// If so, use that. Otherwise, use the request data passed via cli flag.
	var body []byte
	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", fmt.Errorf("error getting file info for stdin fd: %w", err)
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		body, err = io.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("error reading from stdin: %w", err)
		}
	} else {
		body = []byte(c.String("body"))
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("Error creating request: %s", err.Error())
	}

	req.Header.Set("Content-Type", Contenttypes[c.String("ctype")])
	if c.String("token") != "" {
		var bearer = "Bearer " + c.String("token")
		req.Header.Add("Authorization", bearer)
	}
	if c.String("u") != "" && c.String("p") != "" {
		un := c.String("u")
		pw := c.String("p")
		req.Header.Add("Authorization", "Basic "+basicAuth(un, pw))
	}

	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %s", err.Error())
	}
	defer resp.Body.Close()

	return formatresp(resp)
}
