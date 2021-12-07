package methods

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

// BasicRequestWithBody sends put|patch|post|delete requests
func BasicRequestWithBody(c *cli.Context, method string) (string, error) {
	url, err := checkURL(c.Args().Get(0))
	if err != nil {
		return "", err
	}

	var jsonStr []byte

	if c.Bool("editor") {
		var path string
		var editor string

		// check OS
		currentOs := runtime.GOOS

		// based on OS find path for default editor
		if currentOs == "linux" || currentOs == "darwin" {
			editor = "nano"
			path, err = exec.LookPath(editor)
		} else if currentOs == "windows" {
			editor = "notepad"
			path, err = exec.LookPath(editor)
		}

		fmt.Println(string(path))
		if err != nil {
			return "", fmt.Errorf("Error %s while locating %s!!", path, editor)
		}

		// create temp file for writing request body

		tempFile, err := ioutil.TempFile("", "HOPP-CLI-REQUEST*.txt")

		if err != nil {
			return "", fmt.Errorf("Unable to create temp file: %s\n%s", tempFile.Name(), err)
		}

		// open temp file in selected editor and wait until closed
		tempFileName := string(tempFile.Name())
		cmd := exec.Command(path, tempFileName)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		err = cmd.Start()

		if err != nil {
			return "", fmt.Errorf("Unable to open editor: %s", err)
		}

		color.Yellow("Waiting for file to close..\n")
		cmd.Wait()

		// read temp file contents
		fileData, err := ioutil.ReadFile(tempFileName)

		if err != nil {
			return "", fmt.Errorf("Error reading file: %s", err)
		}

		jsonStr = []byte(string(fileData))

		defer os.Remove(tempFileName)
	} else {
		jsonStr = []byte(c.String("body"))
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", fmt.Errorf("Error creating request: %s", err.Error())
	}

	req.Header.Set("Content-Type", Contenttypes[c.String("ctype")])
	if c.String("token") != "" {
		var bearer = "Bearer " + c.String("token")
		req.Header.Add("Authorization", bearer)
	}
	if c.String("u") != "" && c.String("p") != "" {
		un := c.String("u")
		pw := c.String("p")
		req.Header.Add("Authorization", "Basic "+basicAuth(un, pw))
	}

	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %s", err.Error())
	}
	defer resp.Body.Close()

	return formatresp(resp)
}
