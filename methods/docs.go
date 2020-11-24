package methods

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/knadh/stuffbin"
	"github.com/pkg/browser"
)

//Fstruct handles the buffer for generated README.md File
type Fstruct struct {
	b bytes.Buffer
}

// Name holds the FileName, here README.md
func (f *Fstruct) Name() string { return "README.md" }

// Size holds the size of the File
func (f *Fstruct) Size() int64 { return int64(len(f.b.Bytes())) }

// Mode holds the file Mode
func (f *Fstruct) Mode() os.FileMode { return 0755 }

// ModTime holds creation time of File
func (f *Fstruct) ModTime() time.Time { return time.Now() }

// IsDir checks if True
func (f *Fstruct) IsDir() bool { return false }

// Sys - I have no idea
func (f *Fstruct) Sys() interface{} { return nil }

//GenerateDocs generates the Documentation site from the hoppscotch-collection.json
func GenerateDocs(filename string) {
	execPath, err := os.Executable() //get Executable Path for StuffBin
	if err != nil {
		log.Fatal(err)
	}
	fs, err := initFileSystem(execPath) //Init Virtual FS
	if err != nil {
		log.Fatal(err)
	}

	colls, err := ReadCollection(filename)
	if err != nil {
		log.Printf("Error Occured %v", err)
	}
	// FuncMap for the HTML template
	fmap := map[string]interface{}{
		"html": func(val string) string { return val },
	}

	t, err := stuffbin.ParseTemplates(fmap, fs, "/template.md")

	var f Fstruct
	// Execute the template to the file.
	if err = t.Execute(&f.b, colls); err != nil {
		log.Println(err)
	}

	if err := fs.Add(stuffbin.NewFile("/README.md", &f, f.b.Bytes())); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		out, err := fs.Read("templates/index.html")
		if err != nil {
			log.Fatal(err)
		}
		w.Write(out)
	})
	http.Handle("/static/", http.StripPrefix("/static/", fs.FileServer()))

	log.Printf("\033[1;36m%s\033[0m", "Server Listening at http://localhost:1341")

	browser.OpenURL("http://localhost:1341") // AutoOpen the Broswer

	http.ListenAndServe(":1341", nil)
}
