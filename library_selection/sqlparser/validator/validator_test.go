package validator

import (
	"path/filepath"
	"testing"
)

func TestValidateSQLCProject(t *testing.T) {
	t.Parallel()

	results, err := ValidateSQLCProject(filepath.Join("..", "sqlc.yaml"))
	if err != nil {
		t.Fatalf("ValidateSQLCProject returned error: %v", err)
	}

	if len(results) != 5 {
		t.Fatalf("expected 5 results, got %d", len(results))
	}

	var ngCount int
	for _, result := range results {
		if !result.Valid {
			ngCount++
		}
	}

	if ngCount != 1 {
		t.Fatalf("expected 1 invalid statement, got %d", ngCount)
	}
}

func TestValidateSQLFile(t *testing.T) {
	t.Parallel()

	results, err := ValidateSQLFile(filepath.Join("..", "queries", "schools.sql"), newOptions())
	if err != nil {
		t.Fatalf("ValidateSQLFile returned error: %v", err)
	}

	cases := map[string]bool{
		"GetSchoolByID":      true,
		"ListSchoolStudents": true,
		"UpdateSchoolName":   true,
		"DeleteSchool":       true,
		"UnsafeListSchools":  false,
	}

	for _, result := range results {
		want, ok := cases[result.Name]
		if !ok {
			t.Fatalf("unexpected statement %q", result.Name)
		}
		if result.Valid != want {
			t.Fatalf("%s: got valid=%v want %v", result.Name, result.Valid, want)
		}
	}
}

func TestValidateSQLCProject_WithExcludedNames(t *testing.T) {
	t.Parallel()

	results, err := ValidateSQLCProject(
		filepath.Join("..", "sqlc.yaml"),
		WithExcludedNames([]string{"UnsafeListSchools"}),
	)
	if err != nil {
		t.Fatalf("ValidateSQLCProject returned error: %v", err)
	}

	for _, result := range results {
		if result.Name != "UnsafeListSchools" {
			continue
		}
		if !result.Skipped {
			t.Fatalf("expected UnsafeListSchools to be skipped")
		}
		if !result.Valid {
			t.Fatalf("expected skipped result to be treated as valid")
		}
		return
	}

	t.Fatalf("UnsafeListSchools result not found")
}
