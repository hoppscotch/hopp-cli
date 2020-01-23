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
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", c.String("ctype"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s := Formatresp(resp)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", s)
}
