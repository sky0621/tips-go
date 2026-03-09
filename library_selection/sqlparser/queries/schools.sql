-- name: GetSchoolByID :one
SELECT id, name
FROM schools
WHERE schools.id = ?;

-- name: ListSchoolStudents :many
SELECT s.id, s.name
FROM students AS s
JOIN schools AS sc ON sc.id = s.school_id
WHERE sc.id = ?;

-- name: UpdateSchoolName :exec
UPDATE schools
SET name = ?
WHERE id = ?;

-- name: DeleteSchool :exec
DELETE FROM schools
WHERE schools.id = ?;

-- name: UnsafeListSchools :many
SELECT id, name
FROM schools;
