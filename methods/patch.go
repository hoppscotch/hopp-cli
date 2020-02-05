package methods

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

//Patchbasic sends a basic PATCH request
func Patchbasic(c *cli.Context) {
	url := c.Args().Get(0)
	if url == "" {
		fmt.Print("URL is needed")
		os.Exit(0)
	}
	var jsonStr = []byte(c.String("body"))
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))
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
