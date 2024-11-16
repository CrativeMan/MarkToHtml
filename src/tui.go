package main

import "github.com/charmbracelet/huh"

func promptForInputFile() string {
	var inputFile string
	huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Placeholder("Enter the path to the file you want to convert").
				Value(&inputFile),
		),
	).WithShowHelp(true).Run()

	return inputFile
}
