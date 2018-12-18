package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	directoryPath := flag.String("path", "", "Get directory/files total line count")
	flag.Parse()

	var files []string

	root := *directoryPath
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		log.Fatal("Error occured while file path walk: ", err)
	}

	count := 0
	for _, f := range files {
		cmd := "wc -l " + f
		out, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			fmt.Println("Error occured while executing command: ", err)
		}
		value := strings.Replace(string(out), " ", "", -1)

		// regex to find the line count from the response of wc command
		re := regexp.MustCompile("[0-9]+")
		values := re.FindAllString(value, -1)

		finalCount, _ := strconv.Atoi(values[0])
		count = count + finalCount
	}
	fmt.Println("Total Line Count: ", count)
	fmt.Println("Testing goreport branch issue 1")
}
