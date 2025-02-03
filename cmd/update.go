package cmd

import (
	"fmt"

	"github.com/harsh082ip/todolist-CLI-golang/internal/db"
	"github.com/spf13/cobra"
)

var UpdateTodo = &cobra.Command{
	Use:   "update",
	Short: "update a todo using it's Id",
	Run:   UpdateTodos,
}

// var CompTodo = &cobra.

var UpdateCompletedTodos = &cobra.Command{
	Use:   "comp",
	Short: "mark a todo as complete",
	Run:   updateCompTodos,
}

func updateCompTodos(cmd *cobra.Command, args []string) {

	id, _ := cmd.Flags().GetString("id")
	all, _ := cmd.Flags().GetBool("all")
	if all {
		res, err := db.MarkTodosAsCompOrIncomp(id, true, true)
		if err != nil {
			fmt.Println("Error", err.Error())
			return
		}
		fmt.Println(res)
		return
	}

	res, err := db.MarkTodosAsCompOrIncomp(id, true, false)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	fmt.Println(res)
	return

}

var UpdateIncompletedTodos = &cobra.Command{
	Use:   "incomp",
	Short: "mark a todo as incomplete",
	Run:   updateIncompTodos,
}

func updateIncompTodos(cmd *cobra.Command, args []string) {

	id, _ := cmd.Flags().GetString("id")
	all, _ := cmd.Flags().GetBool("all")
	if all {
		res, err := db.MarkTodosAsCompOrIncomp(id, false, true)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}
		fmt.Println(res)
		return
	}
	res, err := db.MarkTodosAsCompOrIncomp(id, false, false)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println(res)
	return
}

func UpdateTodos(cmd *cobra.Command, args []string) {

	id, _ := cmd.Flags().GetString("id")
	title, _ := cmd.Flags().GetString("title")
	desc, _ := cmd.Flags().GetString("desc")
	res, err := db.UpdateTodo(id, title, desc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	return
}

func init() {
	// UpdateTodo.AddCommand(UpdateCompletedTodos)
	UpdateTodo.Flags().String("id", "", "id of the Todo you want to update")
	UpdateTodo.Flags().String("title", "", "updated title of todo")
	UpdateTodo.Flags().String("desc", "", "updated description of todo")
	UpdateCompletedTodos.Flags().String("id", "", "id of the Todo you want to complete")
	UpdateCompletedTodos.Flags().Bool("all", false, "mark all todos as completed")
	UpdateIncompletedTodos.Flags().String("id", "", "id of the Todo you to mark incomplete")
	UpdateIncompletedTodos.Flags().Bool("all", false, "mark all todos as incomplete")
	_ = UpdateTodo.MarkFlagRequired("id")
	_ = UpdateTodo.MarkFlagRequired("title")
	// _ = UpdateCompletedTodos.MarkFlagRequired("id")
	// _ = UpdateIncompletedTodos.MarkFlagRequired("id")
}

// var MarkAllTodosAsComplete = &cobra.Command{
// 	Use:   "compall",
// 	Short: "Mark all todos as complete",
// }

// var MarkAllTodosAsIncomplete = &cobra.Command{
// 	Use:   "incompall",
// 	Short: "Mark all todos as incomplete",
// }
