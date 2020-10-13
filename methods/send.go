package methods

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli"
	"go.uber.org/multierr"
)

//Colls hold the structure of the basic `postwoman-collection.json`
type Colls struct {
	Name    string     `json:"name"`
	Folders []Folders  `json:"folders"`
	Request []Requests `json:"requests"`
}
type Folders struct {
	Name     string     `json:"name"`
	Requests []Requests `json:"requests"`
}
type Headers struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type Requests struct {
	URL               string        `json:"url"`
	Path              string        `json:"path"`
	Method            string        `json:"method"`
	Auth              string        `json:"auth"`
	User              string        `json:"httpUser"`
	Pass              string        `json:"httpPassword"`
	PasswordFieldType string        `json:"passwordFieldType"`
	Token             string        `json:"bearerToken"`
	Headers           []Headers     `json:"headers"`
	Params            []interface{} `json:"params"`
	Bparams           []interface{} `json:"bodyParams"`
	RawParams         string        `json:"rawParams"`
	RawInput          bool          `json:"rawInput"`
	Ctype             string        `json:"contentType"`
	RequestType       string        `json:"requestType"`
	PreRequestScript  string        `json:"preRequestScript"`
	TestScript        string        `json:"testScript"`
	Label             string        `json:"label"`
	Name              string        `json:"name"`
	Collection        int           `json:"collection"`
}
type BodyParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//ReadCollection reads the PostWoman Collection Json File and does the Magic Stuff
func ReadCollection(c *cli.Context) (string, error) {
	data, err := ioutil.ReadFile(c.Args().Get(0))
	if string(data) == "" {
		return "", errors.New("PATH is needed")
	}
	if err != nil {
		return "", err
	}

	var jsonArr []Colls
	var errs error
	var out string
	err = json.Unmarshal([]byte(data), &jsonArr)
	if err != nil {
		return "", fmt.Errorf("Error parsing JSON: %s", err.Error())
	}
	for i, jsondat := range jsonArr {
		var request []Requests
		out := fmt.Sprintf("Name:\t%s\n", color.HiMagentaString(jsondat.Name))
		if len(jsondat.Folders) > 0 {
			for j := range jsondat.Folders[i].Requests {
				tmpOut, err := jsondat.request(i, true)
				if err != nil {
					errs = multierr.Append(errs, err)
				}
				out += tmpOut
				log.Println(jsondat.Folders[i].Requests[j])

			}
		} else {
			request = jsondat.Request
			// jsondat := jsonArr[i]

			for i := range request {
				log.Println(request[i].Name)
				tmpOut, err := jsondat.request(i, false)
				if err != nil {
					errs = multierr.Append(errs, err)
				}
				out += tmpOut
			}
		}
	}
	log.Println(out)
	return out, errs
}

func (c *Colls) request(i int, folders bool) (string, error) {

	var (
		colors = color.New(color.FgHiRed, color.Bold)
		fURL   string
		method string
		out    string
		err    error
	)
	if folders {
		for j := range c.Folders[i].Requests {
			fURL = colors.Sprintf(c.Folders[i].Requests[j].URL + c.Folders[i].Requests[j].Path)
			method = c.Folders[i].Requests[j].Method
		}
	} else {
		fURL = colors.Sprintf(c.Request[i].URL + c.Request[i].Path)
	}

	if method == "GET" {
		out, err = c.getsend(i, "GET")
	} else {
		out, err = c.sendpopa(i, method)
	}
	if err != nil {
		return "", err
	}

	methods := color.HiYellowString(c.Request[i].Method)
	result := fmt.Sprintf("%s |\t%s |\t%s |\t%s\n", color.HiGreenString(c.Request[i].Name), fURL, methods, out)

	return result, nil
}

func (c *Colls) getsend(ind int, method string) (string, error) {
	var (
		color  = color.New(color.FgCyan, color.Bold)
		url    = c.Request[ind].URL + c.Request[ind].Path
		bearer string
	)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", fmt.Errorf("Error creating request: %s", err.Error())
	}

	if c.Request[ind].Token != "" {
		bearer = "Bearer " + c.Request[ind].Token
		req.Header.Add("Authorization", bearer)
	}
	if c.Request[ind].User != "" && c.Request[ind].Pass != "" {
		un := c.Request[ind].User
		pw := c.Request[ind].Pass
		req.Header.Add("Authorization", "Basic "+basicAuth(un, pw))
	}

	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %s", err.Error())
	}
	defer resp.Body.Close()

	s := color.Sprintf("Status: %s\tStatusCode:\t%d\n", resp.Status, resp.StatusCode)
	return s, nil
}

func (c *Colls) sendpopa(ind int, method string) (string, error) {

	var (
		color   = color.New(color.FgCyan, color.Bold)
		jsonStr []byte
		url     = c.Request[ind].URL + c.Request[ind].Path
		reqData = c.Request[ind]
		Bpar    = reqData.Bparams
	)
	switch {
	case reqData.RawInput:
		jsonStr = []byte(reqData.RawParams)
	case len(Bpar) > 0:
		// var dataAsJSON string
		for _, s := range Bpar {
			if k, v := s.(map[string]interface{}); v {
				log.Println(k["key"], k["value"])
				jsonStr = append(jsonStr, fmt.Sprintf("\"%s\":\"%v\",", k["key"], k["value"])...)
			}
		}
		dataAsJSON := strings.TrimSuffix(string(jsonStr), ",")
		jsonStr = []byte(fmt.Sprintf("{%s}", dataAsJSON))
		// log.Println(string(jsonStr))
	}

	// log.Println(string(jsonStr))
	finalBytes, err := json.RawMessage(jsonStr).MarshalJSON()
	if err != nil {
		return "", fmt.Errorf("Error Marhsaling JSON: %s", err.Error())
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(finalBytes))
	if err != nil {
		return "", fmt.Errorf("Error creating request: %s", err.Error())
	}

	req.Header.Set("Content-Type", c.Request[ind].Ctype)
	if c.Request[ind].Token != "" {
		var bearer = "Bearer " + c.Request[ind].Token
		req.Header.Add("Authorization", bearer)
	}
	if c.Request[ind].User != "" && c.Request[ind].Pass != "" {
		un := c.Request[ind].User
		pw := c.Request[ind].Pass
		req.Header.Add("Authorization", "Basic "+basicAuth(un, pw))
	}

	client := getHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %s", err.Error())
	}
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	s := color.Sprintf("Status: %s\tStatusCode:\t%d\n", resp.Status, resp.StatusCode)
	log.Println(s, string(text))
	return s, nil
}
