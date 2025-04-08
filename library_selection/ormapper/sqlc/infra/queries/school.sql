-- name: CreateSchool :execlastid
INSERT INTO school(school_name)
VALUES ('テスト学校');
-- name: CreateGrade :execlastid
INSERT INTO grade(grade_name, school_id)
VALUES ('１年', 1);
-- name: CreateClassBatch :execresult
INSERT INTO class(class_name, grade_id)
VALUES ('１年１組', 1),
       ('１年２組', 1);
-- name: CreateStudentBatch :execresult
INSERT INTO student(name, class_id)
VALUES ('山田太郎', 1),
       ('田中花子', 1),
       ('佐藤次郎', 1),
       ('鈴木三郎', 1),
       ('高橋四郎', 2),
       ('伊藤五郎', 2),
       ('渡辺六郎', 2);

-- name: ListStudentsWithClassWithGradeWithSchool :many
SELECT sc.school_id,
       sc.school_name,
       g.grade_id,
       g.grade_name,
       c.class_id,
       c.class_name,
       s.student_id,
       s.name
FROM student s
         INNER JOIN class c ON s.class_id = c.class_id
         INNER JOIN grade g ON c.grade_id = g.grade_id
         INNER JOIN school sc ON g.school_id = sc.school_id
WHERE sc.school_id = ?;

-- name: GetSchoolByID :one
SELECT * FROM school WHERE school_id = ?;

-- name: ListGradeBySchoolID :many
SELECT * FROM grade WHERE school_id = ?;

-- name: ListClassByGradeID :many
SELECT * FROM class WHERE grade_id = ?;

-- name: ListStudentByClassID :many
SELECT * FROM student WHERE class_id = ?;
