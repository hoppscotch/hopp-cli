package methods

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

// BasicRequestWithBody sends put|patch|post|delete requests
func BasicRequestWithBody(c *cli.Context, method string) (string, error) {
	url, err := checkURL(c.Args().Get(0))
	if err != nil {
		return "", err
	}

	// Check if we're being passed a request body from stdin.
	// If so, use that. Otherwise, use the request data passed via cli flag.
	var body []byte
	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", fmt.Errorf("error getting file info for stdin fd: %w", err)
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		body, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("error reading from stdin: %w", err)
		}
	} else if c.Bool("editor") {
		var path string
		var editor = os.Getenv("EDITOR")

		if editor == "" {
			// check OS
			currentOs := runtime.GOOS

			// based on OS find assign default editor
			if currentOs == "linux" || currentOs == "darwin" {
				editor = "nano"
			} else if currentOs == "windows" {
				editor = "notepad"
			}
		}

		path, err = exec.LookPath(editor)
		if err != nil {
			return "", fmt.Errorf("Unable to locate %s: %w", editor, err)
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

		if err := cmd.Start(); err != nil {
			return "", fmt.Errorf("Unable to open editor: %w", err)
		}

		color.Yellow("Waiting for file to close..\n\n")

		if err := cmd.Wait(); err != nil {
			return "", fmt.Errorf("Error waiting for file: %w", err)
		}

		// read temp file contents
		body, err = ioutil.ReadFile(tempFileName)
		if err != nil {
			return "", fmt.Errorf("Error reading file %s: %w", tempFileName, err)
		}
	} else {
		body = []byte(c.String("body"))
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", Contenttypes[c.String("ctype")])
	if c.String("token") != "" {
		var bearer = "Bearer " + c.String("token")
		req.Header.Add("Authorization", bearer)
	}

	for _, h := range c.StringSlice("header") {
		kv := strings.Split(h, ": ")
		req.Header.Add(kv[0], kv[1])
	}

	if c.String("u") != "" && c.String("p") != "" {
		un := c.String("u")
		pw := c.String("p")
		req.Header.Add("Authorization", "Basic "+basicAuth(un, pw))
	}

	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	return formatresp(resp)
}
