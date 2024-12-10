package main

import (
	// Str "scripts/operations"

	"fmt"
	// nginxClient "github.com/nginxinc/nginx-plus-go-client/client"
)

func main() {
	fmt.Println("Inside main in todo app")
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	// todos.add("First todo")
	// todos.add("learn go")
	// todos.add("learn docker")
	// fmt.Printf("%+v\n\n", todos)
	// todos.delete(0)
	// fmt.Printf("%+v\n\n", todos)
	// todos.toggle(0)
	// todos.display()
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)
}
