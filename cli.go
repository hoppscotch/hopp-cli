package main

import (
	"log"
	"os"

	mets "github.com/athul/pwcli/methods"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Postwoman CLI"
	app.Usage = "Test API endpoints without the hassle"
	app.Description = "Made with <3 by Postwoman Team"

	getFlags := []cli.Flag{
		cli.StringFlag{
			Name:     "url",
			Value:    "https://reqres.in/api/users",
			Usage:    "The URL/Endpoint you want to check",
			Required: true,
		},
		cli.StringFlag{
			Name:  "token",
			Usage: "Send the Request with Bearer Token",
		},
		cli.StringFlag{
			Name:  "u",
			Usage: "Add the Username",
		},
		cli.StringFlag{
			Name:  "p",
			Usage: "Add the Password",
		},
	}
	postFlags := []cli.Flag{
		cli.StringFlag{
			Name:     "url",
			Value:    "https://reqres.in/api/users",
			Usage:    "The URL/Endpoint you want to check",
			Required: true,
		},
		cli.StringFlag{
			Name:  "token",
			Usage: "Send the Request with Bearer Token",
		},
		cli.StringFlag{
			Name:  "u",
			Usage: "Add the Username",
		},
		cli.StringFlag{
			Name:  "p",
			Usage: "Add the Password",
		},
		cli.StringFlag{
			Name:  "ctype",
			Value: "application/json",
			Usage: "Change the Content Type",
		},
		cli.StringFlag{
			Name:  "body",
			Usage: "Body of the Post Request",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "get",
			Usage: "Send a GET request",
			Flags: getFlags,
			Action: func(c *cli.Context) error {
				switch {
				case c.String("token") != "":
					mets.Authwtoken(c, "GET")
					break
				case c.String("u") != "":
					mets.Authbasic(c, "GET")
				default:
					mets.Basicreq(c)
				}

				return nil
			},
		},
		{
			Name:  "post",
			Usage: "Send a POST Request",
			Flags: postFlags,
			Action: func(c *cli.Context) error {
				mets.Postbasic(c)
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
