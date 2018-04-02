package io

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

// CSV define a csv struct.
type CSV struct {
	values [][]string
}

func (c *CSV) String() string {
	lines := make([]string, len(c.values))
	for i, l := range c.values {
		lines[i] = strings.Join(l, ",")
	}
	return strings.Join(lines, "\n")
}

// Load take a file name and return a error.
func (c *CSV) Load(name string) (err error) {
	f, err := os.Open(name)
	if err != nil {
		return
	}
	defer func() {
		if closeErr := close(f, err); closeErr != nil {
			err = closeErr
		}
	}()

	reader := csv.NewReader(f)
	reader.LazyQuotes = true
	for {
		var rec []string
		rec, err = reader.Read()
		if err == io.EOF {
			err = nil
			return
		}
		if err != nil {
			return
		}
		c.values = append(c.values, rec)
	}
}

// Save take a file name and return a error.
func (c *CSV) Save(name string) error {
	return save(c, name)
}

// Values return values.
func (c *CSV) Values() [][]string {
	return c.values
}
