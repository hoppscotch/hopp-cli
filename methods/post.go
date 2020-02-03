package methods

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/urfave/cli"
)

//Postbasic sends a basic POST request
func Postbasic(c *cli.Context) {
	url := c.String("url")
	var jsonStr = []byte(c.String("body"))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", c.String("ctype"))
	if c.String("token") != "" {
		var bearer = "Bearer " + c.String("token")
		req.Header.Add("Authorization", bearer)
	}
	if c.String("u") != "" && c.String("p") != "" {
		un := c.String("u")
		pw := c.String("p")
		req.Header.Add("Authorization", "Basic "+basicAuth(un, pw))
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//defer resp.Body.Close()
	s := formatresp(resp)
	fmt.Println(s)
}
