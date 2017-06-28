package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	name := "Tod"
	tpl, err := template.ParseFiles("text_message.phpjsonhtml")
	if err != nil {
		fmt.Println("There was an err parsing the file", err)
	}

	err = tpl.Execute(os.Stdout, name)
	if err != nil {
		fmt.Println("There was an err Executing the file", err)

	}

}
