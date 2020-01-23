// Package zippity creates zip files, quickly.
package zippity

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"os"
)

// Version retuns the SemVer for this library.
func Version() string {
	return "v1.0.1"
}

// File models the data that will be included in a Zipfile.
type File struct {
	Name string
	Body []byte
}

// Zipfile encapsulates the data and behviour required to create zip file
// archives.
type Zipfile struct {
	Body   *bytes.Buffer
	Client *zip.Writer
}

// New creates a new instance of Zipfile.
func New() *Zipfile {
	zf := &Zipfile{
		Body: new(bytes.Buffer),
	}

	zf.Client = zip.NewWriter(zf.Body)
	return zf

}

// ReadFile reads the File.Body from the given path, and returns a new instance
// of File using the given name.
func ReadFile(name string, path string) *File {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return &File{
		Name: name,
		Body: body,
	}
}

// Add a File to the Zipfile. This method can be chained to add multiple files.
func (zf *Zipfile) Add(file *File) *Zipfile {
	f, err := zf.Client.Create(file.Name)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(file.Body)
	if err != nil {
		panic(err)
	}

	return zf
}

// Done returns the Zipfile as bytes. It must be called after you are done
// Adding Files to the Zipfile. You should only call Done() or Save(), but never
// both, or the second call will panic.
func (zf *Zipfile) Done() []byte {
	err := zf.Client.Close()
	if err != nil {
		panic(err)
	}

	return zf.Body.Bytes()
}

// Save the Zipfile to the given path. It must be called after you are done
// Adding Files to the Zipfile. You should only call Save() or Done(), but never
// both, or the second call will panic.
func (zf *Zipfile) Save(path string) {
	ioutil.WriteFile(path, zf.Done(), 0644)
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
