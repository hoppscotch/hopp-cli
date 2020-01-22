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
			Usage:    "The URL/Endpoint you want to check",
			Required: true,
		},
		cli.StringFlag{
			Name:  "token",
			Value: "Bearer Token",
			Usage: "Send the Request with Bearer Token",
		},
		cli.StringFlag{
			Name:     "u",
			Value:    "Username",
			Usage:    "Add the Username",
			Required: true,
		},
		cli.StringFlag{
			Name:     "p",
			Value:    "Password",
			Usage:    "Add the Password",
			Required: true,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "get",
			Usage: "Send a GET request",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				if c.String("u") != "" && c.String("p") != "" {
					mets.Getbasic(c)
				} else if c.String("token") != "" {
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
