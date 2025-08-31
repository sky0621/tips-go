CREATE TABLE IF NOT EXISTS Issues (
    issue_id SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS Bugs (
    issue_id BIGINT UNSIGNED PRIMARY KEY,
    FOREIGN KEY (issue_id) REFERENCES Issues(issue_id),
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Features (
    issue_id BIGINT UNSIGNED PRIMARY KEY,
    FOREIGN KEY (issue_id) REFERENCES Issues(issue_id),
    request TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Accounts (
    account_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS Comments (
    comment_id SERIAL PRIMARY KEY,
    issue_id BIGINT UNSIGNED NOT NULL,
    author BIGINT UNSIGNED NOT NULL,
    comment_date DATETIME,
    comment TEXT,
    FOREIGN KEY (issue_id) REFERENCES Issues(issue_id),
    FOREIGN KEY (author) REFERENCES Accounts(account_id)
);
