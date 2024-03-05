package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func hashReaderRun() {
	payload := []byte("Happy Chinese New Year")
	if err := hashAndBroadcase(bytes.NewReader(payload)); err != nil {
		fmt.Println("Error!!!")
	}
}

func newHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func hashAndBroadcase(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	hash := sha1.Sum(b)
	fmt.Println("Get Encoding...")
	fmt.Println(hex.EncodeToString(hash[:]))
	return nil
}
