

todo
====

`todo` is a small command line todo list program. It was tested with Go 1.11 on Ubuntu 18.04 LTS running on Windows Subsystem for Linux (WSL). You can download and install Go by following the [installation instructions](https://golang.org/doc/install).

Installation
------------
 
	go get github.com/wilsjame/todo

About
-----

```
NAME
	todo - an interactive todo list

SYNOPSIS
	todo [-a=add] [-r=remove] item

DESCRIPTION
	todo creates todo.txt in the same directory and 
	adds and removes items to and from todo.txt.
	Run todo with no arguments to display todo.txt. 
	New todos are automatically assigned an ID.

OPTIONS
	-a string
		Add a todo.
	-r int
		Remove a todo by specifying its ID.

EXAMPLES
	todo  
		Display todo.txt.

	todo -a practice
		Add a todo.

	todo -a "vote for pedro"
		Add a longer todo.

	todo -r 1
		Remove todo with ID number 1.
```
