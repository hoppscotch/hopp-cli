package methods

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/tidwall/pretty"
	"github.com/yosssi/gohtml"
)

// Formatresp formats the Response with Indents and Colors
func formatresp(resp *http.Response) (string, error) {
	var retbody string
	heads := fmt.Sprint(resp.Header)
	c := color.New(color.FgCyan, color.Bold)
	magenta := color.New(color.FgHiMagenta)
	yellow := color.New(color.FgHiYellow)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body: %s", err.Error())
	}

	str := string(body)
	if strings.Contains(heads, "json") {
		for key, value := range resp.Header {
			c.Print(key, " : ")
			magenta.Print(value, "\n")
		}
		retbody = yellow.Sprintf("\nStatus:\t\t%s\n\nStatusCode:\t%d\n", resp.Status, resp.StatusCode) + fmt.Sprintf("\n%s\n", string(pretty.Color(pretty.Pretty(body), nil)))
	} else if strings.Contains(heads, "xml") || strings.Contains(heads, "html") || strings.Contains(heads, "plain") {
		for key, value := range resp.Header {
			c.Print(key, " : ")
			magenta.Print(value, "\n")
		}
		var s string
		if strings.Contains(heads, "plain") {
			s = str
		} else {
			s = c.Sprint(gohtml.Format(str))
		}
		retbody = yellow.Sprintf("\nStatus:\t\t%s\n\nStatusCode:\t%d\n", resp.Status, resp.StatusCode) + fmt.Sprintf("\n%s\n", s)
	}
	return retbody, nil
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
