package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sky0621/tips-go/library_selection/sqlparser/validator"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(_ context.Context) error {
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	var excludes string
	fs.StringVar(&excludes, "exclude", "", "comma-separated sqlc query names to skip")
	if err := fs.Parse(os.Args[1:]); err != nil {
		return err
	}

	configPath := "sqlc.yaml"
	if fs.NArg() > 0 {
		configPath = fs.Arg(0)
	}

	results, err := validator.ValidateSQLCProject(configPath, validator.WithExcludedNames(parseCSV(excludes)))
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

func parseCSV(value string) []string {
	if value == "" {
		return nil
	}

	parts := strings.Split(value, ",")
	names := make([]string, 0, len(parts))
	for _, part := range parts {
		name := strings.TrimSpace(part)
		if name == "" {
			continue
		}
		names = append(names, name)
	}
	return names
}
