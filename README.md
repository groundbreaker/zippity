# :suspension_railway: zippity

[![GoDoc](https://godoc.org/github.com/groundbreaker/zippity?status.svg)](https://godoc.org/github.com/groundbreaker/zippity)

Creates Zip Files, quickly.


---

## Install

    go get github.com/groundbreaker/zippity

## Usage

```go
import (
  "fmt"

  "github.com/groundbreaker/zippity"
)

func main() {
  // Read a file
	pdf := zippity.ReadFile("test.pdf", "fine.pdf")

  // Or create one from a []byte
	txt := &zippity.File{
		Name: "test.txt",
		Body: []byte("Already have the bytes? Then, create a literal File."),
	}

  // Create a new Zipfile
	zf := zippity.New()

  // Chain as many Add
	zf.Add(pdf).Add(txt)

  zip := zf.Done() // returns the Zipfile as []byte

  // or you can save it to disk with:
  //   zf.Save("path/to/write/file.zip")

  fmt.Printf("zip is %d bytes", len(zip))
}
```
