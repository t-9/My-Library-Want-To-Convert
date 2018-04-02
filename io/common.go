package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type textizationFile interface {
	String() string
	Load(string) error
	Save(string) error
}

func close(c io.Closer, err error) (closeErr error) {
	if closeErr = c.Close(); closeErr != nil && err != nil {
		closeErr = fmt.Errorf("%s\n%s", err.Error(), closeErr.Error())
	}
	return
}

func save(t textizationFile, name string) (err error) {
	f, err := os.Create(name)
	if err != nil {
		return
	}
	defer func() {
		if closeErr := close(f, err); closeErr != nil {
			err = closeErr
		}
	}()

	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(t.String())
	if err == nil {
		writer.Flush()
	}
	return
}
