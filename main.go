package main

import (
	"log"

	"github.com/t-9/My-Library-Want-To-Convert/chuni"
	"github.com/t-9/My-Library-Want-To-Convert/io"
)

const (
	plainTextFileName = "bin/plain.txt"
	errorFileName     = "bin/error.txt"
	resultFileName    = "bin/result.txt"
)

func main() {
	plain := new(io.Text)
	if err := plain.Load(plainTextFileName); err != nil {
		errorText := new(io.Text)
		errorText.Set(err.Error())
		if saveErr := errorText.Save(errorFileName); saveErr != nil {
			log.Fatalln(saveErr)
		}
		return
	}

	result := new(io.Text)
	result.Set(
		new(chuni.Text).Set(
			plain.String(),
		).String(),
	)
	if saveErr := result.Save(resultFileName); saveErr != nil {
		log.Fatalln(saveErr)
	}
}
