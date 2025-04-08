CREATE TABLE IF NOT EXISTS users
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS posts
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    title      VARCHAR(255) NOT NULL,
    content    TEXT,
    user_id    BIGINT,
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS comments
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    content    TEXT,
    user_id    BIGINT,
    post_id    BIGINT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS departments
(
    department_id   BIGINT PRIMARY KEY,
    department_name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS employees
(
    employee_id   BIGINT PRIMARY KEY,
    first_name    VARCHAR(50) NOT NULL,
    last_name     VARCHAR(50) NOT NULL,
    salary        DECIMAL(10, 2),
    department_id BIGINT,
    join_date     DATE,
    FOREIGN KEY (department_id) REFERENCES departments (department_id)
);

CREATE TABLE IF NOT EXISTS school
(
    school_id   BIGINT AUTO_INCREMENT PRIMARY KEY,
    school_name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS grade
(
    grade_id   BIGINT AUTO_INCREMENT PRIMARY KEY,
    grade_name VARCHAR(100) NOT NULL,
    school_id  BIGINT,
    FOREIGN KEY (school_id) REFERENCES school (school_id)
);

CREATE TABLE IF NOT EXISTS class
(
    class_id   BIGINT AUTO_INCREMENT PRIMARY KEY,
    class_name VARCHAR(100) NOT NULL,
    grade_id   BIGINT,
    FOREIGN KEY (grade_id) REFERENCES grade (grade_id)
);

CREATE TABLE IF NOT EXISTS student
(
    student_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name       VARCHAR(100) NOT NULL,
    class_id   BIGINT,
    FOREIGN KEY (class_id) REFERENCES class (class_id)
);
