package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func close(c io.Closer, err error) error {
	if closeErr := c.Close(); closeErr != nil {
		if err == nil {
			return closeErr
		}
		return fmt.Errorf("%s\n%s", err.Error(), closeErr.Error())
	}
	return nil
}

func save(name, content string) (err error) {
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
	_, err = writer.WriteString(content)
	if err != nil {
		return
	}
	writer.Flush()
	return
}
