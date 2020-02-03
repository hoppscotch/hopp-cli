package main

import (
	"fmt"
	"log"
	"os"

	mets "github.com/athul/pwcli/methods"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = color.HiGreenString("Postwoman CLI")
	app.Version = color.HiRedString("0.0.2")
	app.Usage = color.HiYellowString("Test API endpoints without the hassle")
	app.Description = color.HiBlueString("Made with <3 by Postwoman Team")

	getFlags := []cli.Flag{
		cli.StringFlag{
			Name:     "url",
			Value:    "https://reqres.in/api/users",
			Usage:    "The URL/Endpoint you want to check",
			Required: true,
		},
		cli.StringFlag{
			Name:  "token, t",
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
			Name:  "token, t",
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
			Name:  "ctype, c",
			Value: "application/json",
			Usage: "Change the Content Type",
		},
		cli.StringFlag{
			Name:  "body, b",
			Usage: "Body of the Post Request",
		},
	}
	sendFlag := []cli.Flag{
		cli.StringFlag{
			Name:     "pt",
			Usage:    "The Path of Postwoman Collection.json",
			Required: true,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "get",
			Usage: "Send a GET request",
			Flags: getFlags,
			Action: func(c *cli.Context) error {
				mets.Getbasic(c)
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
		{
			Name:  "put",
			Usage: "Send a PUT Request",
			Flags: postFlags,
			Action: func(c *cli.Context) error {
				mets.Putbasic(c)
				return nil
			},
		},
		{
			Name:  "patch",
			Usage: "Send a PATCH Request",
			Flags: postFlags,
			Action: func(c *cli.Context) error {
				mets.Patchbasic(c)
				return nil
			},
		},
		{
			Name:  "delete",
			Usage: "Send a DELETE Request",
			Flags: postFlags,
			Action: func(c *cli.Context) error {
				mets.Deletebasic(c)
				return nil
			},
		},
		{
			Name:  "send",
			Usage: "Test all the Endpoints in the Postwoman Collection.json",
			Flags: sendFlag,
			Action: func(c *cli.Context) error {
				mets.ReadCollection(c)
				return nil
			},
		},
	}
	cli.AppHelpTemplate = fmt.Sprintf(`%sS

	WE REALLY NEED YOUR FEEDBACK, 

	CREATE A NEW ISSUE FOR BUGS AND FEATURE REQUESTS : <http://github.com/athul/pwcli>
	`, cli.AppHelpTemplate)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
