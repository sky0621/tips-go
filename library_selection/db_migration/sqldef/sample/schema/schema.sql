-- users and posts tables synchronized via sqldef
CREATE TABLE users (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(120) NOT NULL,
    display_name VARCHAR(120) NOT NULL,
    login_count INT NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    last_login_at DATETIME NOT NULL,
    `status` ENUM('active', 'archived', 'deleted') NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE posts (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id BIGINT UNSIGNED NOT NULL,
    title VARCHAR(200) NOT NULL,
    body TEXT NOT NULL,
    published_at DATETIME,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT posts_user_id_fk FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE TABLE tags (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(80) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY tags_name_uq (name)
);

CREATE TABLE post_tags (
    post_id BIGINT UNSIGNED NOT NULL,
    tag_id BIGINT UNSIGNED NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (post_id, tag_id),
    CONSTRAINT post_tags_post_id_fk FOREIGN KEY (post_id)
        REFERENCES posts(id)
        ON DELETE CASCADE,
    CONSTRAINT post_tags_tag_id_fk FOREIGN KEY (tag_id)
        REFERENCES tags(id)
        ON DELETE CASCADE
);

CREATE TABLE comments (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    post_id BIGINT UNSIGNED NOT NULL,
    author_name VARCHAR(120) NOT NULL,
    body TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT comments_post_id_fk FOREIGN KEY (post_id)
        REFERENCES posts(id)
        ON DELETE CASCADE
);
