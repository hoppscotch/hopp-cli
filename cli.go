package main

import (
	"fmt"
	"log"
	"os"

	mets "github.com/athul/pwcli/methods"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

// VERSION is set by `make` during the build to the most recent tag
var VERSION = ""

func main() {
	app := cli.NewApp()
	app.Name = color.HiGreenString("Postwoman CLI")
	app.Version = color.HiRedString(VERSION)
	app.Usage = color.HiYellowString("Test API endpoints without the hassle")
	app.Description = color.HiBlueString("Made with <3 by Postwoman Team")

	var out string

	getFlags := []cli.Flag{
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
			Name:  "ctype, c", //Content Type Flag
			Usage: "Change the Content Type",
		},
		cli.StringFlag{
			Name:  "body, b",
			Usage: "Body of the Post Request",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "get",
			Usage: "Send a GET request",
			Flags: getFlags,
			Action: func(c *cli.Context) error {
				var err error
				out, err = mets.Getbasic(c)
				return err
			},
		},
		{
			Name:  "post",
			Usage: "Send a POST Request",
			Flags: postFlags,
			Action: func(c *cli.Context) error {
				var err error
				out, err = mets.BasicRequestWithBody(c, "POST")
				return err
			},
		},
		{
			Name:  "put",
			Usage: "Send a PUT Request",
			Flags: postFlags,
			Action: func(c *cli.Context) error {
				var err error
				out, err = mets.BasicRequestWithBody(c, "PUT")
				return err
			},
		},
		{
			Name:  "patch",
			Usage: "Send a PATCH Request",
			Flags: postFlags,
			Action: func(c *cli.Context) error {
				var err error
				out, err = mets.BasicRequestWithBody(c, "PATCH")
				return err
			},
		},
		{
			Name:  "delete",
			Usage: "Send a DELETE Request",
			Flags: postFlags,
			Action: func(c *cli.Context) error {
				var err error
				out, err = mets.BasicRequestWithBody(c, "DELETE")
				return err
			},
		},
		{
			Name:  "send",
			Usage: "Test all the Endpoints in the Postwoman Collection.json",
			Action: func(c *cli.Context) error {
				var err error
				out, err = mets.ReadCollection(c)
				return err
			},
		},
	}
	cli.AppHelpTemplate = fmt.Sprintf(`%s

	WE REALLY NEED YOUR FEEDBACK,

	CREATE A NEW ISSUE FOR BUGS AND FEATURE REQUESTS : < https://github.com/hoppscotch/hopp-cli >

	`, cli.AppHelpTemplate)

	err := app.Run(os.Args)
	if err != nil {
		l := log.New(os.Stderr, "", 0)
		l.Println(color.HiRedString("\n%s", err.Error()))
		os.Exit(1)
	}
	fmt.Println(out)
}
