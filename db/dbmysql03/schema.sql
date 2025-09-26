CREATE TABLE IF NOT EXISTS bugs (
    bug_id INT AUTO_INCREMENT PRIMARY KEY,
    description TEXT NOT NULL,
    FULLTEXT KEY idx_description (description)
);
