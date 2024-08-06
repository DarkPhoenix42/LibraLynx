CREATE TABLE transactions(
    transaction_id INT AUTO_INCREMENT PRIMARY KEY,
    book_id INT,
    user_id INT,
    date DATE,
    type ENUM('borrow', 'return'),
    status ENUM('pending','rejected','accepted', 'archived'),
    FOREIGN KEY (book_id) REFERENCES books(book_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);