package methods

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/urfave/cli"
)

//Deletebasic sends a basic DELETE request
func Deletebasic(c *cli.Context) (string, error) {
	url, err := checkURL(c.Args().Get(0))
	if err != nil {
		return "", fmt.Errorf("URL validation error: %s", err.Error())
	}

	var jsonStr = []byte(c.String("body"))
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
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
		return "", fmt.Errorf("Request error: %s", err.Error())
	}
	defer resp.Body.Close()

	return formatresp(resp)
}
