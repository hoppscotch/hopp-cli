package methods

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/urfave/cli"
)

//Deletebasic sends a basic DELETE request
func Deletebasic(c *cli.Context) error {
	url, err := checkURL(c.Args().Get(0))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil
	}
	var jsonStr = []byte(c.String("body"))
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
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
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//defer resp.Body.Close()
	s := formatresp(resp)
	fmt.Println(s)
	return nil
}
