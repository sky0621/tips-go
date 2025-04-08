package migration

const (
	createUsers = `CREATE TABLE IF NOT EXISTS users (
        id BIGINT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    )`

	createPosts = `CREATE TABLE IF NOT EXISTS posts (
        id BIGINT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        content TEXT,
        user_id BIGINT,
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    )`

	createComments = `CREATE TABLE IF NOT EXISTS comments (
        id BIGINT PRIMARY KEY,
        content TEXT,
        user_id BIGINT,
        post_id BIGINT,
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    )`

	createDepartments = `CREATE TABLE IF NOT EXISTS departments (
		department_id   BIGINT AUTO_INCREMENT PRIMARY KEY,
		department_name VARCHAR(100) NOT NULL
	)`

	createEmployees = `CREATE TABLE IF NOT EXISTS employees (
		employee_id   BIGINT AUTO_INCREMENT PRIMARY KEY,
		first_name    VARCHAR(50) NOT NULL,
		last_name     VARCHAR(50) NOT NULL,
		salary        DECIMAL(10, 2),
		department_id BIGINT,
		join_date     DATE,
		FOREIGN KEY (department_id) REFERENCES departments (department_id)
	)`

	createSchool = `CREATE TABLE IF NOT EXISTS school (
		school_id   BIGINT AUTO_INCREMENT PRIMARY KEY,
		school_name VARCHAR(100) NOT NULL
	)`

	createGrade = `CREATE TABLE IF NOT EXISTS grade (
		grade_id   BIGINT AUTO_INCREMENT PRIMARY KEY,
		grade_name VARCHAR(100) NOT NULL,
		school_id   BIGINT,
		FOREIGN KEY (school_id) REFERENCES school (school_id)
	)`

	createClass = `CREATE TABLE IF NOT EXISTS class (
		class_id   BIGINT AUTO_INCREMENT PRIMARY KEY,
		class_name VARCHAR(100) NOT NULL,
		grade_id   BIGINT,
		FOREIGN KEY (grade_id) REFERENCES grade (grade_id)
	)`

	createStudent = `CREATE TABLE IF NOT EXISTS student (
		student_id BIGINT AUTO_INCREMENT PRIMARY KEY,
		name       VARCHAR(100) NOT NULL,
		class_id   BIGINT,
		FOREIGN KEY (class_id) REFERENCES class (class_id)
	)`
)
