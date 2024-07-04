SET FOREIGN_KEY_CHECKS = 0;
DELETE FROM books WHERE title IN (
    'And Then There Were None',
    '1984',
    'To Kill a Mockingbird',
    'The Great Gatsby',
    'Pride and Prejudice',
    'The Catcher in the Rye',
    "Harry Potter and the Philosopher's Stone",
    'The Hobbit',
    'The Da Vinci Code',
    'The Lord of the Rings',
    'The Hunger Games',
    'The Alchemist',
    'The Girl with the Dragon Tattoo',
    'Gone with the Wind',
    'The Adventures of Huckleberry Finn',
    'The Shining',
    'Brave New World',
    'The Picture of Dorian Gray',
    'The Road',
    'Frankenstein'
);
SET FOREIGN_KEY_CHECKS = 1;