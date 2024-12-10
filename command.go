package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *cmdFlags) Execute(todos *Todos) {
	if cf.Add != "" {
		todos.add(cf.Add)
	}

	if cf.Del != -1 {
		todos.delete(cf.Del)
	}

	if cf.Edit != "" {
		parts := strings.Split(cf.Edit, ":")
		if len(parts) != 2 {
			fmt.Println("Error, invalid edit format. Please user id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error, invalid id")
			os.Exit(1)
		}
		todos.edit(index, parts[1])
	}

	if cf.Toggle != -1 {
		todos.toggle(cf.Toggle)
	}

	if cf.List {
		todos.display()
	}
}
