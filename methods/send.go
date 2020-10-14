package methods

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"go.uber.org/multierr"
)

//ReadCollection reads the `hoppScotch-collection.json` File and returns a the Loaded Collection Struct
func ReadCollection(filename string) ([]Collection, error) {
	data, err := ioutil.ReadFile(filename)
	if string(data) == "" {
		return nil, errors.New("PATH is needed")
	}
	if err != nil {
		return nil, err
	}

	var jsonArr []Collection
	err = json.Unmarshal([]byte(data), &jsonArr) // Unmarshal JSON to Collection Type
	if err != nil {
		return nil, fmt.Errorf("Error parsing JSON: %s", err.Error())
	}
	return jsonArr, nil
}

// ProcessCollection parses the Collection struct and execute the said requests
func ProcessCollection(jsonArr []Collection) (string, error) {
	var (
		out  string
		errs error
	)
	for _, jsondat := range jsonArr {
		fmt.Printf("\n------------\nName:\t%s\n", color.HiMagentaString(jsondat.Name))
		if len(jsondat.Folders) > 0 {
			jsondat.getDatafromFolders()
		} else {
			request := jsondat.Requests
			for _, req := range request {
				err := jsondat.request(req)
				if err != nil {
					errs = multierr.Append(errs, err)
				}
			}
		}
	}
	return out, errs
}

// request process the Requests type and executes the Requests synchronously
func (c *Collection) request(req Requests) error {
	var (
		colored       = color.New(color.FgHiRed, color.Bold) // Red string color
		fURL          = req.URL + req.Path                   // Full URL
		err           error
		out, paramURL string
	)

	if req.Method == "GET" {
		out, paramURL, err = c.sendGET(req)
	} else {
		out, err = c.sendPOST(req, req.Method)
	}
	if err != nil {
		return err
	}
	if paramURL != "" {
		fmt.Printf("%s |\t%s | %s |\t%s\n", color.HiGreenString(req.Name), colored.Sprintf(paramURL), color.HiYellowString(req.Method), out)
	} else {
		fmt.Printf("%s |\t%s | %s |\t%s\n", color.HiGreenString(req.Name), colored.Sprintf(fURL), color.HiYellowString(req.Method), out)
	}

	return nil
}

// sendGet sends a GET request to the said URL and returns a Response string and if it contains a
// params in the URL
func (c *Collection) sendGET(req Requests) (string, string, error) {
	var (
		url      = req.URL + req.Path
		bearer   string
		paramstr string = "?"
	)
	if len(req.Params) > 0 {
		for _, param := range req.Params {
			if k, v := param.(map[string]interface{}); v {
				if k["type"] == "query" {
					paramstr += fmt.Sprintf("%v=%v&", k["key"], k["value"])
				}

			}
		}
		paramstr = strings.TrimSuffix(paramstr, "&") // Trim any `&` from the Params
		url += paramstr
	}
	reqHTTP, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", "", fmt.Errorf("Error creating request: %s", err.Error())
	}

	if req.Token != "" {
		// Token Auth
		bearer = "Bearer " + req.Token
		reqHTTP.Header.Add("Authorization", bearer)
	}
	if req.User != "" && req.Pass != "" {
		// Basic Auth
		// basicAuth function encodes the username-password combo to base64 encoded string
		reqHTTP.Header.Add("Authorization", "Basic "+basicAuth(req.User, req.Pass))
	}

	client := getHTTPClient()
	resp, err := client.Do(reqHTTP)
	if err != nil {
		return "", "", fmt.Errorf("Error sending request: %s", err.Error())
	}
	defer resp.Body.Close()
	// paramstr is suffixed with a `?` to append to the URL
	if paramstr != "?" {
		s := fmt.Sprintf("Status: %s\tStatusCode:\t%s \n", color.HiBlueString(resp.Status), color.New(color.FgHiBlue).Sprintln(resp.StatusCode))
		return s, url, nil
	}
	s := fmt.Sprintf("Status: %s\tStatusCode:\t%s\n", color.HiBlueString(resp.Status), color.New(color.FgHiBlue).Sprintln(resp.StatusCode))
	return s, "", nil
}

// sendPOST sends all request other than GET requests since
// it's the only one which has BodyParams and RawParams
func (c *Collection) sendPOST(req Requests, method string) (string, error) {

	var (
		// colorcy = color.New(color.FgCyan, color.Bold)
		jsonStr []byte
		url     = req.URL + req.Path
		reqData = req
		Bpar    = reqData.Bparams
	)
	switch {
	case reqData.RawInput:
		// Check if RawInput is enabled and convert it to bytes
		jsonStr = []byte(reqData.RawParams)
	case len(Bpar) > 0:
		// var dataAsJSON string
		for _, s := range Bpar {
			// Since bodyParams is a map[string]interface{}
			// had to loop over them and get the values
			if k, v := s.(map[string]interface{}); v {
				// Update jsonStr with data from BodyParams.
				jsonStr = append(jsonStr, fmt.Sprintf("\"%s\":\"%v\",", k["key"], k["value"])...)
			}
		}
		dataAsJSON := strings.TrimSuffix(string(jsonStr), ",") // Removes any trailing commas from the Input
		jsonStr = []byte(fmt.Sprintf("{%s}", dataAsJSON))      // Appends a `{` and `}` to front and last of jsonStr to make it a json-like string
	}

	finalBytes, err := json.RawMessage(jsonStr).MarshalJSON() // Marshal to JSON from strings
	if err != nil {
		return "", fmt.Errorf("Error Marhsaling JSON: %s", err.Error())
	}
	reqHTTP, err := http.NewRequest(method, url, bytes.NewBuffer(finalBytes))
	if err != nil {
		return "", fmt.Errorf("Error creating request: %s", err.Error())
	}

	reqHTTP.Header.Set("Content-Type", req.Ctype) // Set Content type to said Ctype in Collection
	if req.Token != "" {
		// Bearer Auth / Token based Auth
		var bearer = "Bearer " + req.Token
		reqHTTP.Header.Add("Authorization", bearer)
	}
	if req.User != "" && req.Pass != "" {
		// Basic Auth
		un := req.User
		pw := req.Pass
		reqHTTP.Header.Add("Authorization", "Basic "+basicAuth(un, pw))
	}

	client := getHTTPClient()
	resp, err := client.Do(reqHTTP)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %s", err.Error())
	}

	defer resp.Body.Close()

	s := fmt.Sprintf("Status: %s\tStatusCode:\t%s\n", color.HiBlueString(resp.Status), color.New(color.FgHiBlue).Sprintln(resp.StatusCode))
	return s, nil
}

// getDatafromFolders handle edge cases when requests are saved inside folders
// from hoppscotch itself
// This adds another loop to check for folders and requests
func (c *Collection) getDatafromFolders() error {
	var (
		err           error
		out, paramURL string
	)
	for _, Folder := range c.Folders {
		for j := range Folder.Requests {
			fmt.Printf("Folder Name:\t%s\n", color.HiMagentaString(Folder.Name))
			fURL := fmt.Sprintf(Folder.Requests[j].URL + Folder.Requests[j].Path)
			method := Folder.Requests[j].Method
			if method == "GET" {
				out, paramURL, err = c.sendGET(Folder.Requests[j])
			} else {
				out, err = c.sendPOST(Folder.Requests[j], method)
			}
			if err != nil {
				return err
			}
			if paramURL != "" {
				fmt.Printf("%s |\t%s | %s |\t%s\n", color.HiGreenString(Folder.Requests[j].Name), color.HiRedString(paramURL), color.HiYellowString(method), out)
			} else {
				fmt.Printf("%s |\t%s | %s |\t%s\n", color.HiGreenString(Folder.Requests[j].Name), color.HiRedString(fURL), color.HiYellowString(method), out)
			}
		}
	}
	return nil
}
