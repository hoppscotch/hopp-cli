package main

import (
	"encoding/json"
	"fmt"
	"github.com/TylerBrock/colorjson"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Postwoman CLI"
	app.Usage = "Test API endpoints without the hassle"
	app.Description = "Made with <3 by Postwoman Team"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:     "url",
			Value:    "https://reqres.in/api/users",
			Required: true,
		},
	}
	app.Commands = []cli.Command{{
		Name:  "gt",
		Usage: "Send a GET request",
		Flags: myFlags,
		Action: func(c *cli.Context) error {
			resp, err := http.Get(c.String("url"))
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			str := string(body)
			var obj map[string]interface{}
			json.Unmarshal([]byte(str), &obj)
			f := colorjson.NewFormatter()
			f.Indent = 6
			s, _ := f.Marshal(obj)
			//log.Println()
			fmt.Print(resp)
			fmt.Printf("\n\n %s", string(s))
			return nil
		},
	},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
