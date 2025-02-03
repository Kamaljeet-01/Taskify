package cmd

import (
	"fmt"

	"github.com/harsh082ip/todolist-CLI-golang/internal/db"
	"github.com/spf13/cobra"
)

var DeleteTodo = &cobra.Command{
	Use:   "del",
	Short: "delete a todo by it's id",
	Run:   DeleteTodos,
}

var AllTodoDel = &cobra.Command{
	Use:   "all",
	Short: "delete all todos",
	Run: func(cmd *cobra.Command, args []string) {

		res, err := db.DeleteAllTodos()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
		return
	},
}

var CompTodoDel = &cobra.Command{
	Use:   "comp",
	Short: "delete all completed todos",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := db.DeleteCompOrIncompTODOS(true)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
		return
	},
}

var InCompTodoDel = &cobra.Command{
	Use:   "incomp",
	Short: "delete all incompleted todos",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := db.DeleteCompOrIncompTODOS(false)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
		return
	},
}

func DeleteTodos(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetString("id")

	if id != "" {
		res, err := db.DeleteTodoById(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
		return
	}

	fmt.Println("Exactly one of the flag id, or with command (all, comp, incomp) should be provided")
	return
}

func init() {
	DeleteTodo.Flags().String("id", "", "id of todo you want to delete, Either")
	_ = DeleteTodo.MarkFlagRequired("id")
}
