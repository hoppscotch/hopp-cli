package methods

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

//Getbasic sends a simple GET request to the url with any potential parameters like Tokens or Basic Auth
func Getbasic(c *cli.Context) {
	var url = c.Args().Get(0)
	if url == "" {
		fmt.Print("URL is needed")
		os.Exit(0)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error on Request.\nDid you give a correct URL? Try giving that")
		os.Exit(0)
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
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(color.RedString("Error on response.\nDid you give a correct URL? Try giving that"))
	}

	s := formatresp(resp)
	fmt.Println(s)

}
