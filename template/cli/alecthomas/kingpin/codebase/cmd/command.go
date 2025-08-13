package cmd

import (
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"{{ .ProjectName }}/pkg/echo"
	"{{ .ProjectName }}/pkg/print"
)

var (
	app = kingpin.New("{{ .ProjectName }}", "A command-line application.")

	printCmd = app.Command("print", "Print anything to the screen")
	printMsg = printCmd.Arg("text", "Text message to print").Strings()

	echoCmd   = app.Command("echo", "Echo anything to the screen")
	echoText  = echoCmd.Arg("text", "Text message to echo").Strings()
	echoTimes = echoCmd.Flag("times", "Echo anything to the screen more times").Int()
)

// Commands represents commands to execute
func Commands(args []string) {
	switch kingpin.MustParse(app.Parse(args)) {
	// print message
	case printCmd.FullCommand():
		print.Message(strings.Join(*printMsg, " "))

		// Echo message
	case echoCmd.FullCommand():
		if *echoTimes != 0 {
			echo.WithTimes(strings.Join(*echoText, " "), *echoTimes)
		} else {
			echo.WithoutTime(strings.Join(*echoText, " "))
		}

	}
}
