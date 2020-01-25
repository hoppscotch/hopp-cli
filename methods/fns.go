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
	c := color.New(color.FgCyan)
	body, err := ioutil.ReadAll(resp.Body)
	str := string(body)
	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)
	f := colorjson.NewFormatter()
	f.Indent = 6
	s, _ := f.Marshal(obj)
	for key, value := range resp.Header {
		c.Print(key, " : ", value, "\n")
	}
	retbody := fmt.Sprintf("\nStatus:\t\t%s\n\nStatusCode:\t%d\n\n%s\n", resp.Status, resp.StatusCode, string(s))
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	return retbody
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
