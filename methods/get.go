package methods

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/urfave/cli"
)

// Getbasic sends a simple GET request to the url with any potential parameters like Tokens or Basic Auth
func Getbasic(c *cli.Context) (string, error) {
	var url, err = checkURL(c.Args().Get(0))
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	if c.String("token") != "" {
		var bearer = "Bearer " + c.String("token")
		req.Header.Add("Authorization", bearer)
	}
	if c.String("u") != "" && c.String("p") != "" {
		un := c.String("u")
		pw := c.String("p")
		req.Header.Add("Authorization", "Basic "+basicAuth(un, pw))
	}

	for _, h := range c.StringSlice("header") {
		kv := strings.Split(h, ": ")
		req.Header.Add(kv[0], kv[1])
	}

	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	return formatresp(resp)
}
