package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var short = "todolist is a simple golang based CLI application"
var long = "Using todolist CLI you can easily \ncreate, read, update, delete your todolike you do in any GUI based application"

var RootCmd = &cobra.Command{
	Use:   "todolist",
	Short: short,
	Long:  long,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(short)
			fmt.Printf(long, "\n \n \n")
			fmt.Println("Use 'todolist --help' for usage information")
			return
		}
		return
	},
}
