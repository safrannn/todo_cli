package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	//create a cli app
	app := cli.NewApp()
	app.Name = "todo_cli"
	app.Usage = "list tasks"
	app.Action = func(c *cli.Context) error {
		return nil
	}

	//create cli commands, including showing todo list, add a task and remove a task
	app.Commands = []cli.Command{}
	app = addCommand(app)
	fmt.Println(os.Args)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
