package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sky0621/tips-go/library_selection/sqlparser/validator"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(_ context.Context) error {
	configPath := "sqlc.yaml"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	results, err := validator.ValidateSQLCProject(configPath)
	if err != nil {
		return err
	}

	var hasError bool
	for _, result := range results {
		if result.Err != nil || !result.Valid {
			hasError = true
		}
		fmt.Println(result.String())
	}

	if hasError {
		return fmt.Errorf("validation failed")
	}

	return nil
}
