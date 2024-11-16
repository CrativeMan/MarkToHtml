package main

import (
	"fmt"
	"os"
	"strings"
)

var isStandalone bool = false

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-s" {
			isStandalone = true
		}
	}
	run()
}

func run() {
	inputPath := promptForInputFile()

	fileinfo, err := os.Stat(inputPath)
	if err != nil {
		fmt.Println(err)
	}

	if fileinfo.IsDir() {
		// get all files in the directory
		files, names, err := getFilesInDirectory(inputPath)
		if err != nil {
			fmt.Println(err)
		}

		// for each file, read the contents
		for i, file := range files {
			contents, err := readFile(file)
			if err != nil {
				fmt.Println(err)
			}

			// optimize the contents
			optimized, err := optimizeMarkdown(contents)
			if err != nil {
				fmt.Println(err)
			}

			// convert the optimized contents to standalone HTML
			if isStandalone {
				optimized, err = parseToStandalone(optimized)
				if err != nil {
					fmt.Println(err)
				}
			}

			// write the optimized contents to a new file, remove old extension and add .html
			outputPath := "output/" + strings.TrimSuffix(names[i], ".md") + ".html"
			err = writeFile(outputPath, optimized)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("File written to", outputPath)
			}

		}
	} else {
		contents, err := readFile(inputPath)
		if err != nil {
			fmt.Println(err)
		}

		optimized, err := optimizeMarkdown(contents)
		if err != nil {
			fmt.Println(err)
		}

		if isStandalone {
			optimized, err = parseToStandalone(optimized)
			if err != nil {
				fmt.Println(err)
			}
		}

		outputPath := "output/" + strings.TrimSuffix(fileinfo.Name(), ".md") + ".html"
		err = writeFile(outputPath, optimized)
		if err != nil {
			fmt.Println(err)
		}
	}
}
