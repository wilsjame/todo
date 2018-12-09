//TODO lol
// Change todo.txt when removing
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

var todos map[int]string

func display(m map[int]string) {
	var keys []int

	// Print todo list in key order
	for k := range todos {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Println(key, todos[key])
	}
}

func main() {
	var addPtr = flag.String("a", "null", "add todo") // -a adds a todo item
	var removePtr = flag.Int("r", 0, "remove todo")   // -r removes a todo item

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
	flag.Parse() // Check command line arguments
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
		for i := *removePtr; i < len(todos)+1; i++ {
			todos[i] = todos[i+1]
		}
		// Remove trailing empty todo
		if todos[len(todos)] == "" {
			delete(todos, len(todos))
		}
	} else {
		display(todos)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}
