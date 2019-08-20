package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

type Todo struct {
	Text   string `json:"text"`
	Status string `json:"status"`
}

type Todos struct {
	Todos []Todo `json:"todos"`
}

func createFromFile() Todos {
	jsonFile, errJSONFile := os.Open("todo.json")
	if errJSONFile != nil {
		fmt.Println(errJSONFile)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var todos Todos

	errUnmarshal := json.Unmarshal(byteValue, &todos)
	if errUnmarshal != nil {
		fmt.Print(errUnmarshal)
	}

	return todos
}

func writeToJSON(todos Todos) {
	var jsonData []byte
	jsonData, errJSONData := json.Marshal(todos)
	if errJSONData != nil {
		fmt.Println(errJSONData)
	}

	errWriteFile := ioutil.WriteFile("todo.json", jsonData, 0644)
	if errWriteFile != nil {
		fmt.Println(errWriteFile)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "todo_cli"
	app.Usage = "list tasks"
	app.Action = func(c *cli.Context) error {
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list all tasks",
			Action: func(c *cli.Context) error {
				todos := createFromFile()

				fmt.Println("todo lists")
				for index, todo := range todos.Todos {
					if todo.Status == "NEW" {
						fmt.Println(index, "   NEW    ", todo.Text)
					}
				}

				for index, todo := range todos.Todos {
					if todo.Status == "DONE" {
						fmt.Println(index, "   DONE   ", todo.Text)
					}
				}

				// defer jsonFile.Close()
				return nil
			},
		},
		{
			Name:    "done",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				todos := createFromFile()
				text := strings.TrimSpace(c.Args().First())
				textIndex := -1
				if _, errTextIndex := strconv.Atoi(text); errTextIndex == nil {
					textIndex, _ = strconv.Atoi(text)
				}
				for index, todo := range todos.Todos {
					fmt.Println("index : ", index, "txtindex: ", textIndex)
					if index == textIndex || text == todo.Text {
						todos.Todos[index].Status = "DONE"
					}
				}
				writeToJSON(todos)
				fmt.Println("task [", text, "] completed")
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				newTodo := Todo{Text: strings.TrimSpace(c.Args().First()), Status: "NEW"}
				//read everything from json file then add newTodo to it
				todos := createFromFile()
				todos.Todos = append(todos.Todos, newTodo)
				//write changed todo list to json file
				writeToJSON(todos)
				fmt.Println("task [", newTodo.Text, "] added")
				return nil
			},
		},
	}
	fmt.Println(os.Args)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
