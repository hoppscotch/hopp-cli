package methods

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli"
)

//Getbasic sends a simple GET request to the url with any potential parameters like Tokens or Basic Auth
func Getbasic(c *cli.Context) (string, error) {
	var url, err = checkURL(c.Args().Get(0))
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("Error creating request: %s", err.Error())
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
	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %s", err.Error())
	}
	defer resp.Body.Close()

	return formatresp(resp)
}
