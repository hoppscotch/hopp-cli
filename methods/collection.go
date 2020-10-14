package methods

//Collection hold the structure of the basic `postwoman-collection.json`
type Collection struct {
	// Name of the Whole Collection
	Name string `json:"name"`
	// Folders JSON Type
	Folders []Folders `json:"folders"`
	// Requests inside the Collection
	Requests []Requests `json:"requests"`
}

// Folders can be organized to Folders
type Folders struct {
	// Folder name
	Name string `json:"name"`
	// Requests inside the Folder
	Requests []Requests `json:"requests"`
}

// Headers are the Request Headers
type Headers struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Requests are the Request Model in JSON
type Requests struct {
	// Base URL of the Request
	URL string `json:"url"`
	// Path is the enpoint path
	// URL+PATH = Full URL
	Path string `json:"path"`
	// Request Method - GET,POST,PUT,PATCH,DELETE
	Method string `json:"method"`
	// Authentication Type - Bearer Token or Basic Auth
	Auth string `json:"auth"`
	// Username for Basic Auth
	User string `json:"httpUser"`
	// Password for Basic Auth
	Pass              string `json:"httpPassword"`
	PasswordFieldType string `json:"passwordFieldType"`
	// Bearer token
	Token string `json:"bearerToken"`
	// Request Headers if any- Key,Value pairs
	Headers []Headers `json:"headers"`
	// Params for Get Requests
	Params []interface{} `json:"params"`
	// Body Params for POST requests and forth
	Bparams []interface{} `json:"bodyParams"`
	// Raw Input. Not Formatted JSON
	RawParams string `json:"rawParams"`
	// If RawInputs are used or Not
	RawInput bool `json:"rawInput"`
	// Content Type of Request
	Ctype            string `json:"contentType"`
	RequestType      string `json:"requestType"`
	PreRequestScript string `json:"preRequestScript"`
	TestScript       string `json:"testScript"`
	// Label of Collection
	Label string `json:"label"`
	// Name of the Request
	Name string `json:"name"`
	// Number of Collection
	Collection int `json:"collection"`
}

// BodyParams include the Body Parameters
type BodyParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
