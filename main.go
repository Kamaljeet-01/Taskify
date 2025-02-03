package main

import (
	"fmt"
	"os"

	"github.com/harsh082ip/todolist-CLI-golang/cmd"
	// "github.com/spf13/cobra"
	// "github.com/harsh082ip/todolist-CLI-golang/internal/db"
)

func main() {

	cmd.RootCmd.AddCommand(cmd.ListTodo)
	cmd.RootCmd.AddCommand(cmd.AddTodo)
	cmd.RootCmd.AddCommand(cmd.UpdateTodo)
	cmd.RootCmd.AddCommand(cmd.DeleteTodo)

	cmd.DeleteTodo.AddCommand(cmd.AllTodoDel)
	cmd.DeleteTodo.AddCommand(cmd.CompTodoDel)
	cmd.DeleteTodo.AddCommand(cmd.InCompTodoDel)

	cmd.ListTodo.AddCommand(cmd.ListCompTodos)
	cmd.ListTodo.AddCommand(cmd.ListIncompTodos)
	// cmd.Root2.AddCommand(cmd.RootCmd)

	cmd.UpdateTodo.AddCommand(cmd.UpdateCompletedTodos)
	cmd.UpdateTodo.AddCommand(cmd.UpdateIncompletedTodos)
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
}
