package main

import (
	"fmt"

	"github.com/t-9/My-Library-Want-To-Convert/chuni"
	chuniio "github.com/t-9/My-Library-Want-To-Convert/io"
)

const (
	plainTextFileName = "bin/plain.txt"
)

func main() {

	plain := new(chuniio.Text)
	if err := plain.Load(plainTextFileName); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(new(chuni.Text).Set(plain.String()))
}
