package methods

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

//GenerateDocs generates the Documentation site from the hoppscotch-collection.json
func GenerateDocs(filename string) {
	cwd, _ := os.Getwd()
	fmt.Println(filepath.Join(cwd, "methods/template/index.html"))
	colls, err := ReadCollection(filename)
	if err != nil {
		log.Printf("Error Occured %v", err)
	}
	// paths := []string{
	// 	"index.html",
	// }
	t := template.Must(template.ParseFiles(filepath.Join(cwd, "methods/templates/index.html")))
	if err != nil {
		log.Println(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := colls
		t.Execute(w, data)
	})
	http.ListenAndServe(":1341", nil)
}
