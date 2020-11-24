package methods

import (
	"log"

	"github.com/knadh/stuffbin"
)

func initFileSystem(binPath string) (stuffbin.FileSystem, error) {
	fs, err := stuffbin.UnStuff(binPath)
	// If files are not stuffed with the binary,
	// try to pick up files from local file system.
	if err == stuffbin.ErrNoID {
		// Running in local mode. Load the required static assets into
		// the in-memory stuffbin.FileSystem.
		log.Printf("unable to initialize embedded filesystem: %v", err)
		log.Printf("using local filesystem for static assets")

		files := []string{
			"./templates/index.html",
			"./templates/template.md:/template.md",
		}

		// mutates err object.
		fs, err = stuffbin.NewLocalFS("/", files...)
	}

	// Either unstuff or NewLocalFS throws error,
	// mutated error value will be returned
	return fs, err
}
