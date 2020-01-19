package main

import (
	mets "github.com/athul/pwcli/methods"
	"github.com/urfave/cli"
	"log"
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
		cli.StringFlag{
			Name:  "token",
			Value: "Bearer Token",
		},
	}
	app.Commands = []cli.Command{{
		Name:  "get",
		Usage: "Send a GET request",
		Flags: myFlags,
		Action: func(c *cli.Context) error {
			if c.String("token") != "" {
				mets.Getwtoken(c)
			} else {
				mets.Getreq(c)
			}
			return nil
		},
	},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
