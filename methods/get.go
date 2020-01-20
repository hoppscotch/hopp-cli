package methods

import (
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

//Getwtoken send a get request with the Token for Auth
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
	fmt.Print(resp)
	fmt.Printf("\n\n %s", s)

	return nil
}

//Getbasic send a request with Baic Auth
func Getbasic(c *cli.Context) error {
	return nil
}

// Dummy Code
/* func basicAuth() string {
    var username string = "foo"
    var passwd string = "bar"
    client := &http.Client{}
    req, err := http.NewRequest("GET", "mydomain.com", nil)
    req.SetBasicAuth(username, passwd)
    resp, err := client.Do(req)
    if err != nil{
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)
    s := string(bodyText)
    return s
} */
