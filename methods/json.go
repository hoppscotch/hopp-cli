package methods

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli"
)

//Colls hold the format of the basic `postwoman-collection.json`
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
func ReadCollection(c *cli.Context) {
	data, err := ioutil.ReadFile(c.String("pt"))
	if err != nil {
		fmt.Print(err)
	}
	//fmt.Print(string(data))
	jsondat := []Colls{}

	err = json.Unmarshal([]byte(data), &jsondat)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(jsondat[0].Request); i++ {
		fmt.Printf(`
		URL: %s%s 
		Method: %s
		Auth: %s
		-------`, jsondat[0].Request[i].URL, jsondat[0].Request[i].Path, jsondat[0].Request[i].Method, jsondat[0].Request[i].Auth)
		request(jsondat, i)
	}

}
func request(c []Colls, i int) {
	if c[0].Request[i].Method == "GET" || c[0].Request[i].Method == "POST" {
		//fmt.Print(c[0].Request[i].Method)
		out, err := getsend(c, i, c[0].Request[i].Method)
		if err != nil {
			fmt.Print(err)
		}

		fmt.Printf("%s  %d", out, i)
	}

}
