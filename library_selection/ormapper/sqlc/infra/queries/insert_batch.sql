-- name: CreateUsersBatch :execlastid
INSERT INTO users (name, created_at, updated_at) VALUES
   (?, NOW(), NOW()),
   (?, NOW(), NOW()),
   (?, NOW(), NOW()),
   (?, NOW(), NOW()),
   (?, NOW(), NOW());
