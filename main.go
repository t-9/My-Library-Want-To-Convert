package main

import (
	"fmt"

	"github.com/t-9/My-Library-Want-To-Convert/chuni"
)

func main() {
	plain := "我がライブラが変換せしめん"
	fmt.Println(new(chuni.Text).Set(plain))
}
