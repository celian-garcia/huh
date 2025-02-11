package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func main() {
	var file string

	huh.NewForm(
		huh.NewGroup(
			huh.NewFilePicker().
				Title("Select a file:").
				Description("This will be your profile image.").
				AllowedTypes([]string{".png", ".jpeg", ".webp", ".gif"}).
				Value(&file),
		),
	).WithShowHelp(true).Run()
	fmt.Println(file)
}
