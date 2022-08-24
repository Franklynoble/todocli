package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Franklynoble/todocli"
)

//Hardcoding  the file name

const todoFileName = ".todo.json"

func main() {
	//Define an  item List

	//Parsing command line flags

	// Keep in mind that the assigned variables are pointers so, to be used later,
	// they have to be dereferenced using the operator * .

	task := flag.String("task", "", "Task to be  included in the Todo list")
	list := flag.Bool("list", false, "List all Tasks")

	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	//Define an  items list
	l := &todocli.List{}

	//Use the Get Method to read to do items from file

	if err := l.Get(todoFileName); err != nil {

		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//Decide what to do based on the  number of arguments  provided

	switch {
	//for no extra arguments, print the list
	case len(os.Args) == 1:
		//List current to do items

		for _, item := range *l {
			fmt.Println(item.Task)
		}
		// Concatenate all provided arguments with a space and
		// add to the list as an item
	default:
		//Contatenate all arguments with a space
		item := strings.Join(os.Args[1:], "")
		//Add the task
		l.Add(item)

		//save the new List
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

}
