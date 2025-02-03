package cmd

import (
	"fmt"

	"github.com/harsh082ip/todolist-CLI-golang/internal/db"
	"github.com/spf13/cobra"
)

var ListTodo = &cobra.Command{
	Use:   "list",
	Short: "list all todos",
	Run:   ListTodos,
}

var ListCompTodos = &cobra.Command{
	Use:   "comp",
	Short: "list all completed todos",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := db.ListTodoByStatus(true)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, todo := range res {
			fmt.Printf("Id: %v, \n Title: %v, \n Desc: %v \n, Time Created: %v \n Status: %v \n \n ", todo.Id, todo.Title, todo.Description, todo.Time, "Completed")
		}
	},
}

var ListIncompTodos = &cobra.Command{
	Use:   "incomp",
	Short: "list all incompleted todos",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := db.ListTodoByStatus(false)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, todo := range res {
			fmt.Printf("Id: %v, \n Title: %v, \n Desc: %v \n, Time Created: %v \n Status: %v \n \n ", todo.Id, todo.Title, todo.Description, todo.Time, "Not Completed")
		}
	},
}

func ListTodos(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetString("id")
	if id != "" {
		res, err := db.ListTodoByID(id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		status := "Not Completed"
		if res.IsCompleted {
			status = "Completed"
		}
		fmt.Printf("Id: %v, \n Title: %v, \n Desc: %v \n, Time Created: %v \n Status: %v \n", res.Id, res.Title, res.Description, res.Time, status)
		return
	}

	todos, err := db.ListAllTodos()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, todo := range todos {
		status := "Not Completed"
		if todo.IsCompleted {
			status = "Completed"
		}
		fmt.Printf("Id: %v, \n Title: %v, \n Desc: %v \n, Time Created: %v \n Status: %v \n \n ", todo.Id, todo.Title, todo.Description, todo.Time, status)
	}
}

func init() {
	ListTodo.Flags().String("id", "", "optional: fetch specific todos by ID")
	// _ = ListTodo.MarkFlagRequired("id")
}
