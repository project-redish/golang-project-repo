package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"{{ .ProjectName }}/pkg/echo"
)

var echoTimes int

// cmdEcho represents the echo command
var cmdEcho = &cobra.Command{
	Use:   "echo [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		echo.WithoutTime(strings.Join(args, " "))
	},
}

// cmdTimes represents the times sub command of echo
var cmdTimes = &cobra.Command{
	Use:   "times [string to echo]",
	Short: "Echo anything to the screen more times",
	Long: `echo things multiple times back to the user by providing
a count and a string.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		echo.WithTimes(strings.Join(args, " "), echoTimes)
	},
}

func init() {
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	rootCmd.AddCommand(cmdEcho)
	cmdEcho.AddCommand(cmdTimes)
}
