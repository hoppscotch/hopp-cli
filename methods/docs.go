package methods

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

//GenerateDocs generates the Documentation site from the hoppscotch-collection.json
func GenerateDocs(filename string) {
	cwd, _ := os.Getwd()
	colls, err := ReadCollection(filename)
	if err != nil {
		log.Printf("Error Occured %v", err)
	}
	t := template.Must(template.ParseFiles(filepath.Join(cwd, "methods/templates/index.html")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := colls
		t.Execute(w, data)
	})
	log.Printf("\033[1;36m%s\033[0m", "Server Listening at http://localhost:1341")
	http.ListenAndServe(":1341", nil)
}
