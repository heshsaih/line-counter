package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var verbose = false
var recursive = true
var dirPath string

func countLinesInFile(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := 1
	for scanner.Scan() {
		result++
	}
	if verbose {
		fmt.Println("File name: ", fileName)
		fmt.Println("Line count: ", result)
		fmt.Println()
	}
	return result
}

func getFileNamesRecursively(dirPath string) []string {
	var fileNames []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileNames = append(fileNames, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return fileNames
}

func getFileNames(dirPath string) []string {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}
	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, dirPath+"\\"+file.Name())
		}
	}
	return fileNames
}

func parseFlags() {
	vFlag := flag.Bool("v", false, "More verbose output")
	nrFlag := flag.Bool("nr", false, "Don't count files in subdirectories")
	flag.Parse()
	if flag.Arg(0) == "" {
		pwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		dirPath = pwd

	} else {
    dirPath = flag.Arg(0)
  }
	verbose = *vFlag
	recursive = !*nrFlag
}

func main() {
	parseFlags()
	var files []string
	if recursive {
		files = getFileNamesRecursively(dirPath)
	} else {
		files = getFileNames(dirPath)
	}
	result := 0
	for _, fileName := range files {
		result += countLinesInFile(fileName)
	}
	fmt.Println("Found files: ", len(files))
	fmt.Println("Total amount of lines in directory: ", result)
}
