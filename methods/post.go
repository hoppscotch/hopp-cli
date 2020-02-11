package methods

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

//Postbasic sends a basic POST request
func Postbasic(c *cli.Context) {
	url, err := checkURL(c.Args().Get(0))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(0)
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
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//defer resp.Body.Close()
	s := formatresp(resp)
	fmt.Println(s)
}
