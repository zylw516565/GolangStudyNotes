package main

import (
	"io"
	"os"
	// "fmt"
)

func writeHeader2(w io.Writer, contentType string) error {
	if _, err := w.Write([]byte("Content-Type: ")); err != nil {
		return err
	}

	if _, err := w.Write([]byte(contentType)); err != nil {
		return err
	}
	// ...
	return nil
}

//*********************************************************
func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(s string) (n int, err error)
	}

	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s)
	}

	return w.Write([]byte(s))
}

func writeHeader(w io.Writer, contentType string) error {
	if _, err := writeString(w, "Content-Type: " + contentType); err != nil {
		return err
	}

	return nil
}

func main() {
	writeHeader(os.Stdout, "text/html\n")
}