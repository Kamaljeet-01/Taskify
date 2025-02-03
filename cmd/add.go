package cmd

import (
	"fmt"

	"github.com/harsh082ip/todolist-CLI-golang/internal/db"
	"github.com/spf13/cobra"
)

var AddTodo = &cobra.Command{
	Use:   "add",
	Short: "Add a Todo",
	Run:   AddTodos,
}

func AddTodos(cmd *cobra.Command, args []string) {
	title, _ := cmd.Flags().GetString("title")
	desc, _ := cmd.Flags().GetString("desc")

	res, err := db.CreateTodo(title, desc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}

func init() {
	AddTodo.Flags().String("title", "", "Title of the Todo")
	AddTodo.Flags().String("desc", "", "Optional: Description of Todo")
	_ = AddTodo.MarkFlagRequired("title")
	// _ = AddTodo.MarkFlagRequired("desc")
}
