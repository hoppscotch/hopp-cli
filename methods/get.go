package methods

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

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
	defer resp.Body.Close()
	s := Formatresp(resp)
	//log.Println()
	fmt.Print(resp)
	fmt.Printf("\n\n %s", s)
	return nil
}

//Getwtoken send a get request with the Token for Authorization Header
func Getwtoken(c *cli.Context) error {
	var url = c.String("url")
	var bearer = "Bearer " + c.String("token")
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	defer resp.Body.Close()
	s := Formatresp(resp)
	if s != "" {
		fmt.Printf("%s", s)
	} else {
		fmt.Print(resp)
	}
	return nil
}

//Getbasic helps you send a request with Basic Auth as Authorization Method
func Getbasic(c *cli.Context) error {
	un := c.String("u")
	pw := c.String("p")
	url := c.String("url")
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(un, pw))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	defer resp.Body.Close()
	s := Formatresp(resp)
	if s != "" {
		fmt.Printf("%s", s)
	} else {
		fmt.Print(resp)
	}

	return nil
}
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
