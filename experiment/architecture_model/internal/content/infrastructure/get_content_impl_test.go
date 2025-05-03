package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/content/application/query"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/shared/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

// testDB is a helper function to create a test database connection
// setupTestData sets up test data in the database
func setupTestData(t *testing.T, db *sql.DB) (string, error) {
	t.Helper()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// Create a new content ID
	contentID := uuid.New().String()

	// Create a new content
	_, err = tx.Exec("INSERT INTO contents (id, name) VALUES (UUID_TO_BIN(?), ?)", contentID, "Test Content")
	if err != nil {
		return "", fmt.Errorf("failed to insert content: %w", err)
	}

	// Create program IDs
	program1ID := uuid.New().String()
	program2ID := uuid.New().String()

	// Create programs
	_, err = tx.Exec("INSERT INTO programs (id, question, answer, content_id) VALUES (UUID_TO_BIN(?), ?, ?, UUID_TO_BIN(?))",
		program1ID, "Test Question 1", "Test Answer 1", contentID)
	if err != nil {
		return "", fmt.Errorf("failed to insert program 1: %w", err)
	}

	_, err = tx.Exec("INSERT INTO programs (id, question, answer, content_id) VALUES (UUID_TO_BIN(?), ?, ?, UUID_TO_BIN(?))",
		program2ID, "Test Question 2", "Test Answer 2", contentID)
	if err != nil {
		return "", fmt.Errorf("failed to insert program 2: %w", err)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return contentID, nil
}

// cleanupTestData cleans up test data from the database
func cleanupTestData(t *testing.T, db *sql.DB, contentID string) {
	t.Helper()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// Delete programs
	_, err = tx.Exec("DELETE FROM programs WHERE content_id = UUID_TO_BIN(?)", contentID)
	if err != nil {
		t.Fatalf("Failed to delete programs: %v", err)
	}

	// Delete content
	_, err = tx.Exec("DELETE FROM contents WHERE id = UUID_TO_BIN(?)", contentID)
	if err != nil {
		t.Fatalf("Failed to delete content: %v", err)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}
}

func TestGetContentImpl_Exec(t *testing.T) {
	// Skip if not running integration tests
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Get test database connection
	db := test.ConnectTestDB(t)
	defer db.Close()

	// Setup test data
	contentID, err := setupTestData(t, db)
	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
	defer cleanupTestData(t, db, contentID)

	// Create the GetContent implementation
	getContent := NewGetContent(db)

	// Test cases
	tests := []struct {
		name    string
		id      string
		want    func(t *testing.T, result query.GetContentReadModel, err error)
		wantErr bool
	}{
		{
			name: "Valid content ID",
			id:   contentID,
			want: func(t *testing.T, result query.GetContentReadModel, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "Test Content", result.Name)
				assert.Len(t, result.Programs, 2)

				// Check that the programs are returned in the correct order
				// Note: The order might not be guaranteed, so we might need to sort them
				questions := []string{result.Programs[0].Question, result.Programs[1].Question}
				answers := []string{result.Programs[0].Answer, result.Programs[1].Answer}

				assert.Contains(t, questions, "Test Question 1")
				assert.Contains(t, questions, "Test Question 2")
				assert.Contains(t, answers, "Test Answer 1")
				assert.Contains(t, answers, "Test Answer 2")
			},
			wantErr: false,
		},
		{
			name: "Invalid content ID",
			id:   uuid.New().String(),
			want: func(t *testing.T, result query.GetContentReadModel, err error) {
				assert.NoError(t, err)
				assert.True(t, result.IsEmpty())
			},
			wantErr: false,
		},
		{
			name: "Empty content ID",
			id:   "",
			want: func(t *testing.T, result query.GetContentReadModel, err error) {
				assert.Error(t, err)
			},
			wantErr: true,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getContent.Exec(context.Background(), tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContentImpl.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				tt.want(t, result, err)
			}
		})
	}
}
