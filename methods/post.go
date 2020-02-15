package methods

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/urfave/cli"
)

//Postbasic sends a basic POST request
func Postbasic(c *cli.Context) (string, error) {
	url, err := checkURL(c.Args().Get(0))
	if err != nil {
		return "", err
	}
	var jsonStr = []byte(c.String("body"))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", Contenttypes[c.String("ctype")])
	fmt.Print(Contenttypes[c.String("ctype")] + "\n")
	if c.String("token") != "" {
		var bearer = "Bearer " + c.String("token")
		req.Header.Add("Authorization", bearer)
	}
	if c.String("u") != "" && c.String("p") != "" {
		un := c.String("u")
		pw := c.String("p")
		req.Header.Add("Authorization", "Basic "+basicAuth(un, pw))
	}
	fmt.Print(req.Header)
	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %s", err.Error())
	}
	defer resp.Body.Close()

	return formatresp(resp)
}
