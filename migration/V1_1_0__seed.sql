INSERT INTO publisher (publisher_id, name)
VALUES (1, 'Bloomsbury'),
       (2, 'Alfred A. Knopf'),
       (3, 'HarperOne'),
       (4, 'Lemniscaat'),
       (5, 'Arthur A. Levine Books');
SELECT SETVAL('publisher_publisher_id_seq', (SELECT MAX(publisher_id) FROM publisher));


INSERT INTO genre (genre_id, genre)
VALUES (1, 'fiction'),
       (2, 'fantasy'),
       (3, 'children'),
       (4, 'young adult'),
       (5, 'non-fiction'),
       (6, 'business'),
       (7, 'crime'),
       (8, 'self-help'),
       (9, 'inspirational'),
       (10, 'philosophy'),
       (11, 'adult'),
       (12, 'historical');
SELECT SETVAL('genre_genre_id_seq', (SELECT MAX(genre_id) FROM genre));


INSERT INTO author (author_id, name)
VALUES (1, 'J.K. Rowling'),
       (2, 'John Carreyrou'),
       (3, 'Paulo Coelho'),
       (4, 'Mark Manson'),
       (5, 'Thea Beckman'),
       (6, 'Jack Thorne'),
       (7, 'John Tiffany');
SELECT SETVAL('author_author_id_seq', (SELECT MAX(author_id) FROM author));

INSERT INTO book (book_id, title, language, isbn, pages, reading_age, publisher_id, published_at)
VALUES (1, 'Harry Potter and the Sorcerer''s Stone', 'English', '0747532699', 223, 7, 1, '1997-01-01'),
       (2, 'Harry Potter and the Chamber of Secrets', 'English', '0747538492', 251, NULL, 1, '1998-07-02'),
       (3, 'Harry Potter and the Prisoner of Azkaban', 'English', '0747542155', 317, 8, 1, '1999-07-08'),
       (4, 'Harry Potter and the Goblet of Fire', 'English', '074754624X', 636, 8, 1, '2000-07-08'),
       (5, 'Harry Potter and the Order of the Phoenix', 'English', '0747551006', 766, 9, 1, '2003-06-27'),
       (6, 'Harry Potter and the Half-Blood Prince', 'English', '0747581088', 607, 9, 1, '2005-07-16'),
       (7, 'Harry Potter and the Deathly Hallows', 'English', '0545010225', 607, NULL, 1, '2007-07-14'),
       (8, 'Harry Potter and the Cursed Child', 'English', '133821666X', 336, 10, 5,
        '2017-07-25'),
       (9, 'Bad Blood: Secrets and Lies in a Silicon Valley Startup', 'English', '9781524731656', 352, NULL, 2,
        '2018-05-21'),
       (10, 'The Alchemist', 'English', '9780062315007', 175, NULL, 3,
        '1988-01-21'),
       (11, 'The Subtle Art of Not Giving a Fuck', 'English', '9780062457714', 224, NULL, 3,
        '2016-09-13'),
       (12, 'Kruistocht in spijkerbroek', 'Dutch', '9789060691670', 363, NULL, 4,
        '1974-01-01');
SELECT SETVAL('book_book_id_seq', (SELECT MAX(book_id) FROM book));


INSERT INTO book_author (book_id, author_id)
VALUES (1, 1),
       (2, 1),
       (3, 1),
       (4, 1),
       (5, 1),
       (6, 1),
       (7, 1),
       (8, 1),
       (8, 6),
       (8, 7),
       (9, 2),
       (10, 3),
       (11, 4),
       (12, 5);

INSERT INTO book_genre (book_id, genre_id)
VALUES (1, 1),
       (1, 2),
       (1, 3),
       (2, 1),
       (2, 2),
       (2, 3),
       (3, 1),
       (3, 2),
       (3, 3),
       (3, 4),
       (4, 1),
       (4, 2),
       (4, 3),
       (4, 4),
       (5, 1),
       (5, 2),
       (5, 3),
       (5, 4),
       (6, 1),
       (6, 2),
       (6, 3),
       (6, 4),
       (7, 1),
       (7, 2),
       (7, 4),
       (8, 1),
       (8, 2),
       (8, 3),
       (9, 5),
       (9, 6),
       (9, 7),
       (10, 1),
       (10, 8),
       (10, 9),
       (11, 1),
       (11, 8),
       (11, 9),
       (11, 6),
       (11, 10),
       (11, 11),
       (12, 1),
       (12, 2),
       (12, 3),
       (12, 12);
