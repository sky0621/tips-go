# Architecture Model

## Testing

### Database Tests

This project includes tests that interact with a test database. The test database is configured in the `docker-compose.yml` file as a separate service from the main database.

#### Running the Test Database

To start the test database, run:

```bash
docker-compose up -d mysql_test
```

#### Environment Variables

The test database connection can be configured using the following environment variables:

- `TEST_DB_USER`: The database user (default: "test_user")
- `TEST_DB_PASSWORD`: The database password (default: "test_password")
- `TEST_DB_NAME`: The database name (default: "test_db")
- `TEST_DB_HOST`: The database host (default: "localhost")
- `TEST_DB_PORT`: The database port (default: "3307")

#### Running Tests

To run the tests, first make sure the test database is running, then run:

```bash
# Run all tests
go test ./...

# Skip tests that require a database connection
go test -short ./...

# Run tests for a specific package
go test ./internal/content/infrastructure/...
```

#### Test Data

The tests create and clean up their own test data. Each test function that requires database access should:

1. Call `testDB(t)` to get a database connection
2. Set up test data using `setupTestData(t, db)`
3. Clean up test data using `cleanupTestData(t, db, contentID)`

Example:

```go
func TestSomeFunction(t *testing.T) {
    // Skip if not running integration tests
    if testing.Short() {
        t.Skip("Skipping integration test")
    }

    // Get test database connection
    db := testDB(t)
    defer db.Close()

    // Setup test data
    contentID, err := setupTestData(t, db)
    if err != nil {
        t.Fatalf("Failed to setup test data: %v", err)
    }
    defer cleanupTestData(t, db, contentID)

    // Run your test...
}
```