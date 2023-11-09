package main

import "fmt"

type emailContent struct {
	messageId string
	to        string
	content   string
}

func main() {
	fmt.Println("Creating structures")

	email1 := emailContent{
		messageId: "1",
		to:        "To",
		content:   "Content",
	}

	fmt.Println(email1)

}
