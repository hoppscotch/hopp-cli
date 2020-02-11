package methods

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
)

// Formatresp formats the Response with Indents and Colors
func formatresp(resp *http.Response) string {
	c := color.New(color.FgCyan, color.Bold)
	magenta := color.New(color.FgHiMagenta)
	yellow := color.New(color.FgHiYellow)
	body, err := ioutil.ReadAll(resp.Body)
	str := string(body)
	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)
	f := colorjson.NewFormatter()
	f.Indent = 6
	s, _ := f.Marshal(obj)
	for key, value := range resp.Header {
		c.Print(key, " : ")
		magenta.Print(value, "\n")
	}
	retbody := yellow.Sprintf("\nStatus:\t\t%s\n\nStatusCode:\t%d\n", resp.Status, resp.StatusCode) + fmt.Sprintf("\n%s\n", string(s))
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	return retbody
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

//Contenttypes can be used in Place for ctypes
var Contenttypes = map[string]string{
	"html":  "text/html",
	"js":    "application/json",
	"xml":   "application/xml",
	"plain": "text/plain",
}
