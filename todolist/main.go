package main

import (
	"fmt"
	"net/http"
)

var taskItems = []string{"Watch Go crash course", "Reward myself with a nice meal!"}

func main() {
	fmt.Println("####### Welcome to our Todolist app! #######")

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8000", nil)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user, Welcome to our Todolist App!"

	fmt.Fprintln(writer, greeting)
}

func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")
	for index, task := range taskItems {
		fmt.Printf("index: %v, task: %v\n", index, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}
