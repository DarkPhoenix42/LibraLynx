CREATE TABLE books (
    book_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) UNIQUE NOT NULL,
    author VARCHAR(255) NOT NULL DEFAULT 'unknown',
    genre VARCHAR(255) NOT NULL DEFAULT 'unknown',
    available_copies INT UNSIGNED DEFAULT 1
);
