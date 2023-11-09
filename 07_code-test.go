package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type Folders struct {
	Folder     string
	FolderUser string
	Folder1    string
	Folder2    string
	Folder3    string
	Folder4    string
}

type EmailHeaders struct {
	MessageID               string
	Date                    string
	From                    string
	To                      string
	Subject                 string
	Cc                      string
	MimeVersion             string
	ContentType             string
	ContentTransferEncoding string
	Bcc                     string
	XFrom                   string
	XTo                     string
	XCc                     string
	XBcc                    string
	XFolder                 string
	XOrigin                 string
	XFileName               string
	Message                 string
	Folders
}

func validateFolder(folder []string, i int) string {
	lenFolder := len(folder)
	if lenFolder > i {
		return folder[i]
	} else {
		return ""
	}
}

func main() {
	dir := "./folder_test/pollo/imbox/3."

	content, err := os.ReadFile(dir)
	checkErr(err)
	convertContent := strings.Split(string(content), "\n")
	// Strings convertido en Lineas
	// fmt.Println(convertContent)

	folder := Folders{
		Folder:     "",
		FolderUser: "",
		Folder1:    "",
		Folder2:    "",
		Folder3:    "",
		Folder4:    "",
	}

	headerStrings := EmailHeaders{
		MessageID:               "Message-ID: ",
		Date:                    "Date: ",
		From:                    "From: ",
		To:                      "To: ",
		Subject:                 "Subject: ",
		Cc:                      "Cc: ",
		MimeVersion:             "Mime-Version: ",
		ContentType:             "Content-Type: ",
		ContentTransferEncoding: "Content-Transfer-Encoding: ",
		Bcc:                     "Bcc: ",
		XFrom:                   "X-From: ",
		XTo:                     "X-To: ",
		XCc:                     "X-cc: ",
		XBcc:                    "X-bcc: ",
		XFolder:                 "X-Folder: ",
		XOrigin:                 "X-Origin: ",
		XFileName:               "X-FileName: ",
		Message:                 "",
		Folders:                 folder,
	}

	refValues := reflect.ValueOf(&headerStrings).Elem()
	count := 0
	result := ""

	for id, element := range convertContent {
		if count <= 16 {
			currentField := refValues.FieldByIndex([]int{count})
			nextField := refValues.FieldByIndex([]int{count + 1})
			postField := refValues.FieldByIndex([]int{count + 2})
			if strings.Contains(convertContent[id+1], nextField.String()) {
				result += element
				result = strings.Split(result, currentField.String())[1]
				count += 1
				currentField.SetString(result)
				result = ""
			} else if strings.Contains(convertContent[id+1], postField.String()) {
				result += element
				result = strings.Split(result, currentField.String())[1]
				currentField.SetString(result)
				nextField.SetString("")
				count += 2
				result = ""
			} else {
				result += element
			}
		} else {
			result += element
		}

	}

	Folders := strings.Split(headerStrings.XFolder, "\\")
	folder.Folder = validateFolder(Folders, 1)
	folder.FolderUser = validateFolder(Folders, 2)
	folder.Folder1 = validateFolder(Folders, 3)
	folder.Folder2 = validateFolder(Folders, 4)
	folder.Folder3 = validateFolder(Folders, 5)
	folder.Folder4 = validateFolder(Folders, 6)
	headerStrings.Folders = folder

	headerStrings.Message = result
	jsonData, _ := json.Marshal(headerStrings)
	fmt.Println(string(jsonData))
}
