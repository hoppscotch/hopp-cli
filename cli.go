package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	mets "github.com/hoppscotch/hopp-cli/methods"
	"github.com/urfave/cli"
)

// VERSION is set by `make` during the build to the most recent tag
var buildVersion = "unknown"

func main() {
	app := cli.NewApp()
	app.Name = color.HiGreenString("Hoppscotch CLI")
	app.Version = color.HiRedString(buildVersion)
	app.Usage = color.HiYellowString("Test API endpoints without the hassle")
	app.Description = color.HiBlueString("Made with <3 by Hoppscotch Team")

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
	genFlags := []cli.Flag{
		cli.IntFlag{
			Name:  "port, p",
			Usage: "Port at which the server will open to",
			Value: 1341,
		},
		cli.BoolFlag{
			Name:  "browser, b",
			Usage: "Whether to open the browser automatically",
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
			Usage: "Test all the Endpoints in the Hoppscotch Collection.json",
			Action: func(c *cli.Context) error {
				coll, err := mets.ReadCollection(c.Args().Get(0))
				if err != nil {
					return err
				}
				_, err = mets.ProcessCollection(coll)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "gen",
			Usage: "Generate Documentation from the Hoppscotch Collection.json",
			Flags: genFlags,
			Action: func(c *cli.Context) error {
				if err := mets.GenerateDocs(c); err != nil {
					return err
				}
				return nil
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
