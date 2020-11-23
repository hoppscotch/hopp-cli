package methods

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/pkg/browser"
)

//GenerateDocs generates the Documentation site from the hoppscotch-collection.json
func GenerateDocs(filename string) {
	cwd, _ := os.Getwd()
	colls, err := ReadCollection(filename)
	if err != nil {
		log.Printf("Error Occured %v", err)
	}
	t := template.Must(template.ParseFiles(filepath.Join(cwd, "methods/templates/template.md")))

	// Create the file
	f, err := os.Create(filepath.Join(cwd, "methods/templates/README.md"))
	if err != nil {
		log.Printf("File Creation Error: %v", err)
	}

	// Execute the template to the file.
	if err = t.Execute(f, colls); err != nil {
		log.Println(err)
	}

	// Close the file when done.
	f.Close()
	fs := http.FileServer(http.Dir(filepath.Join(cwd, "methods/templates/")))
	http.Handle("/", fs)

	log.Printf("\033[1;36m%s\033[0m", "Server Listening at http://localhost:1341")
	browser.OpenURL("http://localhost:1341/") // AutoOpen the Broswer
	http.ListenAndServe(":1341", nil)
}
