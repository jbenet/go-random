package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		usageError()
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usageError()
	}

	if err := writeRandomBytes(num, os.Stdout); err != nil {
		die(err)
	}
}

func usageError() {
	fmt.Fprintf(os.Stderr, "Usage: %s <int>\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Output <int> random bytes (from Go's crypto/rand)\n")
	os.Exit(-1)
}

func die(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v", err)
	os.Exit(-1)
}

func writeRandomBytes(num int, w io.Writer) error {
	r := &io.LimitedReader{R: rand.Reader, N: int64(num)}
	_, err := io.Copy(w, r)
	return err
}
