package command

import (
	"database/sql"
	"fmt"

	"github.com/urfave/cli"
)

func commandList(db *sql.DB) cli.Command {
	list := cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "list all tasks",
		Action: func(c *cli.Context) error {
			var todos Todos

			//get todo list fomr database
			rows, errrow := db.Query("SELECT * FROM todo")
			checkErr(errrow)
			var index int
			var text string
			var status bool

			for rows.Next() {
				errrow = rows.Scan(&index, &text, &status)
				checkErr(errrow)
				newTodo := Todo{index, text, status}
				todos.Todos = append(todos.Todos, newTodo)
			}

			rows.Close()

			//print todo list
			fmt.Println("todo lists : ")
			for _, todo := range todos.Todos {
				if todo.Status == false {
					fmt.Println("   NEW    ", todo.Index, "   ", todo.Text)
				}
			}

			for _, todo := range todos.Todos {
				if todo.Status == true {
					fmt.Println("   DONE   ", todo.Index, "   ", todo.Text)
				}
			}

			return nil
		},
	}
	return list
}
