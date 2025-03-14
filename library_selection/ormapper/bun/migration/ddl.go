package migration

const (
	createUsers = `CREATE TABLE IF NOT EXISTS users (
        id BIGINT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL
    )`

	createPosts = `CREATE TABLE IF NOT EXISTS posts (
        id BIGINT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        content TEXT,
        user_id BIGINT,
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL
    )`

	createComments = `CREATE TABLE IF NOT EXISTS comments (
        id BIGINT AUTO_INCREMENT PRIMARY KEY,
        content TEXT,
        user_id BIGINT,
        post_id BIGINT,
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL
    )`
)
