package main

import (
	"bufio"
	"os"
)

func writeFile(filePath string, contents []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range contents {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	writer.Flush()

	return nil
}

func readFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	return text, nil
}

func getFilesInDirectory(directoryPath string) ([]string, []string, error) {
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return nil, nil, err
	}

	var filePaths []string
	var fileNames []string
	for _, file := range files {
		name := directoryPath + "/" + file.Name()
		filePaths = append(filePaths, name)
		fileNames = append(fileNames, file.Name())
	}

	return filePaths, fileNames, nil
}
