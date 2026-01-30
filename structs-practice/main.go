package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type Saver interface {
	Save() error
}

type Outputtable interface { //Embedded interface, an interface that embeds some other interface
	Saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

	printSomething(todo)
}

func printSomething(value any) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Integer:", floatVal)
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println(stringVal)
	}

	// switch value.(type) {		special type of switch for looking and doing something with a specific value type if we recieve a parameter that can be of any type... any or interface{}
	// case int:
	//	fmt.Println("Integer", value)
	// case float:
	//	fmt.Println("Float", value)
	// case string
	//	fmt.Println(string)
	//}
}

func outputData(data Outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving failed.")
		return err
	}

	fmt.Println("Saving succeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

// helper function for getting user input
func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n') // stop reading at new line

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
