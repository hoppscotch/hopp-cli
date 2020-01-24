package methods

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/cli"
)

//Basicreq sends a simple GET request to the url with any potential parameters like Tokens or Basic Auth
func Basicreq(c *cli.Context) error {
	var url = c.String("url")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
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
	fmt.Print(req.Header)
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
