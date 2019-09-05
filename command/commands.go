package command

import (
	"database/sql"

	"github.com/urfave/cli"
)

// CreateCommand add list, add and done command to app
func CreateCommand(app *cli.App) *cli.App {
	//connect to database
	db, errdb := sql.Open("sqlite3", "./db/todo_cli.db")
	checkErr(errdb)

	//create cli commands, including showing todo list, add a task and remove a task
	app.Commands = append(app.Commands, commandList(db))
	app.Commands = append(app.Commands, commandAdd(db))
	app.Commands = append(app.Commands, commandDone(db))
	return app
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
