package methods

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"

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
		var path = os.Getenv("EDITOR")
		var editor string

		if path == "" {
			// check OS
			currentOs := runtime.GOOS

			// based on OS find assign default editor
			if currentOs == "linux" || currentOs == "darwin" {
				editor = "nano"
			} else if currentOs == "windows" {
				editor = "notepad"
			}

			path, err = exec.LookPath(editor)
			if err != nil {
				return "", fmt.Errorf("Unable to locate %s: %w", editor, err)
			}
		}

		// create temp file for writing request body
		tempFile, err := ioutil.TempFile("", "HOPP-CLI-REQUEST*.txt")
		if err != nil {
			return "", fmt.Errorf("Unable to create temp file: %w", err)
		}

		tempFileName := string(tempFile.Name())
		defer os.Remove(tempFileName)
		tempFile.Close()

		// open temp file in selected editor and wait until closed
		cmd := exec.Command(path, tempFileName)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		err = cmd.Start()
		if err != nil {
			return "", fmt.Errorf("Unable to open editor: %w", err)
		}

		fmt.Println("Waiting for file to close..")

		err = cmd.Wait()
		if err != nil {
			return "", fmt.Errorf("Error waiting for file: %w", err)
		}

		// read temp file contents
		jsonStr, err = ioutil.ReadFile(tempFileName)
		if err != nil {
			return "", fmt.Errorf("Error reading file %s: %w", tempFileName, err)
		}
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
