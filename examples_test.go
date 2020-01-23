package zippity_test

import (
	"fmt"

	"github.com/groundbreaker/zippity"
)

func Example() {
	pdf := zippity.ReadFile("test.pdf", "fine.pdf")

	zf := zippity.New()
	zip := zf.Add(pdf).Done()
	fmt.Printf("zip is %d bytes", len(zip))
	// Output:
	// zip is 705454 bytes
}

func Example_chained() {
	pdf := zippity.ReadFile("test.pdf", "fine.pdf")

	txt := &zippity.File{
		Name: "test.txt",
		Body: []byte("Already have the bytes? Then, create a literal File."),
	}

	zf := zippity.New()
	zip := zf.Add(pdf).Add(txt).Done()
	fmt.Printf("zip is %d bytes", len(zip))
	// Output:
	// zip is 705620 bytes
}

func ExampleZipfile_Save() {
	pdf := zippity.ReadFile("test.pdf", "fine.pdf")

	txt := &zippity.File{
		Name: "test.txt",
		Body: []byte("Already have the bytes? Then, create a literal File."),
	}

	zf := zippity.New()
	zf.Add(pdf).Add(txt).Save("test.zip")
}
