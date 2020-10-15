package methods

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
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
		out     string
		errs    error
		tabData [][]string
	)
	for _, jsondat := range jsonArr {
		fmt.Printf("\nName:\t%s\n", color.HiMagentaString(jsondat.Name))
		if len(jsondat.Folders) > 0 {
			if err := jsondat.getDatafromFolders(); err != nil {
				return "", err
			}
		} else {
			request := jsondat.Requests
			for _, req := range request {

				tabdata, err := jsondat.request(req)
				tabData = append(tabData, tabdata)
				if err != nil {
					errs = multierr.Append(errs, err)
				}
			}
			genTables(tabData)
		}
	}
	return out, errs
}

// request process the Requests type and executes the Requests synchronously
func (c *Collection) request(req Requests) ([]string, error) {
	var (
		colored                = color.New(color.FgHiRed, color.Bold) // Red string color
		fURL                   = req.URL + req.Path                   // Full URL
		err                    error
		paramURL, status, code string
		tabData                []string
	)

	if req.Method == "GET" {
		status, code, paramURL, err = c.sendGET(req)
	} else {
		status, code, err = c.sendPOST(req, req.Method)
	}
	if err != nil {
		return nil, err
	}
	if paramURL != "" {
		tabData = []string{color.HiGreenString(req.Name), colored.Sprintf(paramURL), color.HiYellowString(req.Method), status, code}
	} else {
		tabData = []string{color.HiGreenString(req.Name), colored.Sprintf(fURL), color.HiYellowString(req.Method), status, code}
	}
	return tabData, nil
}

// sendGet sends a GET request to the said URL and returns a Response string and if it contains a
// params in the URL
func (c *Collection) sendGET(req Requests) (string, string, string, error) {
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
	if len(req.Headers) > 0 {
		for _, head := range req.Headers {
			reqHTTP.Header.Set(head.Key, head.Value)
		}
	}
	if err != nil {
		return "", "", "", fmt.Errorf("Error creating request: %s", err.Error())
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
		return "", "", "", fmt.Errorf("Error sending request: %s", err.Error())
	}
	defer resp.Body.Close()
	// paramstr is suffixed with a `?` to append to the URL
	if paramstr != "?" {
		status := fmt.Sprintf("%s", color.HiBlueString(resp.Status))
		code := fmt.Sprintf("%s", color.New(color.FgHiBlue).Sprintln(resp.StatusCode))
		return status, code, url, nil
	}
	status := fmt.Sprintf("%s", color.HiBlueString(resp.Status))
	code := fmt.Sprintf("%s", color.New(color.FgHiBlue).Sprintln(resp.StatusCode))
	return status, code, "", nil
}

// sendPOST sends all request other than GET requests since
// it's the only one which has BodyParams and RawParams
func (c *Collection) sendPOST(req Requests, method string) (string, string, error) {

	var (
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
		return "", "", fmt.Errorf("Error Marhsaling JSON: %s", err.Error())
	}
	reqHTTP, err := http.NewRequest(method, url, bytes.NewBuffer(finalBytes))
	if len(req.Headers) > 0 {
		for _, head := range req.Headers {
			reqHTTP.Header.Set(head.Key, head.Value)
		}
	}
	if err != nil {
		return "", "", fmt.Errorf("Error creating request: %s", err.Error())
	}

	reqHTTP.Header.Set("Content-Type", req.Ctype) // Set Content type to said Ctype in Collection
	if req.Token != "" {
		// Bearer Auth / Token based Auth
		var bearer = "Bearer " + req.Token
		reqHTTP.Header.Add("Authorization", bearer)
	}
	if req.User != "" && req.Pass != "" {
		// Basic Auth
		reqHTTP.Header.Add("Authorization", "Basic "+basicAuth(req.User, req.Pass))
	}

	client := getHTTPClient()
	resp, err := client.Do(reqHTTP)
	if err != nil {
		return "", "", fmt.Errorf("Error sending request: %s", err.Error())
	}

	defer resp.Body.Close()

	status := fmt.Sprintf("%s", color.HiBlueString(resp.Status))
	code := fmt.Sprintf("%s", color.New(color.FgHiBlue).Sprintln(resp.StatusCode))
	return status, code, nil
}

// getDatafromFolders handle edge cases when requests are saved inside folders
// from hoppscotch itself
// This adds another loop to check for folders and requests
func (c *Collection) getDatafromFolders() error {
	var (
		err                    error
		paramURL, status, code string
		tabData                [][]string
	)
	for _, Folder := range c.Folders {
		for _, fRequest := range Folder.Requests {
			fURL := fmt.Sprintf(fRequest.URL + fRequest.Path)
			method := fRequest.Method
			if method == "GET" {
				status, code, paramURL, err = c.sendGET(fRequest)
			} else {
				status, code, err = c.sendPOST(fRequest, method)
			}
			if err != nil {
				log.Println(err)
			}
			if paramURL != "" {
				tabData = append(tabData, []string{color.HiGreenString(fRequest.Name), color.HiRedString(paramURL), color.HiYellowString(method), status, code})
			} else {
				tabData = append(tabData, []string{color.HiGreenString(fRequest.Name), color.HiRedString(fURL), color.HiYellowString(method), status, code})
			}
		}
	}
	genTables(tabData)
	return nil
}

//genTables generate the output in Tabular Form
func genTables(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "URL", "Method", "Status", "Code"})
	table.AppendBulk(data)
	table.Render()
}
