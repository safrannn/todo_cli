package command

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

func commandDone(db *sql.DB) cli.Command {
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

	return done
}
