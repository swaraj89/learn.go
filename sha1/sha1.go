package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	sig, err := sha1Sum("http.log.gz")

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf(sig)
}

// $ cat http.log.gz | gunzip | sha1sum
func sha1Sum(filename string) (string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return "", nil
	}
	//always at function level
	// multiple defers - Execution order LIFO stack
	defer file.Close()

	r, err := gzip.NewReader(file)
	// io.CopyN(os.Stdout, r, 100)
	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}
