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

	fmt.Println("file ", filepath)
}

func scaffoldClassDirectory(className string) {
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
	createFile(readmeFilepath, "# " + cases.Title(language.English, cases.Compact).String(className) )
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("missing class name");
	} else {
		className := flag.Arg(0)
		scaffoldClassDirectory(className)
	}

	os.Exit(0)
}