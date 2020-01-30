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

func main() {
	data, err := ioutil.ReadFile("/Users/athul/Downloads/pwclc.json")
	if err != nil {
		fmt.Print(err)
	}
	//fmt.Print(string(data))
	jsondat := []Colls{}

	err = json.Unmarshal([]byte(data), &jsondat)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(jsondat)
	//fmt.Print(jsondat[0].Name + "\n")
	//fmt.Println(jsondat[0].Request[4].URL)
	//fmt.Println(jsondat[0].Request[0].Method)
	for i := 0; i < len(jsondat[0].Request); i++ {
		fmt.Printf(`
		URL: %s 
		Method: %s
		Auth: %s
		-------`, jsondat[0].Request[i].URL, jsondat[0].Request[i].Method, jsondat[0].Request[i].Auth)
	}

}
func request(){
	
}