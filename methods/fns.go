package methods

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/TylerBrock/colorjson"
)

// Formatresp formats the Response with Indents and Colors
func Formatresp(resp *http.Response) string {
	body, err := ioutil.ReadAll(resp.Body)
	str := string(body)
	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)
	f := colorjson.NewFormatter()
	f.Indent = 6
	s, _ := f.Marshal(obj)
	retbody := fmt.Sprintf("\nStatus:\t\t%s\n\nStatusCode:\t%d\n\n%s\n", resp.Status, resp.StatusCode, string(s))
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	return retbody
}
