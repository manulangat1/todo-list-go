package main

import (
	"fmt"
	"net/http"
)

const myAppName string = "Langat's App"

var shortGolang string = "Watch Go crashCourse"
var fullGolang = "Watch Nana's Golang full course"
var rewardDessert = "Reward myself with a donut"
var maxItemsInGroup uint = 10
var TaskItems = []string{}
var newTask string

func main() {

	TaskItems = append(TaskItems, shortGolang, fullGolang, rewardDessert)
	printTasks()

	// for len(TaskItems) <= int(maxItemsInGroup) {
	// 	fmt.Println("Enter a new task name")
	// 	fmt.Scan(&newTask)

	// 	isValidInput := validateUserInput(newTask)
	// 	fmt.Println(isValidInput)
	// 	if !isValidInput {
	// 		fmt.Println("Invalid input entered")
	// 		break
	// 	}
	// 	addTask(newTask)

	// 	printTasks()
	// }

	http.HandleFunc("/welcome", helloUser)
	http.ListenAndServe(":8000", nil)
	// fmt.Println(TaskItems)

}
func validateUserInput(newTask string) bool {

	return len(newTask) > 3

}
func printTasks() {
	fmt.Printf("Welcome to %v \n", myAppName)
	// for loop
	for idx, item := range TaskItems {
		fmt.Println(idx+1, item)
	}
}

func addTask(newTask string) string {
	TaskItems = append(TaskItems, newTask)
	return "Done"
}

func helloUser(writer http.ResponseWriter, response *http.Request) {
	for _, task := range TaskItems {
		fmt.Fprintf(writer, ". %v \n", task)
	}
}
