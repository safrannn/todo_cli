package command

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

func commandAdd(db *sql.DB) cli.Command {
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

	return add
}
