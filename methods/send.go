package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

//Colls hold the structure of the basic `postwoman-collection.json`
type Colls struct {
	Name    string    `json:"name"`
	Folders []string  `json:"folders"`
	Request []Reqdata `json:"requests"`
}

//Reqdata hold the format of the request part in `postwoman-collection.json`
type Reqdata struct {
	URL     string     `json:"url"`
	Path    string     `json:"path"`
	Method  string     `json:"method"`
	Auth    string     `json:"auth"`
	User    string     `json:"httpUser"`
	Pass    string     `json:"httpPassword"`
	Token   string     `json:"bearerToken"`
	Ctype   string     `json:"contentType"`
	Name    string     `json:"name"`
	Heads   []string   `json:"headers"`
	Params  []string   `json:"params"`
	Bparams []Bpardata `json:"bodyParams"`
}

//Bpardata hold the format of the bodyParams of `postwoman-collection.json`
type Bpardata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//ReadCollection reads the PostWoman Collection Json File and does the Magic Stuff
func ReadCollection(c *cli.Context) error {
	data, err := ioutil.ReadFile(c.Args().Get(0))
	if string(data) == "" {
		return fmt.Errorf("PATH is needed")
	}
	if err != nil {
		return err
	}
	var jsondat []Colls
	err = json.Unmarshal([]byte(data), &jsondat)
	if err != nil {
		return fmt.Errorf("Error parsing JSON: %s", err.Error())
	}
	fmt.Println("Name:\t" + color.HiMagentaString(jsondat[0].Name))
	for i := 0; i < len(jsondat[0].Request); i++ {
		err := request(jsondat, i)
		if err != nil {
			return err
		}
	}
	return nil
}
func request(c []Colls, i int) error {
	colors := color.New(color.FgHiRed, color.Bold)
	fURL := colors.Sprintf(c[0].Request[i].URL + c[0].Request[i].Path)
	// if err != nil {
	// 	fmt.Printf("%s\n", err.Error())
	// 	return nil
	//}
	if c[0].Request[i].Method == "GET" {
		out, err := getsend(c, i, "GET")
		if err != nil {
			return err
		}
		methods := color.HiYellowString(c[0].Request[i].Method)
		fmt.Printf("%s |\t%s |\t%s |\t%s", color.HiGreenString(c[0].Request[i].Name), fURL, methods, out)
	} else {
		out, err := sendpopa(c, i, c[0].Request[i].Method)
		if err != nil {
			return err
		}
		methods := color.HiYellowString(c[0].Request[i].Method)
		fURL := colors.Sprintf(c[0].Request[i].URL + c[0].Request[i].Path)
		fmt.Printf("%s |\t%s |\t%s |\t%s", color.HiGreenString(c[0].Request[i].Name), fURL, methods, out)
	}
	return nil
}
func getsend(c []Colls, ind int, method string) (string, error) {
	color := color.New(color.FgCyan, color.Bold)
	var url = c[0].Request[ind].URL + c[0].Request[ind].Path
	//fmt.Print(url + "  ")
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", fmt.Errorf("Error creating request: %s", err.Error())
	}
	if c[0].Request[ind].Token != "" {
		var bearer = "Bearer " + c[0].Request[ind].Token
		req.Header.Add("Authorization", bearer)
	}
	if c[0].Request[ind].User != "" && c[0].Request[ind].Pass != "" {
		un := c[0].Request[ind].User
		pw := c[0].Request[ind].Pass
		req.Header.Add("Authorization", "Basic "+basicAuth(un, pw))
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %s", err.Error())
	}
	//defer resp.Body.Close()
	//fmt.Print(resp.Header)
	s := color.Sprintf("Status: %s\tStatusCode:\t%d\n", resp.Status, resp.StatusCode)
	return s, nil
}

func sendpopa(c []Colls, ind int, method string) (string, error) {
	color := color.New(color.FgCyan, color.Bold)
	var jsonStr []byte
	var url = c[0].Request[ind].URL + c[0].Request[ind].Path
	if len(c[0].Request[ind].Bparams) > 0 {
		jsonStr = []byte(string(c[0].Request[ind].Bparams[0].Key[0] + c[0].Request[ind].Bparams[0].Value[0]))
	} else {
		jsonStr = nil
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", fmt.Errorf("Error creating request: %s", err.Error())
	}
	req.Header.Set("Content-Type", c[0].Request[ind].Ctype)
	if c[0].Request[ind].Token != "" {
		var bearer = "Bearer " + c[0].Request[ind].Token
		req.Header.Add("Authorization", bearer)
	}
	if c[0].Request[ind].User != "" && c[0].Request[ind].Pass != "" {
		un := c[0].Request[ind].User
		pw := c[0].Request[ind].Pass
		req.Header.Add("Authorization", "Basic "+basicAuth(un, pw))
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %s", err.Error())
	}
	//defer resp.Body.Close()
	//fmt.Print(resp.Header)
	s := color.Sprintf("Status: %s\tStatusCode:\t%d\n", resp.Status, resp.StatusCode)
	return s, nil

}
