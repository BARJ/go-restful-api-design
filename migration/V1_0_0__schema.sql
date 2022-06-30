CREATE TABLE author
(
    author_id SERIAL PRIMARY KEY,
    name      TEXT NOT NULL,
    UNIQUE (name)
);

CREATE TABLE publisher
(
    publisher_id SERIAL PRIMARY KEY,
    name         TEXT NOT NULL,
    UNIQUE (name)
);

CREATE TABLE book
(
    book_id      SERIAL PRIMARY KEY,
    title        TEXT    NOT NULL,
    language     TEXT    NOT NULL,
    isbn         TEXT    NOT NULL,
    pages        INTEGER NOT NULL,
    reading_age  INTEGER,
    publisher_id INTEGER NOT NULL,
    published_at DATE    NOT NULL,
    FOREIGN KEY (publisher_id) REFERENCES publisher (publisher_id) ON UPDATE CASCADE ON DELETE CASCADE,
    UNIQUE (title),
    UNIQUE (isbn)
);

CREATE INDEX book_language_index ON book (language);
CREATE INDEX book_reading_age_index ON book (reading_age);
CREATE INDEX book_published_at_index ON book (published_at);

CREATE TABLE book_author
(
    book_id   INTEGER NOT NULL,
    author_id INTEGER NOT NULL,
    FOREIGN KEY (book_id) REFERENCES book (book_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (author_id) REFERENCES author (author_id) ON UPDATE CASCADE ON DELETE CASCADE,
    UNIQUE (book_id, author_id)
);

CREATE TABLE genre
(
    genre_id SERIAL PRIMARY KEY,
    genre    TEXT NOT NULL,
    UNIQUE (genre)
);

CREATE TABLE book_genre
(
    book_id  INTEGER NOT NULL,
    genre_id INTEGER NOT NULL,
    FOREIGN KEY (book_id) REFERENCES book (book_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (genre_id) REFERENCES genre (genre_id) ON UPDATE CASCADE ON DELETE CASCADE,
    UNIQUE (book_id, genre_id)
);
