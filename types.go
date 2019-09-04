package main

//Todo is used to store a text and status of single task
type Todo struct {
	Index  int
	Text   string
	Status bool
}

//Todos is used to store a list of Todo task
type Todos struct {
	Todos []Todo
}
