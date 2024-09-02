package migrations

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh"
)

var (
	op            int
	migrationName string
)

func Run() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Select operation:").
				Description("make sure you have sql-migrate installed (https://github.com/rubenv/sql-migrate). If you'd like more options interact with the sql-migrate cli directly").
				Options(
					huh.NewOption("Make Migration", 1),
					huh.NewOption("Migrate up", 2),
					huh.NewOption("Migrate down", 3),
					huh.NewOption("Migrate redo", 4),
					huh.NewOption("Migrate status", 5),
				).
				Value(&op),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	switch o := op; o {
	case 1:
		createMigration()
	case 2:
		migrateUp()
	case 3:
		migrateDown()
	case 4:
		migrateRedo()
	case 5:
		migrateStatus()
	default:
		fmt.Println("No valid option selected")
	}
}

func createMigration() {
	// Create a form to ask for the migration name
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Migration Name").
				Placeholder("Enter migration name").
				Value(&migrationName),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if migrationName == "" {
		fmt.Println("Migration name cannot be empty")
		return
	}
	// Command to run "sql-migrate up"
	cmd := exec.Command("sql-migrate", "new", migrationName)

	// Set the command's standard output and error to the program's output and error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	_ = cmd.Run()
}

func migrateUp() {
	// Command to run "sql-migrate up"
	cmd := exec.Command("sql-migrate", "up")

	// Set the command's standard output and error to the program's output and error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	_ = cmd.Run()
}

func migrateDown() {
	// Command to run "sql-migrate up"
	cmd := exec.Command("sql-migrate", "down")

	// Set the command's standard output and error to the program's output and error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	_ = cmd.Run()
}

func migrateRedo() {
	// Command to run "sql-migrate up"
	cmd := exec.Command("sql-migrate", "redo")

	// Set the command's standard output and error to the program's output and error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	_ = cmd.Run()
}
func migrateStatus() {
	// Command to run "sql-migrate up"
	cmd := exec.Command("sql-migrate", "status")

	// Set the command's standard output and error to the program's output and error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	_ = cmd.Run()
}
