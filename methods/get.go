package methods

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/TylerBrock/colorjson"
	"github.com/urfave/cli"
)

//Getreq sends a simple GET request to the url
func Getreq(c *cli.Context) error {
	var url = c.String("url")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	s := formatresp(body)
	//log.Println()
	fmt.Print(resp)
	fmt.Printf("\n\n %s", s)
	return nil
}

//Getwtoken send a get request with the Token for Auth
func Getwtoken(c *cli.Context) error {
	var url = c.String("url")
	var bearer = "Bearer " + c.String("token")
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	s := formatresp(body)
	//log.Println()
	fmt.Print(resp)
	fmt.Printf("\n\n %s", s)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	return nil
}
func formatresp(body []byte) string {
	str := string(body)
	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)
	f := colorjson.NewFormatter()
	f.Indent = 6
	s, _ := f.Marshal(obj)
	return string(s)
}
