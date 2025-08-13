package cmd

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"{{ .ProjectName }}/pkg/consts"
	"{{ .ProjectName }}/pkg/echo"
	"{{ .ProjectName }}/pkg/print"
)

// NewCmd as
func NewCmd() (app *cli.App) {
	app = cli.NewApp()
	SetInfo(app)
	addCommands(app)
	return
}

// SetInfo Basic Info of cli
func SetInfo(app *cli.App) {
	app.Name = consts.ProjectName
	app.Usage = consts.Usage
	app.Author = consts.Author
	app.Version = consts.Version
}

// Run to start accepting commands
func Run(app *cli.App) {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func addCommands(app *cli.App) {
	app.Commands, _ = getCommands()
}

func getCommands() ([]cli.Command, error) {
	return []cli.Command{
		{
			Name:    "print",
			Aliases: []string{"p"},
			Usage:   "Print anything to the screen",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "msg"},
			},
			Action: func(c *cli.Context) error {
				print.Message(c.String("msg"))
				return nil
			},
		},
		{
			Name:    "echo",
			Aliases: []string{"e"},
			Usage:   "Echo anything to the screen",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "msg"},
				&cli.IntFlag{Name: "times"},
			},
			Action: func(c *cli.Context) error {
				if c.Int("times") == 0 {
					echo.WithoutTime(c.String("msg"))
				} else {
					echo.WithTimes(c.String("msg"), c.Int("times"))
				}
				return nil
			},
		},
	}, nil
}
