package main

import (
	"fmt"
	"os"

	"github.com/Franklynoble/todocli"
)

//Hardcoding  the file name

const todoFileName = ".todo.json"

func main() {
	//Define an  item List

	l := &todocli.List{}

	//Use the Get Method to read to do items from file

	if err := l.Get(todoFileName); err != nil {

		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
