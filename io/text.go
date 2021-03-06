package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Text define a text struct.
type Text struct {
	text string
}

func (t *Text) String() string {
	return t.text
}

// Load take a file name and return a error.
func (t *Text) Load(name string) (err error) {
	f, err := os.Open(name)
	if err != nil {
		return
	}
	defer func() {
		if closeErr := close(f, err); closeErr != nil {
			err = closeErr
		}
	}()

	reader := bufio.NewReader(f)
	for {
		var rec []byte
		rec, _, err = reader.ReadLine()
		if err == io.EOF {
			err = nil
			t.text = strings.TrimRight(t.text, "\n")
			return
		}
		if err != nil {
			return
		}
		t.text = fmt.Sprintf("%s%s\n", t.text, string(rec))
	}
}

// Save take a file name and return a error.
func (t *Text) Save(name string) error {
	return save(t, name)
}

// Set take a content.
func (t *Text) Set(content string) {
	t.text = content
}
