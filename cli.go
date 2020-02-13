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
				return mets.Getbasic(c)
			},
		},
		{
			Name:  "post",
			Usage: "Send a POST Request",
			Flags: postFlags,
			Action: func(c *cli.Context) error {
				return mets.Postbasic(c)
			},
		},
		{
			Name:  "put",
			Usage: "Send a PUT Request",
			Flags: postFlags,
			Action: func(c *cli.Context) error {
				return mets.Putbasic(c)
			},
		},
		{
			Name:  "patch",
			Usage: "Send a PATCH Request",
			Flags: postFlags,
			Action: func(c *cli.Context) error {
				return mets.Patchbasic(c)
			},
		},
		{
			Name:  "delete",
			Usage: "Send a DELETE Request",
			Flags: postFlags,
			Action: func(c *cli.Context) error {
				return mets.Deletebasic(c)
			},
		},
		{
			Name:  "send",
			Usage: "Test all the Endpoints in the Postwoman Collection.json",
			//Flags: sendFlag,
			Action: func(c *cli.Context) error {
				return mets.ReadCollection(c)
			},
		},
	}
	cli.AppHelpTemplate = fmt.Sprintf(`%sS

	WE REALLY NEED YOUR FEEDBACK, 

	CREATE A NEW ISSUE FOR BUGS AND FEATURE REQUESTS : < http://github.com/athul/pwcli >
	`, cli.AppHelpTemplate)

	err := app.Run(os.Args)
	if err != nil {
		l := log.New(os.Stderr, "", 0)
		l.Println(color.HiRedString("\n%s", err.Error()))
	}
}
