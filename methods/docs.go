package methods

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/knadh/stuffbin"
	"github.com/pkg/browser"
	"github.com/urfave/cli"
)

//FileTrunk handles the buffer for generated README.md File
type FileTrunk struct{ bytes.Buffer }

// Name holds the FileName, here README.md
func (f *FileTrunk) Name() string { return "README.md" }

// Size holds the size of the File
func (f *FileTrunk) Size() int64 { return int64(f.Len()) }

// Mode holds the file Mode
func (f *FileTrunk) Mode() os.FileMode { return 0755 }

// ModTime holds creation time of File
func (f *FileTrunk) ModTime() time.Time { return time.Now() }

// IsDir checks if True
func (f *FileTrunk) IsDir() bool { return false }

// Sys - I have no idea
func (f *FileTrunk) Sys() interface{} { return nil }

//GenerateDocs generates the Documentation site from the hoppscotch-collection.json
func GenerateDocs(c *cli.Context) error {
	execPath, err := os.Executable() //get Executable Path for StuffBin
	if err != nil {
		return err
	}
	fs, err := initFileSystem(execPath) //Init Virtual FS
	if err != nil {
		return err
	}

	colls, err := ReadCollection(c.Args().Get(0))
	if err != nil {
		return err
	}
	// FuncMap for the HTML template
	fmap := map[string]interface{}{
		"html": func(val string) string { return val },
	}

	t, err := stuffbin.ParseTemplates(fmap, fs, "/template.md")

	// f will be used to store rendered templates in memory.
	var f FileTrunk

	// Execute the template to the file.
	if err = t.Execute(&f, colls); err != nil {
		return err
	}

	if err := fs.Add(stuffbin.NewFile("/README.md", &f, f.Bytes())); err != nil {
		return err
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		out, err := fs.Read("templates/index.html")
		if err != nil {
			log.Println(err)
		}
		w.Write(out)
	})
	PortStr := ":" + strconv.Itoa(c.Int("port"))
	URL := fmt.Sprintf("http://localhost%s", PortStr)

	http.Handle("/static/", http.StripPrefix("/static/", fs.FileServer()))

	log.Printf("\033[1;36mServer Listening at %s\033[0m", URL)

	if !c.Bool("browser") { //Check if User wants to open the Broswer
		browser.OpenURL(URL) // AutoOpen the Broswer
	}

	http.ListenAndServe(PortStr, nil)
	return nil
}
