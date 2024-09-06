package projectpackage

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
)

var (
	oldModuleName string
	newModuleName string
)

func Run() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the old module name").
				Description("It's really just replacing every occurence of the string with another string and running go mod tidy after.").
				Placeholder("e.g., github.com/yourusername/oldmodulename").
				Value(&oldModuleName),
			huh.NewInput().
				Title("Enter the new module name").
				Placeholder("e.g., github.com/yourusername/newmodulename").
				Value(&newModuleName),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if oldModuleName == "" || newModuleName == "" {
		fmt.Println("Both old and new module names must be provided.")
		return
	}

	err = updateModuleName(newModuleName)
	if err != nil {
		fmt.Printf("Error updating module name in go.mod: %s\n", err)
		return
	}

	err = updateImportPaths(oldModuleName, newModuleName)
	if err != nil {
		fmt.Printf("Error updating import paths: %s\n", err)
		return
	}

	fmt.Println("Running `go mod tidy`...")
	runGoModTidy()

	fmt.Println("Module name update completed successfully.")
}

func updateModuleName(newModuleName string) error {
	// Read go.mod
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return fmt.Errorf("failed to read go.mod: %w", err)
	}

	// Replace module name
	lines := strings.Split(string(data), "\n")
	if len(lines) > 0 && strings.HasPrefix(lines[0], "module ") {
		lines[0] = "module " + newModuleName
	} else {
		return fmt.Errorf("go.mod does not start with a module declaration")
	}

	// Write back to go.mod
	return os.WriteFile("go.mod", []byte(strings.Join(lines, "\n")), 0644)
}

func updateImportPaths(oldModuleName, newModuleName string) error {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Replace all instances of the old module name with the new one
		newData := strings.ReplaceAll(string(data), oldModuleName, newModuleName)
		if newData != string(data) {
			err = os.WriteFile(path, []byte(newData), 0644)
			if err != nil {
				return err
			}
			fmt.Printf("Updated import paths in: %s\n", path)
		}

		return nil
	})
	return err
}

func runGoModTidy() {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running `go mod tidy`: %s\n", err)
	}
}
