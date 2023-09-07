CREATE TABLE
    users (
        id INT PRIMARY KEY AUTO_INCREMENT,
        username VARCHAR(20) NOT NULL,
        password VARCHAR(20) NOT NULL,
        created_at TIMESTAMP CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP DEFAULT NULL,
        updated_at TIMESTAMP DEFAULT NULL
    );

CREATE TABLE
    bookmarks (
        id INT PRIMARY KEY AUTO_INCREMENT,
        ayat VARCHAR(20) NOT NULL,
        arab VARCHAR(20) NOT NULL,
        created_at TIMESTAMP CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP DEFAULT NULL,
        updated_at TIMESTAMP DEFAULT NULL
    );
