package methods

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/urfave/cli"
)

// BasicRequestWithBody sends put|patch|post|delete requests
func BasicRequestWithBody(c *cli.Context, method string) (string, error) {
	url, err := checkURL(c.Args().Get(0))
	if err != nil {
		return "", err
	}

	var jsonStr = []byte(c.String("body"))
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", fmt.Errorf("Error creating request: %s", err.Error())
	}

	req.Header.Set("Content-Type", Contenttypes[c.String("ctype")])
	if c.String("token") != "" {
		var bearer = "Bearer " + c.String("token")
		req.Header.Add("Authorization", bearer)
	}

	for _, h := range c.StringSlice("header") {
		kv := strings.Split(h, ": ")
		req.Header.Add(kv[0], kv[1])
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
