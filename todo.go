//TODO lol
// shift todos after removing
// actually change todo.txt when removing
// do something about map iteration order
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var addPtr = flag.String("a", "null", "add todo") // -a adds a todo item
var removePtr = flag.Int("r", 0, "remove todo")   // -r removes a todo item
var todos map[int]string                          // Maps todos as keys (int) to values (string)

func display(m map[int]string) {
	// Print entire todo list
	for key, value := range todos {
		fmt.Println(key, value)
	}
}

func main() {
	flag.Parse()
	// If todo.txt file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("todo.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	todos = make(map[int]string)
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and store them as todos
	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()
		todos[i] = line
	}
	if *addPtr != "null" {
		// Add todo
		fmt.Println("add: ", *addPtr)
		todos[len(todos)+1] = *addPtr
		if _, err := f.WriteString(*addPtr + "\n"); err != nil {
			log.Fatal(err)
		}
	} else if *removePtr != 0 {
		// Remove todo
		fmt.Println("remove: ", *removePtr)
		delete(todos, *removePtr)
		//TODO shift map key/values
		// update todo.txt
	} else {
		display(todos)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}
