//TODO lol
// actually change todo.txt when adding or removing
// shift todos after removing
// do something about map iteration order
// create todo.txt if it doesn't exist?
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var addPtr = flag.String("a", "null", "add todo") // -a adds a todo item.
var removePtr = flag.Int("r", 0, "remove todo")   // -r removes a todo item.
var todos map[int]string                          // Maps todos as keys (int) to values (string).

func display(m map[int]string) {
	// Print entire todo list
	for key, value := range todos {
		fmt.Println(key, value)
	}
}

func main() {
	flag.Parse()
	// Read todo.txt into todos map.
	f, err := os.Open("todo.txt")
	if err == nil {
		todos = make(map[int]string)
		scanner := bufio.NewScanner(f)
		// Loop over all lines in the file and store them as todos.
		for i := 1; scanner.Scan(); i++ {
			line := scanner.Text()
			todos[i] = line
		}
	}
	if *addPtr != "null" {
		// Add todo.
		fmt.Println("add: ", *addPtr)
		todos[len(todos)+1] = *addPtr
	} else if *removePtr != 0 {
		// Remove todo.
		fmt.Println("remove: ", *removePtr)
		delete(todos, *removePtr)
	} else {
		display(todos)
	}
}
