package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli"
)

// AddCommand abc
func addCommand(app *cli.App) *cli.App {
	//connect to database
	db, errdb := sql.Open("sqlite3", "./db/todo_cli.db")
	checkErr(errdb)

	//create cli commands, including showing todo list, add a task and remove a task
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

	done := cli.Command{
		Name:    "done",
		Aliases: []string{"d"},
		Usage:   "complete a task on the list",
		Action: func(c *cli.Context) error {
			todoText := strings.TrimSpace(c.Args().First())

			//update task
			stmt, errrow := db.Prepare("UPDATE todo SET status=? where text=?")
			checkErr(errrow)

			res, errres := stmt.Exec(true, todoText)
			checkErr(errres)

			rowAff, errAff := res.RowsAffected()
			checkErr(errAff)

			if rowAff == 0 {
				fmt.Println("task [", todoText, "] doesn't exist")
				return nil
			}

			db.Close()

			fmt.Println("task [", todoText, "] completed")
			return nil
		},
	}

	add := cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a task to the list",
		Action: func(c *cli.Context) error {
			newTodoText := strings.TrimSpace(c.Args().First())

			stmt, errstmt := db.Prepare("INSERT INTO todo (text) values(?)")
			checkErr(errstmt)

			_, errres := stmt.Exec(newTodoText)
			checkErr(errres)

			fmt.Println("task [", newTodoText, "] added")
			return nil
		},
	}

	app.Commands = append(app.Commands, list)
	app.Commands = append(app.Commands, add)
	app.Commands = append(app.Commands, done)

	return app
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
