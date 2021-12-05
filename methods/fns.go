package methods

import (
	"encoding/base64"
	"fmt"
	"io"
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
	c := color.New(color.FgHiCyan)
	magenta := color.New(color.FgHiMagenta)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	for key, value := range resp.Header {
		c.Print(key, " : ")
		magenta.Print(value, "\n")
	}

	str := string(body)
	if strings.Contains(heads, "json") {
		retbody = color.HiYellowString("\nStatus:\t\t%s\n\nStatusCode:\t%d\n", resp.Status, resp.StatusCode) + fmt.Sprintf("\n%s\n", string(pretty.Color(pretty.Pretty(body), nil)))
	} else if strings.Contains(heads, "xml") || strings.Contains(heads, "html") || strings.Contains(heads, "plain") {
		var s string
		if strings.Contains(heads, "plain") {
			s = str
		} else {
			s = c.Sprint(gohtml.Format(str))
		}
		retbody = color.HiYellowString("\nStatus:\t\t%s\n\nStatusCode:\t%d\n", resp.Status, resp.StatusCode) + fmt.Sprintf("\n%s\n", s)
	}

	return retbody, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// Contenttypes are used in Place for ctypes for sending requests
var Contenttypes = map[string]string{
	"html":  "text/html",
	"js":    "application/json",
	"xml":   "application/xml",
	"plain": "text/plain",
}
