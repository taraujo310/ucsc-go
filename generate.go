package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func goBasicSetup() string {
	return `package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
`
}

func getLastDirNumber() string {
	var lastDirName string
	entries, err := os.ReadDir("./")

	if err != nil {
		log.Fatal(err)
	}

	for n := len(entries) - 1; n >= 0; n-- {
		if entries[n].IsDir() {
			lastDirName = entries[n].Name()
			break;
		}
	}

	currentNumber, err := strconv.Atoi(lastDirName[0:3])

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%03d-", currentNumber+1)
}

func logAction(action string, message string) {
	// cRed := "\033[31m"
	cGreen := "\033[32m"
	reset := "\033[0m"
	bold := "\033[1m"

	fmt.Println(bold, cGreen, action, reset, " ", message)
}

func createFile(filepath string, content string) {
	goFile, err := os.Create(filepath)

	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(filepath, []byte(content), 0700); err != nil {
		log.Fatal(err)
	}

	defer goFile.Close()

	if err := os.Chmod(filepath, 0700); err != nil {
		log.Fatal(err)
	}

	logAction("create file", filepath)
}

func insertClassIntoSummary(className string, readmeFilepath string, description string) {
	f, err := os.OpenFile("./README.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	title := cases.Title(language.English, cases.Compact).String(className)

	topic := "\n- [" + title + "](" + readmeFilepath + "): " + description + "\n"

	if _, err := f.WriteString(topic); err != nil {
		log.Println(err)
	}

	defer f.Close()

	logAction("append in to summary", title)
}

func scaffoldClassDirectory(className string, description string) {
	number := getLastDirNumber()
	newDir := number + className

	fmt.Println("dir ", newDir)
	completePath := "./" + newDir

	if err := os.Mkdir(completePath, 0700); err != nil {
		log.Fatal(err)
	}

	goFilepath := completePath + "/" + className + ".go"
	readmeFilepath := completePath + "/README.md"

	createFile(goFilepath, goBasicSetup())
	createFile(readmeFilepath, "# " + cases.Title(language.English, cases.Compact).String(className))
	insertClassIntoSummary(className, readmeFilepath, description)
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("missing class name");
	} else {
		className := flag.Arg(0)
		description := flag.Arg(1)
		scaffoldClassDirectory(className, description)
	}

	os.Exit(0)
}